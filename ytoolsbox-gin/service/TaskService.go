package service

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"main/dto"
	"main/model"
	"main/util"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

//连接的配置
type ClientConfig struct {
	Host       string       //ip
	Port       int64        // 端口
	Username   string       //用户名
	Password   string       //密码
	sshClient  *ssh.Client  //ssh client
	sftpClient *sftp.Client //sftp client
	LastResult string       //最近一次运行的结果
}

func (cliConf *ClientConfig) createClient(host string, port int64, username, password string) {
	var (
		sshClient  *ssh.Client
		sftpClient *sftp.Client
		err        error
	)
	cliConf.Host = host
	cliConf.Port = port
	cliConf.Username = username
	cliConf.Password = password
	cliConf.Port = port

	config := ssh.ClientConfig{
		Config: ssh.Config{
			Ciphers: []string{"aes256-cbc", "aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc"},
		},
		User: cliConf.Username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 5 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", cliConf.Host, cliConf.Port)

	if sshClient, err = ssh.Dial("tcp", addr, &config); err != nil {
		log.Println("error occurred:", err)
	}
	cliConf.sshClient = sshClient

	//此时获取了sshClient，下面使用sshClient构建sftpClient
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		log.Println("error occurred:", err)
	}
	cliConf.sftpClient = sftpClient
}

func (cliConf *ClientConfig) RunShell(shell string) string {
	var (
		session *ssh.Session
		err     error
	)

	//获取session，这个session是用来远程执行操作的
	if session, err = cliConf.sshClient.NewSession(); err != nil {
		log.Println("error occurred:", err)
		return err.Error()
	}
	//执行shell
	if output, err := session.CombinedOutput(shell); err != nil {
		log.Println(shell)
		log.Println("error occurred:", err, string(output))
		return err.Error() + string(output)
	} else {
		cliConf.LastResult = string(output)
	}
	return cliConf.LastResult
}

func (cliConf *ClientConfig) Upload(srcPath, dstPath string) string {
	srcFile, err := os.Open(srcPath) //本地
	log.Println(srcFile, err)
	dstFile, _ := cliConf.sftpClient.Create(dstPath) //远程
	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()
	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println("error occurred:", err)
				return err.Error()
			} else {
				break
			}
		}
		_, _ = dstFile.Write(buf[:n])
	}
	fmt.Println(cliConf.RunShell(fmt.Sprintf("ls %s", dstPath)))
	return ">>> 上传成功至" + dstPath + "\r\n目标路径存在如下文件:\r\n" + cliConf.RunShell(fmt.Sprintf("ls %s", dstPath))
}

func (cliConf *ClientConfig) Download(srcPath, dstPath string) {
	srcFile, _ := cliConf.sftpClient.Open(srcPath) //远程
	dstFile, _ := os.Create(dstPath)               //本地
	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()

	if _, err := srcFile.WriteTo(dstFile); err != nil {
		log.Println("error occurred", err)
	}
	fmt.Println("文件下载完毕")
}

func CreateNewTaskService(config dto.BriefToolConfigDTO, resultChannel chan model.Tasks) {
	log.Println(config)

	// 结果的缓存
	buf := bytes.Buffer{}
	// 1.在哪里执行?
	if config.ToolExecuteLocation == "local" {
		log.Println(">>> 直接本地执行")
		buf.WriteString(">>> 直接本地执行\r\n")
		resultChannel <- model.Tasks{Progress: 25, ReturnContent: buf.String()}

		// 准备好连接本地的素材
		port, _ := strconv.Atoi(config.ToolRemoteSSH_Port)
		cliConf := new(ClientConfig)
		cliConf.createClient(
			config.ToolRemoteIP,
			int64(port),
			config.ToolRemoteSSH_Account,
			config.ToolRemoteSSH_Password)

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

	} else if config.ToolExecuteLocation == "remote" {
		log.Println(">>> 进入远程执行")
		buf.WriteString(">>> 进入远程执行\r\n")
		resultChannel <- model.Tasks{Progress: 25, ReturnContent: buf.String()}
		// 准备好远程连接的素材
		port, _ := strconv.Atoi(config.ToolRemoteSSH_Port)
		cliConf := new(ClientConfig)
		cliConf.createClient(
			config.ToolRemoteIP,
			int64(port),
			config.ToolRemoteSSH_Account,
			config.ToolRemoteSSH_Password)

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

	// log.Print(util.Struct2MapViaJson(config))
	// //本地文件上传到服务器
	// cliConf.Upload(`D:\settings.txt`, `/tmp/haha.go`) // /root/haha.go
	// //从服务器中下载文件
	// cliConf.Download(`/root/1.py`, `D:\go\1.py`) //文件下载完毕
}

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

	util.CreateDir(FinalFilePath)

	FileDST := filepath.Join(FinalFilePath, fileName)

	//保存文件到服务器本地
	//SaveUploadedFile(文件头，保存路径)
	if err := ctx.SaveUploadedFile(file, FileDST); err != nil {
		return "", "", err
	}
	return toolName, FileDST, nil
}
