package dao

import (
	"main/dto"
	"main/model"

	"gorm.io/gorm"
)

// 插入工具的基本信息
func InsertNewToolBasicInfo(db *gorm.DB, PostNewToolBasicInfoDTOReq *dto.PostNewToolBasicInfoDTOReq) *gorm.DB {
	result := db.Model(&model.Tool{}).Create(PostNewToolBasicInfoDTOReq)
	return result
}

// 插入工具的配置信息
func InsertNewToolConfigInfo(db *gorm.DB, PostNewToolConfigInfoDTOReq *dto.PostNewToolConfigInfoDTOReq) *gorm.DB {
	result := db.Model(&model.ToolConfig{}).Create(PostNewToolConfigInfoDTOReq)
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
		db.Order("tools.id").Model(&model.Tool{}).Where("tool_configs.\"toolConfigName\" = ?", "默认配置").Select("tools.id, tools.\"toolName\", tools.\"toolDesc\", tools.\"toolAuthor\", tools.\"toolRate\", tools.\"toolRateCount\", tool_configs.\"toolType\"").Joins("join tool_configs on tools.id = tool_configs.\"toolID\"").Find(&briefToolsInfoList)
	} else {
		// toolNames加了双引号，是为了解决pgsql自动变小写的问题
		// 而且这里不能加单引号，否则查询失败
		db.Order("tools.id").Where("tools.\"toolName\" LIKE ?", "%"+query+"%").Model(&model.Tool{}).Select("tools.id, tools.\"toolName\", tools.\"toolDesc\", tools.\"toolAuthor\", tools.\"toolRate\", tools.\"toolRateCount\", tool_configs.\"toolType\"").Joins("join tool_configs on tools.id = tool_configs.\"toolID\"").Find(&briefToolsInfoList)
	}
	// log.Println(briefToolsInfoList)

	return briefToolsInfoList
}

// 查询特定工具的配置
func SelectSpecifiedToolConfig(db *gorm.DB, toolID int) []dto.BriefToolConfigDTO {
	// 配置列表
	configList := []dto.BriefToolConfigDTO{}

	// 1.拿取配置
	// db.Model(&model.Tool{}).Select("tools.id, tools.\"toolName\", tool_configs.\"toolConfigName\", tool_configs.\"toolConfigDesc\"").Joins("join tool_configs on tools.id = tool_configs.\"toolID\"").Where("tool_configs.\"toolID\" = ?", toolID).Find(&configList)
	// log.Println(configList)
	db.Model(&model.ToolConfig{}).Where("tool_configs.\"toolID\" = ?", toolID).Find(&configList)
	return configList
}

// 更新一下脚本存放在本地的位置 ！！！注意，只有Default条目会存放位置，不管这个工具包含多少个配置，都使用的是第一次上传的脚本文件
func UpdateToolConfigScriptLocalPath(db *gorm.DB, toolName string, FileDST string) {
	tool := model.Tool{}
	db.Debug().Model(&model.Tool{}).Select("id").Where("\"toolName\" = ?", toolName).Find(&tool)
	db.Debug().Model(&model.ToolConfig{}).Where("\"toolID\" = ?", tool.ID).Where("\"toolConfigName\" = ?", "默认配置").Update("toolScriptLocalPath", FileDST)
}
