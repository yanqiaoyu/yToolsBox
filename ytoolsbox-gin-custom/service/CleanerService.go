package service

import (
	"fmt"
	"log"
	"main/model"
	"main/utils"
	"strconv"

	"github.com/spf13/viper"
)

// 强制DAS审计并生成日志
func DeleteRiskAndVunlService(POCConfig model.POCConfig) (string, error) {
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

	// 找db容器
	dbContainerName := cliConf.RunShell(fmt.Sprintf("echo -n `docker ps --format \"table {{.Names}}\" | grep %s`", viper.GetString("cleanerconfig.dbcontainername")))
	if dbContainerName != "" {
		log.Print("db容器如下: ", dbContainerName)
	} else {
		return "", fmt.Errorf(fmt.Sprintf("目标环境中没有%s容器", viper.GetString("cleanerconfig.dbcontainername")))
	}

	// 找ck容器
	ckContainerName := cliConf.RunShell(fmt.Sprintf("echo -n `docker ps --format \"table {{.Names}}\" | grep %s`", viper.GetString("cleanerconfig.clickhousecontainername")))
	if ckContainerName != "" {
		log.Print("ck容器如下: ", ckContainerName)
	} else {
		return "", fmt.Errorf(fmt.Sprintf("目标环境中没有%s容器", viper.GetString("cleanerconfig.clickhousecontainername")))
	}

	// 找siem容器
	siemContainerName := cliConf.RunShell(fmt.Sprintf("echo -n `docker ps --format \"table {{.Names}}\" | grep %s`", viper.GetString("cleanerconfig.siemcontainername")))
	if siemContainerName != "" {
		log.Print("siem容器如下: ", siemContainerName)
	} else {
		return "", fmt.Errorf(fmt.Sprintf("目标环境中没有%s容器", viper.GetString("cleanerconfig.siemcontainername")))
	}

	// 删除ck
	result1 := cliConf.RunShell(fmt.Sprintf("docker exec %s clickhouse-client --host 0.0.0.0 --port 9876 -d hecate -m -q \"truncate table api_weak\"", ckContainerName))
	result2 := cliConf.RunShell(fmt.Sprintf("docker exec %s clickhouse-client --host 0.0.0.0 --port 9876 -d hecate -m -q \"truncate table risk_log\"", ckContainerName))

	// 删除db
	result3 := cliConf.RunShell(fmt.Sprintf("docker exec %s psql -h 0.0.0.0 -p 5432 -U postgres -d dsc -c \"truncate table hecate.api_weak_count;truncate table hecate.risk_log_attr\"", dbContainerName))

	// 重启siem
	result4 := cliConf.RunShell(fmt.Sprintf("docker exec %s supervisorctl restart api_log", siemContainerName))

	log.Println(result1, result2, result3, result4)

	return "", nil
}
