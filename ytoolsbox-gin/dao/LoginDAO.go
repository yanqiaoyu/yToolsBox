package dao

import (
	"main/model"

	"gorm.io/gorm"
)

// 用户是否存在
func IsUserExist(db *gorm.DB, username string) bool {
	var User model.User
	db.Where("username = ?", username).First(&User)
	return User.ID != 0
}
