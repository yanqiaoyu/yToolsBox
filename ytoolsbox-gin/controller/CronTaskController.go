package controller

import (
	"log"
	"main/common"
	"main/dao"
	"main/dto"
	"main/response"
	"main/service"
	"main/util"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

// 新增一个定时任务
func PostNewCronTask(ctx *gin.Context) {
	db := common.GetDB()
	// 拿到定时任务控制器的句柄
	schedule := common.GetScheduler()
	PostNewcronTaskParam := dto.PostNewcronTaskDTOReq{}

	if util.ResolveParam(ctx, &PostNewcronTaskParam) != nil {
		return
	}

	log.Print("前端传过来的新增定时任务参数", PostNewcronTaskParam)

	// 写入定时任务信息到数据库
	result, cronTaskOriginID := dao.InsertNewCronTask(db, &PostNewcronTaskParam)

	if result.Error != nil {
		msg := dto.FailResponseMeta{}
		msg.StatusCode = 400
		msg.Message = "定时任务写入数据库失败: " + result.Error.Error()
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
		return
	}

	// 如果要立即执行一次 那就立即调用一次
	if PostNewcronTaskParam.CronRunAtOnce {
		service.AddNewCronTaskService(&PostNewcronTaskParam, cronTaskOriginID)
	}

	// 添加任务, 这里可以复用普通功能的部分代码
	CronTaskScheduleID, err := schedule.AddFunc(PostNewcronTaskParam.CronTaskTime,
		func() { service.AddNewCronTaskService(&PostNewcronTaskParam, cronTaskOriginID) },
	)

	// 在刚刚插入的任务信息列表中，更新一下任务的taskID
	dao.UpdateSpecifiedCronTaskScheduleID(db, cronTaskOriginID, CronTaskScheduleID)

	log.Println(CronTaskScheduleID, err)
	if err != nil {
		msg := dto.FailResponseMeta{}
		msg.StatusCode = 400
		msg.Message = "新增定时任务失败:" + err.Error()
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
		return
	}

	// 返回
	Meta := dto.SuccessResponseMeta{Message: "新建定时任务成功", StatusCode: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 删除所有定时任务
func DeleteAllCronTask(ctx *gin.Context) {
	db := common.GetDB()
	schedule := common.GetScheduler()

	// 获取当前所有任务,AllTask是一个数组
	AllTask := schedule.Entries()

	// value中记录着每一个task的信息,是一个结构体
	for _, value := range AllTask {
		schedule.Remove(value.ID)
	}

	log.Println("删除了所有定时任务: ", AllTask)

	// 从schedule中删除后，还需要从数据库中删除
	dao.DeleteAllCronTask(db)

	Meta := dto.SuccessResponseMeta{Message: "删除所有定时任务成功", StatusCode: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 删除特定定时任务
func DeleteSpecifiedCrontask(ctx *gin.Context) {
	db := common.GetDB()
	schedule := common.GetScheduler()
	DeleteSpecifiedTaskReq := dto.DeleteSpecifiedcronTaskReq{}

	if util.ResolveURI(ctx, &DeleteSpecifiedTaskReq) != nil {
		return
	}

	// 先从schedul中移除任务
	schedule.Remove(cron.EntryID(DeleteSpecifiedTaskReq.CronTaskScheduleID))

	// 再从数据库中移除任务
	result := dao.DeleteSpecifiedCrontask(db, &DeleteSpecifiedTaskReq)
	if result.Error != nil {
		msg := dto.FailResponseMeta{}
		msg.StatusCode = 400
		msg.Message = "删除任务记录失败: " + result.Error.Error()
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
		return
	}

	Meta := dto.SuccessResponseMeta{Message: "删除定时任务成功", StatusCode: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 查询定时任务
func GetAllCronTask(ctx *gin.Context) {
	db := common.GetDB()
	GetAllCronTaskParam := dto.GetAllCronTaskDTOReq{}

	if util.ResolveParam(ctx, &GetAllCronTaskParam) != nil {
		return
	}

	cronTaskItemList, DefaultLength := dao.SelectAllCronTask(db, GetAllCronTaskParam)

	// 构造返回的结构体
	cronTaskItemData := dto.GetAllCronTaskItemDTOResp{Total: int64(DefaultLength), CronTaskItemList: cronTaskItemList}
	Meta := dto.SuccessResponseMeta{Message: "获取定时任务列表成功", StatusCode: 200}

	response.Success(ctx, util.Struct2MapViaJson(cronTaskItemData), util.Struct2MapViaJson(Meta))
}

// 根据scheduleID查询特定定时任务
func GetSpecifiedCrontaskByScheduleID(ctx *gin.Context) {
	db := common.GetDB()
	GetSpecifiedCrontaskByScheduleIDParam := dto.GetSpecifiedCrontaskByScheduleIDDTOReq{}

	if util.ResolveURI(ctx, &GetSpecifiedCrontaskByScheduleIDParam) != nil {
		return
	}

	result, cronTaskItem := dao.SelectSpecifiedCronTaskByScheduleID(db, GetSpecifiedCrontaskByScheduleIDParam.CronTaskScheduleID)
	if result.Error != nil {
		msg := dto.FailResponseMeta{}
		msg.StatusCode = 400
		msg.Message = "查询定时任务信息失败: " + result.Error.Error()
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
		return
	}

	log.Println("获取特定的定时任务信息: ", cronTaskItem)

	// 构造返回的结构体
	Meta := dto.SuccessResponseMeta{Message: "获取定时任务列表成功", StatusCode: 200}
	response.Success(ctx, util.Struct2MapViaJson(cronTaskItem), util.Struct2MapViaJson(Meta))
}

// 查询所有定时任务执行结果
func GetAllCronTaskResult(ctx *gin.Context) {
	db := common.GetDB()
	GetAllCronTaskResultParam := dto.GetAllCronTaskResultDTOReq{}
	if util.ResolveParam(ctx, &GetAllCronTaskResultParam) != nil {
		return
	}

	cronTaskResutlItemList, DefaultLength := dao.SelectAllCronTaskResult(db, GetAllCronTaskResultParam)

	// 构造返回的结构体
	cronTaskItemData := dto.GetAllCronTaskItemDTOResp{Total: int64(DefaultLength), CronTaskItemList: cronTaskResutlItemList}
	Meta := dto.SuccessResponseMeta{Message: "获取定时任务执行结果列表成功", StatusCode: 200}

	response.Success(ctx, util.Struct2MapViaJson(cronTaskItemData), util.Struct2MapViaJson(Meta))
}

// 删除所有定时任务执行结果
func DeleteAllCronTaskResult(ctx *gin.Context) {
	db := common.GetDB()

	// 从数据库中删除
	dao.DeleteAllCronTaskResult(db)

	Meta := dto.SuccessResponseMeta{Message: "删除所有定时任务成功", StatusCode: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 删除特定定时任务执行结果
func DeleteSpecifiedCrontaskResult(ctx *gin.Context) {
	db := common.GetDB()
	DeleteSpecifiedcronTaskResultParam := dto.DeleteSpecifiedcronTaskResultReq{}

	if util.ResolveURI(ctx, &DeleteSpecifiedcronTaskResultParam) != nil {
		return
	}

	// 再从数据库中移除任务
	result := dao.DeleteSpecifiedCrontaskResult(db, &DeleteSpecifiedcronTaskResultParam)
	if result.Error != nil {
		msg := dto.FailResponseMeta{}
		msg.StatusCode = 400
		msg.Message = "删除定时任务执行记录失败: " + result.Error.Error()
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
		return
	}

	Meta := dto.SuccessResponseMeta{Message: "删除定时任务结果成功", StatusCode: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}
