package service

import (
	"log"
	"main/common"
	"main/dao"
	"main/dto"
	"main/model"
	"strings"
)

func AddNewCronTaskService(PostNewcronTaskParam *dto.PostNewcronTaskDTOReq) {
	db := common.GetDB()
	// 这里获取到的PostNewTaskParam是一个字符串形式的数组，所以还需要处理
	PostNewcronTaskParam.CronTaskFinalList = strings.TrimPrefix(PostNewcronTaskParam.CronTaskFinalList, "[")
	PostNewcronTaskParam.CronTaskFinalList = strings.TrimSuffix(PostNewcronTaskParam.CronTaskFinalList, "]")

	configIDList := strings.Split(PostNewcronTaskParam.CronTaskFinalList, ",")
	log.Println(configIDList)

	// 获取每一个配置ID，新增一个任务，新增一个任务进度条目
	for i := 0; i < len(configIDList); i++ {
		// 从库里面把这个配置ID的全部信息查出来传给service
		configList := dao.SelectConfigByToolID(db, string(configIDList[i]))
		// 新增一个任务条目
		TaskID := dao.InsertTaskItem(db, configList)

		// 线程间通信需要用到的chan
		resultChannel := make(chan model.Tasks, 10)
		// 用于接收业务执行结果并更新至数据库
		go dao.UpdateTaskProgress(db, resultChannel, TaskID)
		// 用于关键业务的执行
		go CreateNewTaskService(configList, resultChannel)
	}
}
