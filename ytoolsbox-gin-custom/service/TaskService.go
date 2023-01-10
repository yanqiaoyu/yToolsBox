package service

import (
	"main/dto"
	"main/model"
	"main/utils"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CreateNewTaskService(config dto.BriefToolConfigDTO, resultChannel chan model.Tasks) error {
	utils.ExecuteTask(
		config.ToolExecuteLocation,
		resultChannel,
		config.ToolRemoteSSH_Port,
		config.ToolRemoteIP,
		config.ToolRemoteSSH_Account,
		config.ToolRemoteSSH_Password,
		config.ToolType,
		config.ToolRunCMD,
		config.ToolScriptLocalPath,
		config.ToolScriptPath,
		config.ToolScriptName,
	)
	// 关闭resultChannel
	utils.CloseMyChannel(resultChannel)
	return nil
}

/*
func CreateNewTaskService(config dto.BriefToolConfigDTO, resultChannel chan model.Tasks) error {

	// 存放任务执行结果的缓存
	buf := bytes.Buffer{}
	// Task在工具盒所在的机器上执行
	if config.ToolExecuteLocation == "local" {
		log.Println(">>> 直接本地执行")
		buf.WriteString(">>> 直接本地执行\r\n")
		resultChannel <- model.Tasks{Progress: 25, ReturnContent: buf.String()}

		// 准备好连接本地的素材
		port, _ := strconv.Atoi(config.ToolRemoteSSH_Port)
		cliConf := new(utils.ClientConfig)
		errCreateClient := cliConf.CreateClient(
			config.ToolRemoteIP,
			int64(port),
			config.ToolRemoteSSH_Account,
			config.ToolRemoteSSH_Password)

		if errCreateClient != nil {
			log.Println(">>> 连接异常")
			buf.WriteString(">>> 连接异常\r\n")
			buf.WriteString(errCreateClient.Error())
			buf.WriteString("\r\n")
			resultChannel <- model.Tasks{Progress: 100, ReturnContent: buf.String(), IsDone: true}
			close(resultChannel)
			return nil
		}

		// 2.执行的是容器还是脚本?
		if config.ToolType == "container" {
			log.Println(">>> 执行的是容器工具")
			buf.WriteString(">>> 执行的是容器工具\r\n")
			resultChannel <- model.Tasks{Progress: 50, ReturnContent: buf.String()}

			// 3.开始执行
			log.Println(">>> 开始执行 \r\n", config.ToolRunCMD)
			buf.WriteString(">>> 开始执行 \r\n")
			buf.WriteString(config.ToolRunCMD)
			buf.WriteString("\r\n")
			ExecuteResult := cliConf.RunShell(config.ToolRunCMD)
			resultChannel <- model.Tasks{Progress: 75, ReturnContent: buf.String()}

			// 4.获取执行结果，后续这里要改进，另起一个goroutine，持续获取执行情况
			log.Println(">>> 执行结果", ExecuteResult)
			buf.WriteString(">>> 执行结果")
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			buf.WriteString(ExecuteResult)
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			resultChannel <- model.Tasks{Progress: 100, ReturnContent: buf.String(), IsDone: true}

		} else if config.ToolType == "script" {
			// 执行脚本工具
			log.Println(">>> 执行的是脚本工具")
			buf.WriteString(">>> 执行的是脚本工具\r\n")
			resultChannel <- model.Tasks{Progress: 40, ReturnContent: buf.String()}

			// 3.开始执行
			log.Println(">>> 开始执行 \r\n", config.ToolRunCMD)
			buf.WriteString(">>> 开始执行 \r\n")
			buf.WriteString(config.ToolRunCMD)
			buf.WriteString("\r\n")
			resultChannel <- model.Tasks{Progress: 85, ReturnContent: buf.String()}

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
			resultChannel <- model.Tasks{Progress: 100, ReturnContent: buf.String(), IsDone: true}
		}

	} else if config.ToolExecuteLocation == "remote" { //在别的远程环境执行
		log.Println(">>> 进入远程执行")
		buf.WriteString(">>> 进入远程执行\r\n")
		resultChannel <- model.Tasks{Progress: 25, ReturnContent: buf.String()}
		// 准备好远程连接的素材
		port, _ := strconv.Atoi(config.ToolRemoteSSH_Port)
		cliConf := new(utils.ClientConfig)
		errCreateClient := cliConf.CreateClient(
			config.ToolRemoteIP,
			int64(port),
			config.ToolRemoteSSH_Account,
			config.ToolRemoteSSH_Password)

		if errCreateClient != nil {
			log.Println(">>> 连接异常")
			buf.WriteString(">>> 连接异常\r\n")
			buf.WriteString(errCreateClient.Error())
			buf.WriteString("\r\n")
			resultChannel <- model.Tasks{Progress: 100, ReturnContent: buf.String(), IsDone: true}
			close(resultChannel)
			return nil
		}

		// 2.执行的是容器还是脚本?
		if config.ToolType == "container" {
			log.Println(">>> 执行的是容器工具")
			buf.WriteString(">>> 执行的是容器工具\r\n")
			resultChannel <- model.Tasks{Progress: 50, ReturnContent: buf.String()}

			// 3.开始执行
			log.Println(">>> 开始执行 '\r\n", config.ToolRunCMD)
			buf.WriteString(">>> 开始执行 \r\n")
			buf.WriteString(config.ToolRunCMD)
			buf.WriteString("\r\n")
			ExecuteResult := cliConf.RunShell(config.ToolRunCMD)
			resultChannel <- model.Tasks{Progress: 75, ReturnContent: buf.String()}

			// 4.获取执行结果，后续这里要改进，另起一个goroutine，持续获取执行情况
			buf.WriteString(">>> 执行结果")
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			buf.WriteString(ExecuteResult)
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			resultChannel <- model.Tasks{Progress: 100, ReturnContent: buf.String(), IsDone: true}

		} else if config.ToolType == "script" {
			// 执行脚本工具
			log.Println(">>> 执行的是脚本工具")
			buf.WriteString(">>> 执行的是脚本工具\r\n")
			resultChannel <- model.Tasks{Progress: 40, ReturnContent: buf.String()}

			// 脚本工具，需要先上传脚本到指定位置
			log.Println(">>> 上传脚本至指定位置", config.ToolScriptPath+config.ToolScriptName)
			buf.WriteString(">>> 上传脚本 \r\n" + config.ToolScriptLocalPath)
			buf.WriteString("\r\n>>> 至指定位置 \r\n")
			buf.WriteString(config.ToolScriptPath + config.ToolScriptName)
			buf.WriteString("\r\n")
			resultChannel <- model.Tasks{Progress: 60, ReturnContent: buf.String()}
			uploadResult := cliConf.Upload(config.ToolScriptLocalPath, config.ToolScriptPath+config.ToolScriptName)
			log.Println(">>> 上传结果 \r\n", uploadResult)
			buf.WriteString(">>> 上传结果 \r\n")
			buf.WriteString(uploadResult)
			buf.WriteString("\r\n")
			resultChannel <- model.Tasks{Progress: 70, ReturnContent: buf.String()}

			// 4.开始执行
			log.Println(">>> 开始执行 \r\n", config.ToolRunCMD)
			buf.WriteString(">>> 开始执行 \r\n")
			buf.WriteString(config.ToolRunCMD)
			buf.WriteString("\r\n")
			resultChannel <- model.Tasks{Progress: 85, ReturnContent: buf.String()}

			ExecuteResult := cliConf.RunShell(config.ToolRunCMD)
			// 5.获取执行结果，后续这里要改进，另起一个goroutine，持续获取执行情况
			buf.WriteString(">>> 执行结果")
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			buf.WriteString(ExecuteResult)
			buf.WriteString("\r\n----------------------**********----------------------\r\n")
			resultChannel <- model.Tasks{Progress: 100, ReturnContent: buf.String(), IsDone: true}
		}
	}

	// 关闭resultChannel
	close(resultChannel)
	return nil
}
*/
func SaveScriptFile(ctx *gin.Context) (string, string, error) {
	// 获取文件名
	file, err := ctx.FormFile("file")
	if err != nil {
		return "", "", err
	}
	fileName := file.Filename
	//获取工具名
	toolName := ctx.PostForm("toolName")
	// log.Println("工具名：", toolName)
	// log.Println("文件名：", fileName)

	// 创建Base路径
	AbsPath, _ := os.Getwd()
	BasePath := viper.GetString("tool.scriptBasePath")
	FinalFilePath := filepath.Join(AbsPath, BasePath, toolName)
	// log.Println("文件存放路径: ", FinalFilePath)

	utils.CreateDir(FinalFilePath)

	FileDST := filepath.Join(FinalFilePath, fileName)

	//保存文件到服务器本地
	//SaveUploadedFile(文件头，保存路径)
	if err := ctx.SaveUploadedFile(file, FileDST); err != nil {
		return "", "", err
	}
	return toolName, FileDST, nil
}
