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

	for i := 1; i <= 5; i++ {
		if sshClient, err = ssh.Dial("tcp", addr, &config); err != nil {
			log.Println(fmt.Printf("第%d次SSH拨号失败: %s, 继续尝试", i, err.Error()))
			time.Sleep(time.Duration(2) * time.Second)
			if i == 5 {
				log.Println("5次ssh尝试全部失败")
				return err
			}

		} else {
			log.Println(fmt.Printf("第%d次SSH拨号成功", i))
			break
		}
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
	log.Println("打开本地文件: ", srcFile)
	dstFile, err := cliConf.sftpClient.Create(dstPath) //远程
	if err != nil {
		log.Printf("在远程创建文件%s失败: "+err.Error(), dstPath)
		return "", err
	}

	defer func() error {
		err1 := srcFile.Close()
		err2 := dstFile.Close()
		if err1 != nil {
			log.Println("关闭本地文本文件失败")
			return err1
		}
		if err2 != nil {
			log.Println("关闭远程文本文件失败")
			return err2
		}
		return nil
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
		_, err = dstFile.Write(buf[:n])
		if err != nil {
			return "", err
		}
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

func (cliConf *ClientConfig) CloseSSHClient() error {
	log.Println("当前cliConf中的sshClient: ", cliConf.sshClient)
	if cliConf.sshClient != nil {
		err := cliConf.sshClient.Close()
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
func (cliConf *ClientConfig) CloseSFTPClient() error {
	log.Println("当前cliConf中的sftpClient: ", cliConf.sftpClient)
	if cliConf.sftpClient != nil {
		err := cliConf.sftpClient.Close()
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (cliConf *ClientConfig) TestCreateSSHConnection(host string, port int64, username, password string) error {
	var (
		sshClient *ssh.Client
		err       error
	)
	cliConf.Host = host
	cliConf.Port = port
	cliConf.Username = username
	cliConf.Password = password

	// 捕捉panic异常
	defer func() {
		if errPanic := recover(); errPanic != nil {
			log.Println("Panic:", errPanic)
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
		cliConf.sshClient.Close()
		log.Println("SSH拨号失败:", err)
		return err
	}

	cliConf.sshClient = sshClient
	err = cliConf.sshClient.Close()
	if err != nil {
		log.Println("关闭SSH连接失败:", err)
		return err
	}
	return nil
}
