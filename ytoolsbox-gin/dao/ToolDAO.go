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
