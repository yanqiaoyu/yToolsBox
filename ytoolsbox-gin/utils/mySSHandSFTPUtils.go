package utils

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/pkg/sftp"
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

// 创建ssh client
func (cliConf *ClientConfig) CreateClient(host string, port int64, username, password string) error {
	var (
		sshClient  *ssh.Client
		sftpClient *sftp.Client
		err        error
	)
	cliConf.Host = host
	cliConf.Port = port
	cliConf.Username = username
	cliConf.Password = password

	// 捕捉panic异常
	defer func() {
		if errPanic := recover(); errPanic != nil {
			log.Println("Panic:", errPanic)
			// return "Something Wrong During Create SSH Connection"
		}
	}()

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
		log.Println("Dial error occurred:", err)
		return err
	}

	cliConf.sshClient = sshClient

	//此时获取了sshClient，下面使用sshClient构建sftpClient
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		log.Println("NewClient error occurred:", err)
		return err
	}
	cliConf.sftpClient = sftpClient

	return nil
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
		log.Println("shellCMD:", shell)
		log.Println("error occurred:", err, string(output))
		return err.Error() + string(output)
	} else {
		cliConf.LastResult = string(output)
	}
	return cliConf.LastResult
}

func (cliConf *ClientConfig) Upload(srcPath, dstPath string) (string, error) {
	srcFile, err := os.Open(srcPath) //本地
	if err != nil {
		return "", err
	}
	// log.Println(srcFile, err)
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
				return "", err
			} else {
				break
			}
		}
		_, _ = dstFile.Write(buf[:n])
	}

	cliConf.sftpClient.Close()
	return ">>> 上传成功至" + dstPath + "\r\n目标路径存在如下文件:\r\n" + cliConf.RunShell(fmt.Sprintf("ls %s", dstPath)), nil
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
