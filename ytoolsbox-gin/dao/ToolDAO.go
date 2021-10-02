package dao

import (
	"main/dto"
	"main/model"

	"gorm.io/gorm"
)

func InsertNewTool(db *gorm.DB, PostNewToolReq dto.PostNewToolDTOReq) *gorm.DB {
	result := db.Model(&model.Tool{}).Create(&PostNewToolReq)
	return result
}

func SelectAllTools(db *gorm.DB, obj dto.GetAllToolsDTOReq) []dto.ToolsDTO {
	query := obj.Query
	tools_List := []dto.ToolsDTO{}

	// 不带Query，返回全部
	// 否则返回like搜索后的结果
	if query == "" {
		db.Order("id").Model(&model.Tool{}).Find(&tools_List)
	} else {
		// toolNames加了双引号，是为了解决pgsql自动变小写的问题
		// 而且这里不能加单引号，否则查询失败
		db.Order("id").Where("\"toolName\" LIKE ?", "%"+query+"%").Model(&model.Tool{}).Find(&tools_List)
	}

	return tools_List
}
