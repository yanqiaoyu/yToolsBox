package service

import (
	"bytes"
	"log"
	"main/common"
	"main/dao"
	"main/dto"
	"main/model"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func AddNewCronTaskService(PostNewcronTaskParam *dto.PostNewcronTaskDTOReq, cronTaskOriginID uint) {
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
		// 新增一个定时任务条目
		cronTaskID := dao.InsertCronTaskItem(db, configList, cronTaskOriginID)

		// 线程间通信需要用到的chan
		cronTaskResultChannel := make(chan model.CronTasksResult, 10)
		// 用于接收业务执行结果并更新至数据库
		go dao.UpdateCronTaskProgress(db, cronTaskResultChannel, cronTaskID)
		// 用于关键业务的执行
		go CreateNewCronTaskService(configList, cronTaskResultChannel)
	}
}

func CreateNewCronTaskService(config dto.BriefToolConfigDTO, resultChannel chan model.CronTasksResult) error {
	log.Println(config)

	// 存放任务执行结果的缓存
	buf := bytes.Buffer{}
	// Task在工具盒所在的机器上执行
	if config.ToolExecuteLocation == "local" {
		log.Println(">>> 直接本地执行")
		buf.WriteString(">>> 直接本地执行\r\n")
		resultChannel <- model.CronTasksResult{Progress: 25, ReturnContent: buf.String()}

		// 准备好连接本地的素材
		port, _ := strconv.Atoi(config.ToolRemoteSSH_Port)
		cliConf := new(ClientConfig)
		errCreateClient := cliConf.createClient(
			config.ToolRemoteIP,
			int64(port),
			config.ToolRemoteSSH_Account,
			config.ToolRemoteSSH_Password)

		if errCreateClient != nil {
			log.Println(">>> 连接异常")
			buf.WriteString(">>> 连接异常\r\n")
			buf.WriteString(errCreateClient.Error())
			buf.WriteString("\r\n")
			resultChannel <- model.CronTasksResult{Progress: 100, ReturnContent: buf.String(), IsDone: true}
			close(resultChannel)
			return nil
		}

		// 2.执行的是容器还是脚本?
		if config.ToolType == "container" {
			log.Println(">>> 执行的是容器工具")
			buf.WriteString(">>> 执行的是容器工具\r\n")
			resultChannel <- model.CronTasksResult{Progress: 50, ReturnContent: buf.String()}

			// 3.开始执行
			log.Println(">>> 开始执行 \r\n", config.ToolRunCMD)
			buf.WriteString(">>> 开始执行 \r\n")
			buf.WriteString(config.ToolRunCMD)
			buf.WriteString("\r\n")
			ExecuteResult := cliConf.RunShell(config.ToolRunCMD)
			resultChannel <- model.CronTasksResult{Progress: 75, ReturnContent: buf.String()}

			// 4.获取执行结果，后续这里要改进，另起一个goroutine，持续获取执行情况
			log.Println(">>> 执行结果", ExecuteResult)
			buf.WriteString(">>> 执行结果")
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			buf.WriteString(ExecuteResult)
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			resultChannel <- model.CronTasksResult{Progress: 100, ReturnContent: buf.String(), IsDone: true}

		} else if config.ToolType == "script" {
			// 执行脚本工具
			log.Println(">>> 执行的是脚本工具")
			buf.WriteString(">>> 执行的是脚本工具\r\n")
			resultChannel <- model.CronTasksResult{Progress: 40, ReturnContent: buf.String()}

			// 3.开始执行
			log.Println(">>> 开始执行 \r\n", config.ToolRunCMD)
			buf.WriteString(">>> 开始执行 \r\n")
			buf.WriteString(config.ToolRunCMD)
			buf.WriteString("\r\n")
			resultChannel <- model.CronTasksResult{Progress: 85, ReturnContent: buf.String()}

			// 这是文件在宿主机存放的路径
			HOST_SCRIPT_PATH := os.Getenv("HOST_SCRIPT_PATH")
			//
			sysType := runtime.GOOS
			var tmpStr []string
			if sysType == "linux" {
				tmpStr = strings.Split(config.ToolScriptLocalPath, "/")

			} else if sysType == "windows" {
				tmpStr = strings.Split(config.ToolScriptLocalPath, "\\")
			}

			finalShell := "cd " + HOST_SCRIPT_PATH + " && " + "cd " + tmpStr[len(tmpStr)-2] + " && " + config.ToolRunCMD

			ExecuteResult := cliConf.RunShell(finalShell)
			// 4.获取执行结果
			log.Println(">>> 执行结果", ExecuteResult)
			buf.WriteString(">>> 执行结果")
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			buf.WriteString(ExecuteResult)
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			resultChannel <- model.CronTasksResult{Progress: 100, ReturnContent: buf.String(), IsDone: true}
		}

	} else if config.ToolExecuteLocation == "remote" { //在别的远程环境执行
		log.Println(">>> 进入远程执行")
		buf.WriteString(">>> 进入远程执行\r\n")
		resultChannel <- model.CronTasksResult{Progress: 25, ReturnContent: buf.String()}
		// 准备好远程连接的素材
		port, _ := strconv.Atoi(config.ToolRemoteSSH_Port)
		cliConf := new(ClientConfig)
		errCreateClient := cliConf.createClient(
			config.ToolRemoteIP,
			int64(port),
			config.ToolRemoteSSH_Account,
			config.ToolRemoteSSH_Password)

		if errCreateClient != nil {
			log.Println(">>> 连接异常")
			buf.WriteString(">>> 连接异常\r\n")
			buf.WriteString(errCreateClient.Error())
			buf.WriteString("\r\n")
			resultChannel <- model.CronTasksResult{Progress: 100, ReturnContent: buf.String(), IsDone: true}
			close(resultChannel)
			return nil
		}

		// 2.执行的是容器还是脚本?
		if config.ToolType == "container" {
			log.Println(">>> 执行的是容器工具")
			buf.WriteString(">>> 执行的是容器工具\r\n")
			resultChannel <- model.CronTasksResult{Progress: 50, ReturnContent: buf.String()}

			// 3.开始执行
			log.Println(">>> 开始执行 '\r\n", config.ToolRunCMD)
			buf.WriteString(">>> 开始执行 \r\n")
			buf.WriteString(config.ToolRunCMD)
			buf.WriteString("\r\n")
			ExecuteResult := cliConf.RunShell(config.ToolRunCMD)
			resultChannel <- model.CronTasksResult{Progress: 75, ReturnContent: buf.String()}

			// 4.获取执行结果，后续这里要改进，另起一个goroutine，持续获取执行情况
			buf.WriteString(">>> 执行结果")
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			buf.WriteString(ExecuteResult)
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			resultChannel <- model.CronTasksResult{Progress: 100, ReturnContent: buf.String(), IsDone: true}

		} else if config.ToolType == "script" {
			// 执行脚本工具
			log.Println(">>> 执行的是脚本工具")
			buf.WriteString(">>> 执行的是脚本工具\r\n")
			resultChannel <- model.CronTasksResult{Progress: 40, ReturnContent: buf.String()}

			// 脚本工具，需要先上传脚本到指定位置
			log.Println(">>> 上传脚本至指定位置", config.ToolScriptPath+config.ToolScriptName)
			buf.WriteString(">>> 上传脚本 \r\n" + config.ToolScriptLocalPath)
			buf.WriteString("\r\n>>> 至指定位置 \r\n")
			buf.WriteString(config.ToolScriptPath + config.ToolScriptName)
			buf.WriteString("\r\n")
			resultChannel <- model.CronTasksResult{Progress: 60, ReturnContent: buf.String()}
			uploadResult := cliConf.Upload(config.ToolScriptLocalPath, config.ToolScriptPath+config.ToolScriptName)
			log.Println(">>> 上传结果 \r\n", uploadResult)
			buf.WriteString(">>> 上传结果 \r\n")
			buf.WriteString(uploadResult)
			buf.WriteString("\r\n")
			resultChannel <- model.CronTasksResult{Progress: 70, ReturnContent: buf.String()}

			// 4.开始执行
			log.Println(">>> 开始执行 \r\n", config.ToolRunCMD)
			buf.WriteString(">>> 开始执行 \r\n")
			buf.WriteString(config.ToolRunCMD)
			buf.WriteString("\r\n")
			resultChannel <- model.CronTasksResult{Progress: 85, ReturnContent: buf.String()}

			ExecuteResult := cliConf.RunShell(config.ToolRunCMD)
			// 5.获取执行结果，后续这里要改进，另起一个goroutine，持续获取执行情况
			buf.WriteString(">>> 执行结果")
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			buf.WriteString(ExecuteResult)
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			resultChannel <- model.CronTasksResult{Progress: 100, ReturnContent: buf.String(), IsDone: true}
		}
	}

	// 关闭resultChannel
	close(resultChannel)
	return nil
}
