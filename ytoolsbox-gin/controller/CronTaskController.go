package controller

import (
	"fmt"
	"log"
	"main/common"
	"main/dto"
	"main/response"
	"main/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func task(count int) {
	fmt.Println("I am runnning task.", time.Now(), count)
}

// 新增一个定时任务
func PostNewCronTask(ctx *gin.Context) {
	schedule := common.GetScheduler()
	taskID, err := schedule.AddFunc("*/1 * * * * ?", func() { task(1) })
	log.Println(taskID, err)
	if err != nil {
		msg := dto.FailResponseMeta{}
		msg.StatusCode = 400
		msg.Message = "新增定时任务失败"
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
		return
	}

	Meta := dto.SuccessResponseMeta{Message: "新建定时任务成功", StatusCode: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 删除所有定时任务
func DeleteAllCronTask(ctx *gin.Context) {
	schedule := common.GetScheduler()

	// 获取当前所有任务,AllTask是一个数组
	AllTask := schedule.Entries()

	// value中记录着每一个task的信息,是一个结构体
	for _, value := range AllTask {
		schedule.Remove(value.ID)
	}
	Meta := dto.SuccessResponseMeta{Message: "删除所有定时任务成功", StatusCode: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 删除特定定时任务
func DeleteSpecifiedCrontask(ctx *gin.Context) {
	schedule := common.GetScheduler()
	DeleteSpecifiedTaskReq := dto.DeleteSpecifiedCronTaskReq{}

	if util.ResolveURI(ctx, &DeleteSpecifiedTaskReq) != nil {
		return
	}

	schedule.Remove(cron.EntryID(DeleteSpecifiedTaskReq.TaskID))
	Meta := dto.SuccessResponseMeta{Message: "删除定时任务成功", StatusCode: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}
