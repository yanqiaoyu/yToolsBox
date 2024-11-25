package dspm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"main/common"
	"main/cons"
	"main/dao"
	"net"
	"strconv"
	"time"
)

func CreateSSHClient(host string, port int64, username, password string) (*ssh.Client, error) {

	config := ssh.ClientConfig{
		Config: ssh.Config{
			Ciphers: []string{"aes256-cbc", "aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc"},
		},
		User: username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 5 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	var err error
	var sshClient *ssh.Client
	for i := 1; i <= 5; i++ {
		if sshClient, err = ssh.Dial("tcp", addr, &config); err != nil {
			log.Println(fmt.Printf("第%d次SSH拨号失败: %s, 继续尝试", i, err.Error()))
			time.Sleep(time.Duration(2) * time.Second)
			if i == 5 {
				log.Println("5次ssh尝试全部失败")
				return sshClient, err
			}
		} else {
			log.Println(fmt.Printf("第%d次SSH拨号成功", i))
			return sshClient, nil
		}
	}
	return sshClient, errors.New("5次ssh尝试全部失败")
}

func UpdateDspmConfig() {
	db := common.GetDB()
	pocConfig, err := dao.SelectPOCConfig(db)
	if err != nil {
		log.Printf("获取poc的token失败:%v\n", err)
		return
	}
	sshPort, err := strconv.Atoi(pocConfig.DSCSSHPort)
	if err != nil {
		log.Printf("ssh port format error:%v\n", err)
		return
	}
	client, err := CreateSSHClient(pocConfig.DSCAddress, int64(sshPort), pocConfig.DSCSSHUserName, pocConfig.DSCPassword)
	if err != nil {
		return
	}
	err = OpenPort(client)
	if err != nil {
		log.Printf("暴露本地端口失败 error:%v\n", err)
	}
	err = UpdatePocMode(client)
	if err != nil {
		log.Printf("暴露本地端口失败 error:%v\n", err)
	}
	config, err := GetPwdConfig(client)
	if err != nil {
		log.Printf("获取远程密码信息失败 error:%v\n", err)
	}

	cons.DspmAddr = pocConfig.DSCAddress
	cons.PocAddr = pocConfig.ToolBoxAddress
	log.Printf("DspmAddr:%v;PocAddr:%v;config:%v\n", cons.DspmAddr, cons.PocAddr, config)
	json.Unmarshal([]byte(config), &cons.DspmPwdConfig)
}

// GetPwdConfig 通过ssh获取所有密码配置
func GetPwdConfig(client *ssh.Client) (string, error) {
	command := "kubectl -n xdr-comm-config get secrets root-account -o jsonpath='{.data.account}' | base64 -d"

	// 创建会话
	session, err := client.NewSession()
	if err != nil {
		log.Printf("无法创建SSH会话: %v\n", err)
		return "", err
	}
	defer session.Close()

	// 执行命令并获取输出
	var stdoutBuf, stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf

	if err := session.Run(command); err != nil {
		log.Println("命令执行失败")
		return "", err
	}
	return stdoutBuf.String(), nil
}

func OpenPort(client *ssh.Client) error {
	patchPortCmd := "kubectl patch globalnetworkpolicy allow-external-ingress --type=json -p=\"[{\\\"op\\\": \\\"add\\\", \\\"path\\\": \\\"/spec/ingress/0/destination/ports/-\\\", \\\"value\\\": %d}]\""
	cmdSlice := []string{
		fmt.Sprintf("kubectl patch svc postgresql-stolon-proxy  -p '{\"spec\": {\"type\": \"NodePort\", \"ports\": [{\"port\": 5432, \"nodePort\": 30432}]}}' -n postgresql"),
		fmt.Sprintf("kubectl patch svc dsp-ueba-engine-dsp-ueba-engine-flink-jobmanager -p '{\"spec\": {\"type\": \"NodePort\", \"ports\": [{\"port\": 8081, \"nodePort\": 30081}]}}' -n dsp"),
		fmt.Sprintf("kubectl patch svc cluster-pro-xdr-ck  -p '{\"spec\": {\"type\": \"NodePort\", \"ports\": [{\"port\": 8123, \"nodePort\": 30123},{\"port\": 9000, \"nodePort\": 30900}]}}' -n ck"),
		fmt.Sprintf("kubectl patch svc pulsar-proxy  -p '{\"spec\": {\"type\": \"NodePort\", \"ports\": [{\"port\": 80, \"nodePort\": 30080},{\"port\": 6650, \"nodePort\": 30650}]}}' -n pulsar"),
		fmt.Sprintf("kubectl patch svc tenant-10001001-rest -p '{\"spec\": {\"type\": \"NodePort\", \"ports\": [{\"port\": 8081, \"nodePort\": 30181}]}}' -n dsp"),
		fmt.Sprintf(patchPortCmd, 30432),
		fmt.Sprintf(patchPortCmd, 30081),
		fmt.Sprintf(patchPortCmd, 30900),
		fmt.Sprintf(patchPortCmd, 30123),
		fmt.Sprintf(patchPortCmd, 30650),
		fmt.Sprintf(patchPortCmd, 30080),
		fmt.Sprintf(patchPortCmd, 30181),
	}

	var resultError error
	for _, command := range cmdSlice {
		// 创建会话
		session, err := client.NewSession()
		if err != nil {
			log.Printf("无法创建SSH会话: %v\n", err)
			return err
		}
		if err = session.Run(command); err != nil {
			log.Printf("命令执行失败:%v\n", err)
			resultError = err
		}
		session.Close()
	}
	return resultError
}

func UpdatePocMode(client *ssh.Client) error {
	// 创建会话
	session, err := client.NewSession()
	if err != nil {
		log.Printf("无法创建SSH会话: %v\n", err)
		return err
	}
	defer session.Close()
	command := "touch /dsc/data/meta-data/rare/rare_config/config/debug"
	if err = session.Run(command); err != nil {
		log.Printf("命令执行失败:%v\n", err)
	}
	return err
}
