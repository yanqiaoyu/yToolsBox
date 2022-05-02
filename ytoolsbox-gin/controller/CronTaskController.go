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
	result := dao.InsertNewCronTask(db, &PostNewcronTaskParam)

	if result.Error != nil {
		msg := dto.FailResponseMeta{}
		msg.StatusCode = 400
		msg.Message = "定时任务写入数据库失败: " + result.Error.Error()
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
		return
	}

	// 添加任务, 这里可以复用普通功能的部分代码
	taskID, err := schedule.AddFunc(PostNewcronTaskParam.CronTaskTime,
		func() { service.AddNewCronTaskService(&PostNewcronTaskParam) },
	)
	log.Println(taskID, err)
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
	schedule := common.GetScheduler()
	DeleteSpecifiedTaskReq := dto.DeleteSpecifiedcronTaskReq{}

	if util.ResolveURI(ctx, &DeleteSpecifiedTaskReq) != nil {
		return
	}

	schedule.Remove(cron.EntryID(DeleteSpecifiedTaskReq.TaskID))
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

	// // 构造返回的结构体
	cronTaskItemData := dto.GetAllCronTaskItemDTOResp{Total: int64(DefaultLength), CronTaskItemList: cronTaskItemList}
	Meta := dto.SuccessResponseMeta{Message: "获取定时任务列表成功", StatusCode: 200}

	response.Success(ctx, util.Struct2MapViaJson(cronTaskItemData), util.Struct2MapViaJson(Meta))
}
