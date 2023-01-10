package utils

import (
	"bytes"
	"errors"
	"log"
	"main/model"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// 处理从前端传过来的Config列表
// Task和CronTask可以复用这里
func TreatTaskConfigListFromFrontEnd(ConfigListString string) (TreatedConfigListString []string) {
	// log.Println("未处理前的ConfigList: ", ConfigListString)
	ConfigListString = strings.TrimPrefix(ConfigListString, "[")
	ConfigListString = strings.TrimSuffix(ConfigListString, "]")
	// log.Println("处理后的ConfigList: ", ConfigListString)

	TreatedConfigListString = strings.Split(ConfigListString, ",")
	// log.Println("最终字符串数组化了的ConfigList: ", TreatedConfigListString)
	return TreatedConfigListString
}

// 执行任务
// Task和CronTask可以复用这里
func ExecuteTask(
	// 执行位置
	ToolExecuteLocation string,
	// 接收执行结果的管道
	resultChannel interface{},
	// ssh port
	SSHPort string,
	// ssh ip
	SSHIPAddr string,
	// ssh account
	SSHAcount string,
	// ssh passwd
	SSHPasswd string,

	// local相关
	ToolType string,
	ToolRunCMD string,
	ToolScriptLocalPath string,
	ToolScriptPath string,
	ToolScriptName string,
) error {

	// 存放任务执行结果的缓存
	buf := bytes.Buffer{}

	// 准备好连接本地的素材
	cliConf, err := makeSSHClient(buf, resultChannel, SSHPort, SSHIPAddr, SSHAcount, SSHPasswd)
	if err != nil {
		return err
	}

	// 本地执行
	if ToolExecuteLocation == "local" {
		write2logAndBufAndChan(">>> 直接本地执行\r\n", &buf, resultChannel, 25, false)

		// 2.执行的是容器还是脚本?
		if ToolType == "container" {
			executeContainer(buf, resultChannel, ToolRunCMD, cliConf)
		} else if ToolType == "script" {
			// 执行脚本工具
			write2logAndBufAndChan(">>> 执行的是脚本工具\r\n", &buf, resultChannel, 40, false)

			// 3.开始执行
			write2logAndBufAndChan(">>> 开始执行 \r\n"+ToolRunCMD+"\r\n", &buf, resultChannel, 85, false)

			// 这是文件在宿主机存放的路径,这个环境变量定义在dockre-compose文件中
			HOST_SCRIPT_PATH := os.Getenv("HOST_SCRIPT_PATH")
			sysType := runtime.GOOS
			var tmpStr []string
			if sysType == "linux" {
				tmpStr = strings.Split(ToolScriptLocalPath, "/")
			} else if sysType == "windows" {
				tmpStr = strings.Split(ToolScriptLocalPath, "\\")
			}

			finalShell := "cd " + HOST_SCRIPT_PATH + " && " + "cd " + tmpStr[len(tmpStr)-2] + " && " + ToolRunCMD

			ExecuteResult := cliConf.RunShell(finalShell)
			// 4.获取执行结果
			write2logAndBufAndChan(">>> 执行结果"+"\r\n----------------------**********----------------------\r\n"+ExecuteResult+"\r\n----------------------**********----------------------\r\n", &buf, resultChannel, 100, true)
		}
	} else if ToolExecuteLocation == "remote" { //在别的远程环境执行

		write2logAndBufAndChan(">>> 进入远程执行\r\n", &buf, resultChannel, 25, false)

		// // 准备好远程连接的素材
		// cliConf, err := makeSSHClient(buf, resultChannel, SSHPort, SSHIPAddr, SSHAcount, SSHPasswd)
		// if err != nil {
		// 	return err
		// }

		// 2.执行的是容器还是脚本?
		if ToolType == "container" {
			executeContainer(buf, resultChannel, ToolRunCMD, cliConf)
		} else if ToolType == "script" {
			// 执行脚本工具
			write2logAndBufAndChan(">>> 执行的是脚本工具\r\n", &buf, resultChannel, 40, false)

			// 脚本工具，需要先上传脚本到指定位置
			write2logAndBufAndChan(">>> 上传脚本 \r\n"+ToolScriptLocalPath+"\r\n>>> 至指定位置 \r\n"+ToolScriptPath+ToolScriptName+"\r\n", &buf, resultChannel, 60, false)

			uploadResult, err := cliConf.Upload(ToolScriptLocalPath, ToolScriptPath+ToolScriptName)
			if err != nil {
				write2logAndBufAndChan(">>> 上传文件异常:\r\n"+err.Error()+"\r\n", &buf, resultChannel, 100, true)
			}

			write2logAndBufAndChan(">>> 上传结果 \r\n"+uploadResult+"\r\n", &buf, resultChannel, 70, false)

			// 4.开始执行
			write2logAndBufAndChan(">>> 开始执行 \r\n"+ToolRunCMD+"\r\n", &buf, resultChannel, 85, false)

			ExecuteResult := cliConf.RunShell(ToolRunCMD)

			// 5.获取执行结果，后续这里要改进，另起一个goroutine，持续获取执行情况
			write2logAndBufAndChan(">>> 执行结果"+"\r\n----------------------**********----------------------\r\n"+ExecuteResult+"\r\n----------------------**********----------------------\r\n", &buf, resultChannel, 100, true)
		}
	}

	cliConf.sshClient.Close()
	return nil
}

func write2Channel(resultChannel interface{}, progress int, returnContent string, isDone bool) error {
	// 区分管道类型
	switch v := resultChannel.(type) {
	case chan model.Tasks:
		log.Println("是普通任务的管道")
		resultChannel.(chan model.Tasks) <- model.Tasks{Progress: progress, ReturnContent: returnContent, IsDone: isDone}
		return nil

	case chan model.CronTasksResult:
		log.Println("是定时任务的管道")
		resultChannel.(chan model.CronTasksResult) <- model.CronTasksResult{Progress: progress, ReturnContent: returnContent, IsDone: isDone}
		return nil
	default:
		log.Println("未识别出管道类型:", v)
		return errors.New("未识别出管道类型")
	}
}

func CloseMyChannel(resultChannel interface{}) error {
	// 区分管道类型
	switch v := resultChannel.(type) {
	case chan model.Tasks:
		log.Println("准备关闭普通任务的管道")
		close(resultChannel.(chan model.Tasks))
		return nil

	case chan model.CronTasksResult:
		log.Println("关闭定时任务的管道")
		close(resultChannel.(chan model.CronTasksResult))
		return nil
	default:
		log.Println("未识别出管道类型:", v)
		return errors.New("未识别出管道类型")
	}
}

// 复用，执行容器工具
func executeContainer(buf bytes.Buffer, resultChannel interface{}, ToolRunCMD string, cliConf *ClientConfig) {
	write2logAndBufAndChan(">>> 执行的是容器工具\r\n", &buf, resultChannel, 50, false)

	// 3.开始执行
	write2logAndBufAndChan(">>> 开始执行 \r\n"+ToolRunCMD+"\r\n", &buf, resultChannel, 75, false)

	ExecuteResult := cliConf.RunShell(ToolRunCMD)

	// 4.获取执行结果，后续这里要改进，另起一个goroutine，持续获取执行情况
	write2logAndBufAndChan(">>> 执行结果"+"\r\n----------------------**********----------------------\r\n"+ExecuteResult+"\r\n----------------------**********----------------------\r\n", &buf, resultChannel, 100, true)
}

// 复用，新建ssh client
func makeSSHClient(buf bytes.Buffer, resultChannel interface{}, SSHPort string, SSHIPAddr string, SSHAcount string, SSHPasswd string) (*ClientConfig, error) {
	// 准备好远程连接的素材
	port, _ := strconv.Atoi(SSHPort)
	cliConf := new(ClientConfig)
	errCreateClient := cliConf.CreateClient(
		SSHIPAddr,
		int64(port),
		SSHAcount,
		SSHPasswd)

	if errCreateClient != nil {
		write2logAndBufAndChan(">>> 连接异常\r\n"+errCreateClient.Error()+"\r\n", &buf, resultChannel, 100, true)

		CloseMyChannel(resultChannel)
		return nil, errCreateClient
	}
	return cliConf, nil
}

// 复用，log以及buf以及chan的写入
func write2logAndBufAndChan(
	info string,
	buf *bytes.Buffer,
	resultChannel interface{},
	progress int,
	isDone bool,
) {
	log.Print(info)
	buf.WriteString(info)
	write2Channel(resultChannel, progress, buf.String(), isDone)
}
