package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/common"
	"main/dao"
	"main/model"
	"main/utils"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var InnerSourceIPList = []string{
	// sz
	"119.132.243.144",
	// dongguan
	"121.12.145.25",
	// sh
	"121.5.26.3",
	// yun fu
	"121.10.25.74",
	// tang shan
	"121.20.25.74",
	// zhong guo
	"121.5.26.3",
	// shang hai
	"121.5.26.3",
}

var ForiegnSourceIPList = []string{
	"172.15.36.25",
	"25.110.36.25",
	"27.110.36.25",
	"94.77.2.12",
	"92.2.62.45",
	"1.6.224.28",
	"95.3.63.84",
	"122.1.16.79",
}

func GenerateLocalAgentConfigFile(POCConfig model.POCConfig) (string, error) {
	// 先根据工具盒的IP反查一下网卡名称
	networkCardName, err := GetNetworkCardNameByIP(POCConfig)
	if err != nil || networkCardName == "" {
		networkCardName = "eth0"
	}

	fileContent := fmt.Sprintf(`
# 必须保证是unix格式!
[server]
# 数据包发送地址
url               = /api/input/pcap?token=%s
host              = %s
port              = 4430
# http:0, https:1
proto             = 1
# 心跳包发送地址
heartbeat_url     = https://%s:4430/api/devices/agents/heartbeat
token             = %s
# 获取配置更新的发送地址
config_update_url = https://%s:4430/api/devices/config/dp/agent

[capture]
dev           = %s
filter        = port 443 or port 80 or port 3306 or port 1521 
prod_cons_cnt = 150000
# 是否保存文件，0不保存，1保存
save_file     = 0
# 0:hash ip, 1: hash tcpseq
hash_alg      = 0
# 1: cap by libpcap, 0:raw socket
cap_type      = 1
# pcap ring gbuffer size: default 2MB
pcap_buffer   = 64
# 如果IP显示错误(例如NAT场景，如果配置下面的client_ip)
client_ip     = 10.87.68.16

[sender]
cache             = 0
conn_cnt          = 1
# 是否发送心跳包，0不发送，1发送
send_heartbeat    = 1
# 心跳间隔时间，每heart_rate秒发送一次心跳包
heart_rate        = 60
# max value:60000
sendslot_per_once = 40000
thread_num        = 4
send_sleep        = 200000
ssl               = ECDHE-RSA-AES128-SHA
keepalive_max_req = 40000
wait_response     = 0
# 0:normal; 1:send none, only capture
send_debug        = 0
# system read/write socket buffer, MB
socket_buffer     = 64
read_timeout      = 150000

[limit]
# 限制cpu使用率不能超过百分之80
cpu                   = 160
# dump packet counter to file(60 seconds)
dump_counter_interval = 60
	`, POCConfig.DP_Token, POCConfig.DSCAddress, POCConfig.DSCAddress, POCConfig.DP_Token, POCConfig.DSCAddress, networkCardName)

	log.Println("Agent配置内容", fileContent)

	byteFileContent := []byte(fileContent)
	filePath := "./" + viper.GetString("pocconfig.dscAgentConfigName")

	err = ioutil.WriteFile(filePath, byteFileContent, 0666) //写入文件(字节数组)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func UploadLocalAgentConfigFile(filePath string, POCConfig model.POCConfig, RemoteFilePath string) error {
	// 准备好远程连接的素材
	port, _ := strconv.Atoi(POCConfig.ToolBoxSSHPort)
	cliConf := new(utils.ClientConfig)
	errCreateClient := cliConf.CreateClient(
		POCConfig.ToolBoxAddress,
		int64(port),
		POCConfig.ToolBoxSSHUserName,
		POCConfig.ToolBoxSSHPassword)
	if errCreateClient != nil {
		return errCreateClient
	}

	defer func() {
		log.Print("关闭SFTP与SSH连接")
		cliConf.CloseSFTPClient()
		cliConf.CloseSSHClient()
	}()

	_, err := cliConf.Upload(filePath, RemoteFilePath)
	if err != nil {
		return err
	}

	log.Println("准备重启hecate_agent")
	cliConf.RunShell(viper.GetString("pocconfig.restartDscAgentShell"))

	cliConf.CloseSSHClient()

	return nil
}

func UploadLocalAccountConfigFile(filePath string, POCConfig model.POCConfig, RemoteFilePath string) error {
	// 准备好远程连接的素材
	port, _ := strconv.Atoi(POCConfig.DSCSSHPort)
	cliConf := new(utils.ClientConfig)
	errCreateClient := cliConf.CreateClient(
		POCConfig.DSCAddress,
		int64(port),
		POCConfig.DSCSSHUserName,
		POCConfig.DSCPassword)
	if errCreateClient != nil {
		return errCreateClient
	}

	defer func() {
		log.Print("关闭SFTP与SSH连接")
		cliConf.CloseSFTPClient()
		cliConf.CloseSSHClient()
	}()

	_, err := cliConf.Upload(filePath, RemoteFilePath)
	if err != nil {
		return err
	}

	// 上传第二个文件之后, 重启api rich容器
	if filePath == "./"+viper.GetString("pocconfig.appConfigName") {
		result := cliConf.RunShell(fmt.Sprintf("echo -n `docker ps | grep %s | cut -d ' ' -f 1`", viper.GetString("pocconfig.accountExtractContainerName")))
		if result == "" {
			return fmt.Errorf("没有api rich容器")
		}
		log.Println(viper.GetString("pocconfig.accountExtractContainerName"), "容器ID为: ", result)
		result = cliConf.RunShell(fmt.Sprintf("docker exec %s supervisorctl restart %s", result, viper.GetString("pocconfig.accountExtractContainerName")))
		log.Println("重启"+viper.GetString("pocconfig.accountExtractContainerName")+"容器结果: ", result)
	}

	return nil
}

func RemoveLocalPOCConfigFile(filePath string) error {
	// 先判断文件是否存在
	exist := utils.CheckFileIsExist(filePath)
	if exist {
		log.Println("本地存在agent配置文件,准备删除")
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	} else {
		log.Println("本地无agent配置文件,不执行任何动作")
		return nil
	}

	return nil
}

func GenerateLocalAccountExtractFile(POCConfig model.POCConfig) (string, string, error) {
	authConfContent := fmt.Sprintf(`
	{
		"version": "1.0.0",
		"config": [
			{
				"id": "test_auth_001",
				"relation_app_ids": [
					"my_tool_box"
				],
				"type": "token",
				"url": "regex:%s/api/auth/custom/mock/vulnerability/WeakPasswd",
				"field_extract_conf": {
					"account": [
						{
							"pos": "req_body_orig",
							"operator": [
								"www-form-urldecode",
								"eq:name"
							]
						}
					],
					"password": [
						{
							"pos": "req_body_orig",
							"operator": [
								"www-form-urldecode",
								"eq:password"
							]
						}
					],
					"app_token": [
						{
							"pos": "res_body",
							"operator": [
								"eq:token"
							]
						}
					]
				}
			},
			{
				"id": "test_auth_002",
				"relation_app_ids": [
					"my_tool_box2"
				],
				"type": "token",
				"url": "regex:10.87.68.19/api/auth/custom/mock/risk/LocalIPWithMultiAccount",
				"field_extract_conf": {
					"account": [
						{
							"pos": "req_body_orig",
							"operator": [
								"www-form-urldecode",
								"eq:name"
							]
						}
					],
					"password": [
						{
							"pos": "req_body_orig",
							"operator": [
								"www-form-urldecode",
								"eq:password"
							]
						}
					],
					"app_token": [
						{
							"pos": "res_body",
							"operator": [
								"eq:token"
							]
						}
					]
				}
			}
		]
	}
`, POCConfig.ToolBoxAddress)

	appConfContent := fmt.Sprintf(`
	{
		"version": "1.0.0",
		"config": [
			{
				"id": "my_tool_box",
				"auth_id": "test_auth_001",
				"url": "regex:%s/*",
				"field_extract_conf": {
					"app_token": [
						{
							"pos": "req_header:authorization",
							"operator": [
								"regex:(?<=Bearer ).*"
							]
						}
					]
				}
			},
			{
				"id": "my_tool_box2",
				"auth_id": "test_auth_002",
				"url": "regex:10.87.68.19/*",
				"field_extract_conf": {
					"app_token": [
						{
							"pos": "req_header:authorization",
							"operator": [
								"regex:(?<=Bearer ).*"
							]
						}
					]
				}
			}
		]
	}
`, POCConfig.ToolBoxAddress)

	log.Println("auth_Conf内容", authConfContent)
	log.Println("app_Conf内容", appConfContent)

	byteAuthConfContent := []byte(authConfContent)
	byteAppConfContent := []byte(appConfContent)

	authConfPath := "./" + viper.GetString("pocconfig.authConfigName")
	appConfPath := "./" + viper.GetString("pocconfig.appConfigName")

	err := ioutil.WriteFile(authConfPath, byteAuthConfContent, 0666) //写入文件(字节数组)
	if err != nil {
		return "", "", err
	}
	err = ioutil.WriteFile(appConfPath, byteAppConfContent, 0666) //写入文件(字节数组)
	if err != nil {
		return "", "", err
	}

	return authConfPath, appConfPath, nil
}

func CheckConfigContent(shell string, string_port string, ipaddr string, username string, pwd string) (string, error) {
	// 准备好远程连接的素材
	port, _ := strconv.Atoi(string_port)
	cliConf := new(utils.ClientConfig)
	errCreateClient := cliConf.CreateClient(
		ipaddr,
		int64(port),
		username,
		pwd)
	if errCreateClient != nil {
		return "", errCreateClient
	}

	result := cliConf.RunShell(shell)

	cliConf.CloseSSHClient()

	return result, nil
}

// 强制DAS审计并生成日志
func ForceDASContainerAuditAndGenerateLog(POCConfig model.POCConfig) (string, error) {
	// 准备好远程连接的素材

	port, _ := strconv.Atoi(POCConfig.DSCSSHPort)
	cliConf := new(utils.ClientConfig)
	errCreateClient := cliConf.CreateClient(
		POCConfig.DSCAddress,
		int64(port),
		POCConfig.DSCSSHUserName,
		POCConfig.DSCPassword)
	if errCreateClient != nil {
		return "", errCreateClient
	}

	defer func() {
		log.Print("关闭SFTP与SSH连接")
		cliConf.CloseSFTPClient()
		cliConf.CloseSSHClient()
	}()

	// 拿到所有名称里面含有DAS的容器ID
	result := cliConf.RunShell(fmt.Sprintf("echo -n `docker ps  | grep %s: | cut -d ' ' -f 1`", viper.GetString("pocconfig.dscDasContainerName")))
	if result != "" {
		log.Print("DAS容器ID如下: ", result)
		resultList := strings.SplitAfter(result, " ")
		for cnt := 0; cnt < len(resultList); cnt++ {
			fmt.Println("\n强制DAS容器: ", result[cnt], " 进行日志审计")
			cliConf.RunShell(fmt.Sprintf("docker exec %s python3 /usr/bin/flush_log.py", resultList[cnt]))
		}
	} else {
		return "", fmt.Errorf(fmt.Sprintf("目标环境中没有%s容器", viper.GetString("pocconfig.dscDasContainerName")))
	}

	return "", nil
}

// 上传agent install的安装包 并安装
func UploadAndInstallAgentPackage(AgentInstallConfigParam model.AgentInstallConfig) (string, error) {

	// 准备好远程连接的素材
	port, _ := strconv.Atoi(AgentInstallConfigParam.Port)
	cliConf := new(utils.ClientConfig)
	errCreateClient := cliConf.CreateClient(
		AgentInstallConfigParam.IP,
		int64(port),
		AgentInstallConfigParam.Username,
		AgentInstallConfigParam.Password)
	if errCreateClient != nil {
		return "", errCreateClient
	}

	defer func() {
		log.Print("关闭SFTP与SSH连接")
		cliConf.CloseSFTPClient()
		cliConf.CloseSSHClient()
	}()

	// 准备好各种路径
	dscAgentInstallPath := viper.GetString("agentinstall.dscAgentInstallPath")
	dscAgentPackageName := viper.GetString("agentinstall.dscAgentPackageName")

	dasAgentInstallPath := viper.GetString("agentinstall.dasAgentInstallPath")
	dasAgentPackageName := viper.GetString("agentinstall.dasAgentPackageName")

	var InstallResult string

	// 新建目录,上传文件
	if AgentInstallConfigParam.Type == "dsc" {
		cliConf.RunShell("mkdir -p " + dscAgentInstallPath)
		log.Println("新建目录: ", dscAgentInstallPath)
		_, err := cliConf.Upload("./"+dscAgentPackageName, dscAgentInstallPath+dscAgentPackageName)
		if err != nil {
			return "", err
		}
		// 执行安装
		InstallResult = cliConf.RunShell("rpm -ivh " + dscAgentInstallPath + dscAgentPackageName)

	} else if AgentInstallConfigParam.Type == "das" {
		cliConf.RunShell("mkdir -p " + dasAgentInstallPath)
		log.Println("新建目录: ", dasAgentInstallPath)
		_, err := cliConf.Upload("./"+dasAgentPackageName, dasAgentInstallPath+dasAgentPackageName)
		if err != nil {
			return "", err
		}
		// 执行安装
		InstallResult = cliConf.RunShell("cd  " + dasAgentInstallPath + " | " + "chmod 777 " + dasAgentPackageName + " | " + " ./" + dasAgentPackageName + " `pwd`")
	}

	return InstallResult, nil
}

// 回放包
func ReplayPcap(sourceIP string, netcardName string, POCConfig model.POCConfig, originPcap string, cachPcap string, tempPcap string) error {

	// 准备好远程连接的素材
	port, _ := strconv.Atoi(POCConfig.ToolBoxSSHPort)
	cliConf := new(utils.ClientConfig)
	errCreateClient := cliConf.CreateClient(
		POCConfig.ToolBoxAddress,
		int64(port),
		POCConfig.ToolBoxSSHUserName,
		POCConfig.ToolBoxSSHPassword)
	if errCreateClient != nil {
		return errCreateClient
	}

	// 修改源IP
	shellModifySourceIP := fmt.Sprintf("docker exec %s tcprewrite --fixcsum --endpoints=%s:%s -i %s%s -o %s%s -c %s%s",
		viper.GetString("tcpreplayconfig.replaycontainername"),
		sourceIP,
		POCConfig.ToolBoxAddress,
		viper.GetString("tcpreplayconfig.pcappath"),
		originPcap,
		viper.GetString("tcpreplayconfig.pcappath"),
		tempPcap,
		viper.GetString("tcpreplayconfig.pcappath"),
		cachPcap,
	)

	cliConf.RunShell(shellModifySourceIP)
	log.Println("修改源IP: ", shellModifySourceIP)

	// 执行回放动作
	shellExecReplay := fmt.Sprintf("docker exec %s tcpreplay -i %s %s%s",
		viper.GetString("tcpreplayconfig.replaycontainername"),
		netcardName,
		viper.GetString("tcpreplayconfig.pcappath"),
		// viper.GetString("tcpreplayconfig.multiaccounttemp"),
		tempPcap,
	)
	result := cliConf.RunShell(shellExecReplay)
	log.Println("回放包结果: ", result)

	cliConf.CloseSSHClient()
	return nil
}

// 根据IP反查网卡名称
func GetNetworkCardNameByIP(POCConfig model.POCConfig) (string, error) {
	// 准备好远程连接的素材
	port, _ := strconv.Atoi(POCConfig.ToolBoxSSHPort)
	cliConf := new(utils.ClientConfig)
	errCreateClient := cliConf.CreateClient(
		POCConfig.ToolBoxAddress,
		int64(port),
		POCConfig.ToolBoxSSHUserName,
		POCConfig.ToolBoxSSHPassword)
	if errCreateClient != nil {
		return "", errCreateClient
	}

	shell := fmt.Sprintf("echo -n `ip addr | grep -B 2 %s | head -n 1 | awk -F: '{ print $2 }' | tr -d [:blank:]`", POCConfig.ToolBoxAddress)
	netcardName := cliConf.RunShell(shell)

	log.Println("根据IP: ", POCConfig.ToolBoxAddress, " 反查出的网卡名称: ", netcardName)

	cliConf.CloseSSHClient()

	return netcardName, nil
}

// 回放包
func ReplayMultiAccount(IPList []string) error {
	db := common.GetDB()
	POCConfig, err := dao.SelectPOCConfig(db)
	if err != nil {
		return fmt.Errorf("获取配置失败")
	}

	// 获取网卡名称
	networkCardName, err := GetNetworkCardNameByIP(POCConfig)
	if err != nil {
		return fmt.Errorf("获取网卡名称失败")
	}

	// 回放包
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		err := ReplayPcap(
			IPList[rand.Intn(len(IPList)-1)],
			networkCardName,
			POCConfig,
			viper.GetString("tcpreplayconfig.multiaccountpcap"),
			viper.GetString("tcpreplayconfig.multiaccountcach"),
			viper.GetString("tcpreplayconfig.multiaccounttemp"),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// 修改大脑时间
func ModifyDate(POCConfig model.POCConfig, date string) error {
	// 准备好远程连接的素材
	port, _ := strconv.Atoi(POCConfig.DSCSSHPort)
	cliConf := new(utils.ClientConfig)
	errCreateClient := cliConf.CreateClient(
		POCConfig.DSCAddress,
		int64(port),
		POCConfig.DSCSSHUserName,
		POCConfig.DSCPassword)
	if errCreateClient != nil {
		return errCreateClient
	}

	defer func() {
		log.Print("关闭SFTP与SSH连接")
		cliConf.CloseSFTPClient()
		cliConf.CloseSSHClient()
	}()

	shell := fmt.Sprintf("date -s \"%s\"", date)
	shell += ` +"%Y-%m-%d"`
	log.Println("修改时间的语句为:", shell)
	modifyResult := cliConf.RunShell(shell)
	log.Println("修改的结果为:", modifyResult)

	// log.Println("根据IP: ", POCConfig.ToolBoxAddress, " 反查出的网卡名称: ", netcardName)

	return nil
}
