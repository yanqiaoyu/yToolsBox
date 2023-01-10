package controller

import (
	"fmt"
	"log"
	"main/common"
	"main/dao"
	"main/dto"
	grpcclient "main/grpc/grpc_client"
	"main/model"
	"main/response"
	"main/service"
	"main/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 获取POC配置
func GetPOCConfig(ctx *gin.Context) {
	db := common.GetDB()
	POCConfig, err := dao.SelectPOCConfig(db)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "获取配置失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
	} else {
		Meta := dto.SuccessResponseMeta{Message: "获取配置成功", StatusCode: 200}
		response.Success(ctx, utils.Struct2MapViaJson(POCConfig), utils.Struct2MapViaJson(Meta))
	}

}

// 保存POC配置
func PostPOCConfig(ctx *gin.Context) {
	db := common.GetDB()

	SavePOCConfigParam := model.POCConfig{}
	if utils.ResolveParam(ctx, &SavePOCConfigParam) != nil {
		return
	}
	log.Println("需要保存的POC配置如下", utils.Struct2MapViaJson(SavePOCConfigParam))

	insertResult := dao.InsertSavePOCConfig(db, SavePOCConfigParam)

	if insertResult != nil {
		Meta := dto.SuccessResponseMeta{Message: "保存配置失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
	} else {
		Meta := dto.SuccessResponseMeta{Message: "保存配置成功", StatusCode: 200}
		response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))
	}

}

// 更新工具盒中的agent
func UpdateDSCAgentInToolBox(ctx *gin.Context) {
	// 先拿一下POC配置
	db := common.GetDB()
	POCConfig, err := dao.SelectPOCConfig(db)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "拿取POC配置失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	// 在本地生成一个配置文件
	filePath, err := service.GenerateLocalAgentConfigFile(POCConfig)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "生成新的Agent配置文件失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	// 然后上传
	// 执行 hecate_agent.sh restart 指令
	// remoteFilePath := "/etc/hecate_agent/das_agent.ini"
	remoteFilePath := viper.GetString("pocconfig.dscAgentConfigPath") + viper.GetString("pocconfig.dscAgentConfigName")
	err = service.UploadLocalAgentConfigFile(filePath, POCConfig, remoteFilePath)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "新的Agent配置文件上传失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	// 删除本地的配置文件
	err = service.RemoveLocalPOCConfigFile(filePath)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "删除本地Agent配置文件失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}
	Meta := dto.SuccessResponseMeta{Message: "更新本地Agent配置文件成功", StatusCode: 200}
	response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))
}

// 更新工具盒中的账号提取
func UpdateDSCAccountExtract(ctx *gin.Context) {
	// 先拿一下POC配置
	db := common.GetDB()
	POCConfig, err := dao.SelectPOCConfig(db)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "拿取POC配置失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	// 在本地生成两个账号提取文件
	authConfPath, appConfPath, err := service.GenerateLocalAccountExtractFile(POCConfig)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "生成新的账号提取配置文件失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	// 然后上传
	// remoteAuthConfFilePath := "/hecate/data/meta-data/account/auth.conf"
	remoteAuthConfFilePath := viper.GetString("pocconfig.accountExtractConfigPath") + viper.GetString("pocconfig.authConfigName")
	err = service.UploadLocalAccountConfigFile(authConfPath, POCConfig, remoteAuthConfFilePath)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "新的auth.conf配置文件上传失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	// remoteAppConfFilePath := "/hecate/data/meta-data/account/app.conf"
	remoteAppConfFilePath := viper.GetString("pocconfig.accountExtractConfigPath") + viper.GetString("pocconfig.appConfigName")
	err = service.UploadLocalAccountConfigFile(appConfPath, POCConfig, remoteAppConfFilePath)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "新的app.conf配置文件上传失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	// 删除本地的配置文件
	err = service.RemoveLocalPOCConfigFile(authConfPath)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "删除本地auth.conf失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	err = service.RemoveLocalPOCConfigFile(appConfPath)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "删除本地app.conf失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	Meta := dto.SuccessResponseMeta{Message: "更新安全大脑账号提取配置文件成功", StatusCode: 200}
	response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))

}

// 测试ssh链接
func TestSSHConnection(ctx *gin.Context) {
	SSHConfig := dto.TestSSHDTO{}
	if utils.ResolveParam(ctx, &SSHConfig) != nil {
		return
	}
	log.Println("需要测试SSH的链接配置为: ", utils.Struct2MapViaJson(SSHConfig))

	cliConf := new(utils.ClientConfig)
	port, _ := strconv.Atoi(SSHConfig.Port)
	err := cliConf.CreateClient(SSHConfig.IP, int64(port), SSHConfig.Username, SSHConfig.Password)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "测试链接失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}
	cliConf.CloseSFTPClient()
	cliConf.CloseSSHClient()

	Meta := dto.SuccessResponseMeta{Message: "测试链接成功", StatusCode: 200}
	response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))
}

// 查看工具盒中Agent的配置
func GetDSCAgentConfig(ctx *gin.Context) {
	db := common.GetDB()
	//查询配置
	POCConfig, err := dao.SelectPOCConfig(db)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "查询Agent的配置失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
	}

	// 根据配置执行语句
	result, err := service.CheckConfigContent(fmt.Sprintf("cat %s", viper.GetString("pocconfig.dscAgentConfigPath")+viper.GetString("pocconfig.dscAgentConfigName")),
		POCConfig.ToolBoxSSHPort,
		POCConfig.ToolBoxAddress,
		POCConfig.ToolBoxSSHUserName,
		POCConfig.ToolBoxSSHPassword)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "查询Agent的配置失败:" + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}
	Meta := dto.SuccessResponseMeta{Message: "查询Agent的配置成功", StatusCode: 200}
	response.Success(ctx, gin.H{"Content": result}, utils.Struct2MapViaJson(Meta))
}

// 查看大脑中账号提取的配置
func GetAccountExtractConfig(ctx *gin.Context) {
	db := common.GetDB()
	//查询配置
	POCConfig, err := dao.SelectPOCConfig(db)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "查询大脑中账号提取的配置失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
	}

	// 根据配置执行语句
	result, err := service.CheckConfigContent(fmt.Sprintf("cat %s %s", viper.GetString("pocconfig.accountExtractConfigPath")+viper.GetString("pocconfig.appConfigName"), viper.GetString("pocconfig.accountExtractConfigPath")+viper.GetString("pocconfig.authConfigName")),
		POCConfig.DSCSSHPort,
		POCConfig.DSCAddress,
		POCConfig.DSCSSHUserName,
		POCConfig.DSCPassword)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "查询大脑中账号提取的配置失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}
	Meta := dto.SuccessResponseMeta{Message: "查询大脑中账号提取的配置成功", StatusCode: 200}
	response.Success(ctx, gin.H{"Content": result}, utils.Struct2MapViaJson(Meta))
}

// 强制审计
func ForceAudit(ctx *gin.Context) {
	db := common.GetDB()
	//查询配置
	POCConfig, err := dao.SelectPOCConfig(db)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "查询大脑中账号提取的配置失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
	}

	_, err = service.ForceDASContainerAuditAndGenerateLog(POCConfig)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "强制审计失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}
	Meta := dto.SuccessResponseMeta{Message: "强制审计成功", StatusCode: 200}
	response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))
}

// 获取Agent安装的配置
func GetAgentInstallConfig(ctx *gin.Context) {
	db := common.GetDB()
	AgentInstallConfig, Defaultlength, err := dao.SelectAgentInstallConfig(db)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "获取配置失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
	} else {
		Meta := dto.SuccessResponseMeta{Message: "获取配置成功", StatusCode: 200}
		response.Success(ctx, gin.H{"Total": Defaultlength, "AgentInstallConfigList": AgentInstallConfig}, utils.Struct2MapViaJson(Meta))
	}

}

// 保存Agent安装的配置
func PostAgentInstallConfig(ctx *gin.Context) {
	db := common.GetDB()

	SaveAgentInstallConfigParam := model.AgentInstallConfig{}
	if utils.ResolveParam(ctx, &SaveAgentInstallConfigParam) != nil {
		return
	}
	log.Println("需要保存的Agent安装的配置如下", utils.Struct2MapViaJson(SaveAgentInstallConfigParam))

	insertResult := dao.InsertAgentInstallConfig(db, SaveAgentInstallConfigParam)

	if insertResult != nil {
		Meta := dto.SuccessResponseMeta{Message: "保存配置失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
	} else {
		Meta := dto.SuccessResponseMeta{Message: "保存配置成功", StatusCode: 200}
		response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))
	}
}

// 执行安装agent的操作
func PostInstallAgent(ctx *gin.Context) {
	AgentInstallConfigParam := model.AgentInstallConfig{}
	if utils.ResolveParam(ctx, &AgentInstallConfigParam) != nil {
		return
	}

	// 上传安装包至特定位置, 并安装
	_, err := service.UploadAndInstallAgentPackage(AgentInstallConfigParam)
	if err != nil {
		Meta := dto.FailResponseMeta{Message: "安装Agent失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	Meta := dto.SuccessResponseMeta{Message: "安装Agent成功", StatusCode: 200}
	response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))

}

// 调整大脑的阈值
func ModifyDSCThreshold(ctx *gin.Context) {
	db := common.GetDB()

	mode := ctx.Request.FormValue("mode")
	log.Println("调整阈值的模式为: ", mode)

	// 从数据库中,拿到大脑的IP,前端账号,前端密码
	result, err := dao.SelectModifyDSCThresholdConfig(db)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "大脑配置查询失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	log.Print("需要调整阈值的大脑IP: ", result.DSCAddress)
	log.Print("需要调整阈值的大脑前端账号: ", result.DSCWebUserName)
	log.Print("需要调整阈值的大脑前端密码: ", result.DSCWebPassword)

	if mode == "poc" {
		err := grpcclient.ModifyThreshold(mode, result)
		if err != nil {
			Meta := dto.SuccessResponseMeta{Message: "调整阈值失败", StatusCode: 401}
			response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
			return
		}

		Meta := dto.SuccessResponseMeta{Message: "调整阈值成功", StatusCode: 200}
		response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	} else if mode == "default" {
		err := grpcclient.ModifyThreshold(mode, result)
		if err != nil {
			Meta := dto.SuccessResponseMeta{Message: "调整阈值失败", StatusCode: 401}
			response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
			return
		}

		Meta := dto.SuccessResponseMeta{Message: "还原阈值成功", StatusCode: 200}
		response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}
}
