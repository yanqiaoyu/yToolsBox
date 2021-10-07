package service

import (
	"fmt"
	"io"
	"log"
	"main/dto"
	"main/util"
	"net"
	"os"
	"path/filepath"
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
		User: cliConf.Username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 5 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", cliConf.Host, cliConf.Port)

	if sshClient, err = ssh.Dial("tcp", addr, &config); err != nil {
		log.Fatalln("error occurred:", err)
	}
	cliConf.sshClient = sshClient

	//此时获取了sshClient，下面使用sshClient构建sftpClient
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		log.Fatalln("error occurred:", err)
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
		log.Fatalln("error occurred:", err)
	}
	//执行shell
	if output, err := session.CombinedOutput(shell); err != nil {
		fmt.Println(shell)
		log.Fatalln("error occurred:", err)
	} else {
		cliConf.LastResult = string(output)
	}
	return cliConf.LastResult
}

func (cliConf *ClientConfig) Upload(srcPath, dstPath string) {
	srcFile, _ := os.Open(srcPath)                   //本地
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
				log.Fatalln("error occurred:", err)
			} else {
				break
			}
		}
		_, _ = dstFile.Write(buf[:n])
	}
	fmt.Println(cliConf.RunShell(fmt.Sprintf("ls %s", dstPath)))
}

func (cliConf *ClientConfig) Download(srcPath, dstPath string) {
	srcFile, _ := cliConf.sftpClient.Open(srcPath) //远程
	dstFile, _ := os.Create(dstPath)               //本地
	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()

	if _, err := srcFile.WriteTo(dstFile); err != nil {
		log.Fatalln("error occurred", err)
	}
	fmt.Println("文件下载完毕")
}

func CreateNewTaskService(config dto.BriefToolConfigDTO) {
	cliConf := new(ClientConfig)
	cliConf.createClient("103.44.241.227", 22, "yqy", "")

	// 1.在哪里执行?
	if config.ToolExecuteLocation == "local" {
		log.Println("直接本地执行")
	} else if config.ToolExecuteLocation == "remote" {
		log.Println("进入远程执行")
	}

	// log.Print(util.Struct2MapViaJson(config))
	// //本地文件上传到服务器
	// cliConf.Upload(`D:\settings.txt`, `/tmp/haha.go`) // /root/haha.go
	// //从服务器中下载文件
	// cliConf.Download(`/root/1.py`, `D:\go\1.py`) //文件下载完毕
}

func SaveScriptFile(ctx *gin.Context) (string, error) {
	// 获取文件名
	file, err := ctx.FormFile("file")
	if err != nil {
		return "", err
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
		return "", err
	}
	return FileDST, nil
}
