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

func SelectAllTools(db *gorm.DB) []dto.ToolsDTO {
	tools_List := []dto.ToolsDTO{}
	db.Model(&model.Tool{}).Find(&tools_List)
	// log.Print(tools_List)
	return tools_List
}
