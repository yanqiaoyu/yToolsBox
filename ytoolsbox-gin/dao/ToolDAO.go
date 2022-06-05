package dao

import (
	"log"
	"main/dto"
	"main/model"
	"main/utils"

	"gorm.io/gorm"
)

// 删除所有工具
func DeleteAllTools(db *gorm.DB) {
	db.Debug().Unscoped().Where("1 = 1").Delete(&model.Tool{})
	db.Debug().Unscoped().Where("1 = 1").Delete(&model.ToolConfig{})
	db.Debug().Unscoped().Where("\"isDone\" = true").Delete(&model.Tasks{})
}

// 插入工具的基本信息
func InsertNewToolBasicInfo(db *gorm.DB, PostNewToolBasicInfoDTOReq *dto.PostNewToolBasicInfoDTOReq) *gorm.DB {
	result := db.Debug().Model(&model.Tool{}).Create(PostNewToolBasicInfoDTOReq)
	return result
}

// 插入工具的配置信息
func InsertNewToolConfigInfo(db *gorm.DB, PostNewToolConfigInfoDTOReq *dto.PostNewToolConfigInfoDTOReq) *gorm.DB {
	result := db.Model(&model.ToolConfig{}).Create(PostNewToolConfigInfoDTOReq)
	return result
}

// 插入已有工具的配置信息
func InsertExistToolConfigInfo(db *gorm.DB, PostToolConfig *dto.PostNewToolConfigInfoDTOReq) *gorm.DB {
	result := db.Model(&model.ToolConfig{}).Create(PostToolConfig)
	return result
}

// 查询所有的工具
func SelectAllTools(db *gorm.DB, obj dto.GetAllToolsDTOReq) []dto.BriefToolsInfoDTO {
	query := obj.Query
	briefToolsInfoList := []dto.BriefToolsInfoDTO{}

	// 不带Query，返回全部
	// 否则返回like搜索后的结果
	if query == "" {
		// 内联查询
		// 添加了一个where语句的限定条件，是为了去重,这里其实也可以用distinct
		db.Order("tools.id").Model(&model.Tool{}).Where("tool_configs.\"toolConfigName\" = ?", "默认配置").Select("tools.id, tools.\"toolName\", tools.\"toolDesc\", tools.\"toolAuthor\", tools.\"toolRate\", tools.\"toolRateCount\", tool_configs.\"toolType\", tools.\"toolTutorial\"").Joins("join tool_configs on tools.id = tool_configs.\"toolID\"").Find(&briefToolsInfoList)
	} else {
		// toolNames加了双引号，是为了解决pgsql自动变小写的问题
		// 而且这里不能加单引号，否则查询失败
		db.Order("tools.id").Where("tools.\"toolName\" LIKE ?", "%"+query+"%").Model(&model.Tool{}).Select("tools.id, tools.\"toolName\", tools.\"toolDesc\", tools.\"toolAuthor\", tools.\"toolRate\", tools.\"toolRateCount\", tool_configs.\"toolType\", tools.\"toolTutorial\"").Joins("join tool_configs on tools.id = tool_configs.\"toolID\"").Find(&briefToolsInfoList)
	}
	// log.Println(briefToolsInfoList)

	return briefToolsInfoList
}

// 查询特定工具的所有配置
func SelectSpecifiedToolConfig(db *gorm.DB, toolID int, QueryInfo dto.GetSpecifiedToolConfigDTOReqQuery) ([]map[string]interface{}, int) {
	// 配置列表
	configList := []dto.BriefToolConfigDTO{}

	query := QueryInfo.Query
	pagenum := QueryInfo.Pagenum
	pagesize := QueryInfo.Pagesize
	map_configList := []map[string]interface{}{}

	// 不带Query，返回全部
	// 否则返回like搜索后的结果
	if query == "" {
		// 按照时间降序
		db.Order("created_at asc").Model(&model.ToolConfig{}).Where("tool_configs.\"toolID\" = ?", toolID).Find(&configList)
	} else {
		// 按照时间降序
		db.Order("created_at asc").Model(&model.ToolConfig{}).Where("tool_configs.\"toolID\" = ?", toolID).Where("\"toolConfigName\" LIKE ?", "%"+query+"%").Find(&configList)
	}

	DefaultLength := len(configList)

	// 把一个自定义结构体的array 转换成map的array
	// 这里用了json的方法 虽然效率低 但是解决了返回给前端大小写的问题
	for i := 0; i < DefaultLength; i++ {
		map_item := utils.Struct2MapViaJson(configList[i])
		map_configList = append(map_configList, map_item)
	}

	// 计算一下需要如何切割数组
	ArrayStart, ArrayEnd := utils.CalculateReturnMapLength(pagenum, pagesize, map_configList)
	// 返回切片后的结果
	return map_configList[ArrayStart:ArrayEnd], DefaultLength

}

// 根据配置ID查询配置
func SelectSpecifiedToolConfigByConfigID(db *gorm.DB, configID uint) dto.BriefToolConfigDTO {
	config := dto.BriefToolConfigDTO{}
	db.Model(&model.ToolConfig{}).Where("id = ?", configID).Find(&config)
	return config
}

// 更新一下脚本存放在本地的位置 ！！！注意，只有Default条目会存放位置，不管这个工具包含多少个配置，都使用的是第一次上传的脚本文件
func UpdateToolConfigScriptLocalPath(db *gorm.DB, toolName string, FileDST string) {
	tool := model.Tool{}
	db.Debug().Model(&model.Tool{}).Select("id").Where("\"toolName\" = ?", toolName).Find(&tool)
	db.Debug().Model(&model.ToolConfig{}).Where("\"toolID\" = ?", tool.ID).Where("\"toolConfigName\" = ?", "默认配置").Update("toolScriptLocalPath", FileDST)
}

// 删除特定的工具底下的配置条目
func DeleteSpecifiedConfig(db *gorm.DB, configID uint) {
	db.Debug().Delete(&model.ToolConfig{}, configID)
}

// 更新配置信息
func UpdateSpecifiedToolConfigByConfigID(db *gorm.DB, configID uint, obj dto.PutSpecifiedToolConfigByConfigIDDTOReqQuery) {
	log.Println(utils.Struct2MapViaJson(obj))
	db.Debug().Model(&model.ToolConfig{}).Where("id = ?", configID).Updates(
		model.ToolConfig{
			ToolExecuteLocation:    obj.ToolConfig.ToolExecuteLocation,
			ToolConfigName:         obj.ToolConfig.ToolConfigName,
			ToolConfigDesc:         obj.ToolConfig.ToolConfigDesc,
			ToolPythonVersion:      obj.ToolConfig.ToolPythonVersion,
			ToolShellVersion:       obj.ToolConfig.ToolShellVersion,
			ToolDockerImageName:    obj.ToolConfig.ToolDockerImageName,
			ToolOptions:            obj.ToolConfig.ToolOptions,
			ToolRunCMD:             obj.ToolConfig.ToolRunCMD,
			ToolScriptPath:         obj.ToolConfig.ToolScriptPath,
			ToolRemoteIP:           obj.ToolConfig.ToolRemoteIP,
			ToolRemoteSSH_Port:     obj.ToolConfig.ToolRemoteSSH_Port,
			ToolRemoteSSH_Account:  obj.ToolConfig.ToolRemoteSSH_Account,
			ToolRemoteSSH_Password: obj.ToolConfig.ToolRemoteSSH_Password,
		},
	)
}

func UpdateSpecifiedToolTutorialByToolID(db *gorm.DB, PutSpecifiedToolTutorialByToolIDParam dto.PutSpecifiedToolTutorialByToolIDDTOReq) {

	db.Debug().Model(&model.Tool{}).Where("id = ?", PutSpecifiedToolTutorialByToolIDParam.ToolID).Updates(
		model.Tool{
			ToolTutorial: PutSpecifiedToolTutorialByToolIDParam.ToolTutorial,
		},
	)
}
