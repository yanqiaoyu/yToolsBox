package controller

import (
	"fmt"
	"main/common"
	"time"

	"github.com/gin-gonic/gin"
)

func task() {
	fmt.Println("I am runnning task.", time.Now())
}

// 新增一个定时任务
func PostNewCronTask(ctx *gin.Context) {
	schedule := common.GetScheduler()
	schedule.Every(1).Second().Do(task)
}
