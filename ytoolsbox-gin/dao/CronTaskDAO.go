package dao

import (
	"log"
	"main/dto"
	"main/model"
	"main/utils"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

// 插入定时任务
func InsertNewCronTask(db *gorm.DB, PostNewcronTaskDTOReq *dto.PostNewcronTaskDTOReq) (*gorm.DB, uint) {
	// 构造写入数据库的结构体
	CronTaskItem := model.CronTasks{
		CronTaskName:                 PostNewcronTaskDTOReq.CronTaskName,
		CronTaskDesc:                 PostNewcronTaskDTOReq.CronTaskDesc,
		CronTaskFinalList:            PostNewcronTaskDTOReq.CronTaskFinalList,
		CronTaskTime:                 PostNewcronTaskDTOReq.CronTaskTime,
		CronRunAtOnce:                PostNewcronTaskDTOReq.CronRunAtOnce,
		CronTaskCascaderAllInfo:      PostNewcronTaskDTOReq.CronTaskCascaderAllInfo,
		CronTaskCascaderSelectedInfo: PostNewcronTaskDTOReq.CronTaskCascaderSelectedInfo,
	}
	// 写入数据库
	result := db.Debug().Model(&model.CronTasks{}).Create(&CronTaskItem)
	return result, CronTaskItem.ID
}

// 新增一个定时任务执行结果的条目
func InsertCronTaskItem(db *gorm.DB, config dto.BriefToolConfigDTO, cronTaskOriginID uint) uint {
	// 这里是查一些工具的信息
	tool := model.Tool{}
	db.Model(&model.Tool{}).Select("\"toolName\"").Where("id = ?", config.ToolID).First(&tool)

	// 这里是查一些定时任务的信息
	cronTask := model.CronTasks{}

	db.Model(&model.CronTasks{}).Where("id = ?", cronTaskOriginID).Find(&cronTask)

	CronTaskItem := model.CronTasksResult{
		ToolConfigName:     config.ToolConfigName,
		ToolName:           tool.ToolName,
		Progress:           0,
		IsDone:             false,
		ReturnContent:      "任务开始:\r\n",
		CronTaskID:         cronTaskOriginID,
		CronTaskScheduleID: uint(cronTask.CronTaskScheduleID),
		CronTaskName:       cronTask.CronTaskName,
		CronTaskDesc:       cronTask.CronTaskDesc,
	}
	db.Model(&model.CronTasksResult{}).Create(&CronTaskItem)
	return CronTaskItem.ID
}

// 根据CronTask ID 更新定时任务详情条目
func UpdateCronTaskProgress(db *gorm.DB, resultChannel chan model.CronTasksResult, cronTaskID uint) {
	// 我们可以使用for循环，持续的从一个channel中接受数据
	// 当channel为空时，for循环会被阻塞。当channel被关闭时，则会跳出for循环。
	for result := range resultChannel {
		// log.Printf("receive %s\n", result)
		db.Model(&model.CronTasksResult{}).Where("id = ?", cronTaskID).Updates(result)
	}
	log.Println("管道被关闭，结束这个更新定时任务信息的协程")
}

// 查询定时任务
func SelectAllCronTask(db *gorm.DB, GetAllCronTaskParam dto.GetAllCronTaskDTOReq) ([]map[string]interface{}, int) {
	cronTaskItem := []model.CronTasks{}
	query := GetAllCronTaskParam.Query
	pagenum := GetAllCronTaskParam.Pagenum
	pagesize := GetAllCronTaskParam.Pagesize
	map_taskItemList := []map[string]interface{}{}

	// 不带Query，返回全部
	// 否则返回like搜索后的结果
	if query == "" {
		// 按照时间升序
		db.Order("created_at desc").Model(&model.CronTasks{}).Find(&cronTaskItem)
	} else {
		// 按照时间升序
		db.Order("created_at desc").Where("\"cronTaskName\" LIKE ?", "%"+query+"%").Model(&model.CronTasks{}).Find(&cronTaskItem)
	}

	DefaultLength := len(cronTaskItem)

	// 把一个自定义结构体的array 转换成map的array
	// 这里用了json的方法 虽然效率低 但是解决了返回给前端大小写的问题
	for i := 0; i < DefaultLength; i++ {
		map_item := utils.Struct2MapViaJson(cronTaskItem[i])
		map_taskItemList = append(map_taskItemList, map_item)
	}

	// 计算一下需要如何切割数组
	ArrayStart, ArrayEnd := utils.CalculateReturnMapLength(pagenum, pagesize, map_taskItemList)
	// 返回切片后的结果
	return map_taskItemList[ArrayStart:ArrayEnd], DefaultLength
}

// 查询所有定时任务执行结果
func SelectAllCronTaskResult(db *gorm.DB, GetAllCronTaskResultParam dto.GetAllCronTaskResultDTOReq) ([]map[string]interface{}, int) {
	cronTaskResultItem := []model.CronTasksResult{}
	query := GetAllCronTaskResultParam.Query
	pagenum := GetAllCronTaskResultParam.Pagenum
	pagesize := GetAllCronTaskResultParam.Pagesize
	map_cronTaskResultList := []map[string]interface{}{}

	// 不带Query，返回全部
	// 否则返回like搜索后的结果
	if query == "" {
		// 按照时间升序
		db.Order("created_at desc").Model(&model.CronTasksResult{}).Find(&cronTaskResultItem)
	} else {
		// 按照时间升序
		db.Order("created_at desc").Where("\"cronTaskName\" LIKE ?", "%"+query+"%").Model(&model.CronTasks{}).Find(&cronTaskResultItem)
	}

	DefaultLength := len(cronTaskResultItem)

	// 把一个自定义结构体的array 转换成map的array
	// 这里用了json的方法 虽然效率低 但是解决了返回给前端大小写的问题
	for i := 0; i < DefaultLength; i++ {
		map_item := utils.Struct2MapViaJson(cronTaskResultItem[i])
		map_cronTaskResultList = append(map_cronTaskResultList, map_item)
	}

	// 计算一下需要如何切割数组
	ArrayStart, ArrayEnd := utils.CalculateReturnMapLength(pagenum, pagesize, map_cronTaskResultList)
	// 返回切片后的结果
	return map_cronTaskResultList[ArrayStart:ArrayEnd], DefaultLength
}

// 根据scheduleID查询特定定时任务
func SelectSpecifiedCronTaskByScheduleID(db *gorm.DB, CronTaskScheduleID uint) (*gorm.DB, model.CronTasks) {
	cronTaskItem := model.CronTasks{}
	result := db.Debug().Model(&model.CronTasks{}).Where("\"cronTaskScheduleID\" = ?", CronTaskScheduleID).Order("created_at desc").Find(&cronTaskItem)
	return result, cronTaskItem
}

// 删除所有定时任务
func DeleteAllCronTask(db *gorm.DB) {
	// 根据gorm的官方文档，如果在没有任何条件的情况下执行批量删除，GORM 不会执行该操作
	// 必须加一些条件
	db.Unscoped().Where("\"id\" != -1").Delete(&model.CronTasks{})
}

// 删除特定定时任务
func DeleteSpecifiedCrontask(db *gorm.DB, DeleteSpecifiedcronTaskReq *dto.DeleteSpecifiedcronTaskReq) *gorm.DB {
	return db.Debug().Delete(&model.CronTasks{}, DeleteSpecifiedcronTaskReq.CronTaskOriginID, DeleteSpecifiedcronTaskReq.CronTaskScheduleID)
}

// 删除所有定时任务执行结果
func DeleteAllCronTaskResult(db *gorm.DB) {
	// 根据gorm的官方文档，如果在没有任何条件的情况下执行批量删除，GORM 不会执行该操作
	// 必须加一些条件
	db.Unscoped().Where("\"id\" != -1").Delete(&model.CronTasksResult{})
}

// 删除特定定时任务执行结果
func DeleteSpecifiedCrontaskResult(db *gorm.DB, DeleteSpecifiedcronTaskResultReq *dto.DeleteSpecifiedcronTaskResultReq) *gorm.DB {
	return db.Debug().Delete(&model.CronTasksResult{}, DeleteSpecifiedcronTaskResultReq.CronTaskResultID)
}

// 更新特定的定时任务的taskID
func UpdateSpecifiedCronTaskScheduleID(db *gorm.DB, cronTaskOriginID uint, CronTaskScheduleID cron.EntryID) {
	db.Model(&model.CronTasks{}).Where("id = ?", cronTaskOriginID).Update("cronTaskScheduleID", CronTaskScheduleID)
}
