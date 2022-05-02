package dao

import (
	"main/dto"
	"main/model"
	"main/util"

	"gorm.io/gorm"
)

// 插入定时任务
func InsertNewCronTask(db *gorm.DB, PostNewcronTaskDTOReq *dto.PostNewcronTaskDTOReq) *gorm.DB {
	// 构造写入数据库的结构体
	CronTaskItem := model.CronTasks{
		CronTaskName:      PostNewcronTaskDTOReq.CronTaskName,
		CronTaskDesc:      PostNewcronTaskDTOReq.CronTaskDesc,
		CronTaskFinalList: PostNewcronTaskDTOReq.CronTaskFinalList,
		CronTaskTime:      PostNewcronTaskDTOReq.CronTaskTime,
		CronRunAtOnce:     PostNewcronTaskDTOReq.CronRunAtOnce,
	}
	// 写入数据库
	result := db.Debug().Model(&model.CronTasks{}).Create(&CronTaskItem)
	return result
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
		map_item := util.Struct2MapViaJson(cronTaskItem[i])
		map_taskItemList = append(map_taskItemList, map_item)
	}

	// 计算一下需要如何切割数组
	ArrayStart, ArrayEnd := util.CalculateReturnMapLength(pagenum, pagesize, map_taskItemList)
	// 返回切片后的结果
	return map_taskItemList[ArrayStart:ArrayEnd], DefaultLength
}

// 删除所有任务
func DeleteAllCronTask(db *gorm.DB) {
	// 根据gorm的官方文档，如果在没有任何条件的情况下执行批量删除，GORM 不会执行该操作
	// 必须加一些条件
	db.Unscoped().Where("\"id\" != -1").Delete(&model.CronTasks{})
}
