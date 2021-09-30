package dao

import (
	"log"
	"main/dto"
	"main/model"

	"gorm.io/gorm"
)

func InsertNewTool(db *gorm.DB, PostNewTolReq dto.PostNewToolDTOReq) {
	newTool := model.Tool{}
	log.Println(newTool)
	// result := db.Create(&newTool)

	// return result
}
