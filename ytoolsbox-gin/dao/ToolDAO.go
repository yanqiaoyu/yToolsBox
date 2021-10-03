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
		db.Order("tools.id").Model(&model.Tool{}).Select("tools.id, tools.\"toolName\", tools.\"toolDesc\", tools.\"toolAuthor\", tools.\"toolRate\", tools.\"toolRateCount\", tool_configs.\"toolType\"").Joins("join tool_configs on tools.id = tool_configs.\"toolID\"").Find(&briefToolsInfoList)
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
	db.Order("id").Where("tool_configs.\"toolID\" = ?", toolID).Model(&model.ToolConfig{}).Find(&configList)
	// log.Println(configList)
	return configList

}
