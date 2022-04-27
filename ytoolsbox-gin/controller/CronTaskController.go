package controller

import (
	"fmt"
	"log"
	"main/common"
	"time"

	"github.com/gin-gonic/gin"
)

func task(count int) {
	fmt.Println("I am runnning task.", time.Now(), count)
}

// 新增一个定时任务
func PostNewCronTask(ctx *gin.Context) {
	schedule := common.GetScheduler()
	taskID, err := schedule.AddFunc("*/1 * * * * ?", func() { task(1) })

	log.Println(taskID, err)
}

// 删除所有定时任务
func DeleteAllCronTask(ctx *gin.Context) {
	schedule := common.GetScheduler()
	AllTask := schedule.Entries()
	log.Println("AllTask: ", AllTask)

	for _, value := range AllTask {
		schedule.Remove(value.ID)
		// TaskMap := util.Struct2MapViaReflect(value)
		// log.Println(TaskMap)
	}

	// schedule.re
}
