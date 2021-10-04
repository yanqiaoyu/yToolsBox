package dao

import (
	"main/dto"
	"main/model"

	"gorm.io/gorm"
)

// 用户是否存在
func IsUserExist(db *gorm.DB, loginParam dto.LoginDTO) bool {
	User := model.User{}
	db.Where("username = ?", loginParam.UserName).First(&User)
	return User.ID != 0
}

// 检查账号密码是否正确
func CheckPassword(db *gorm.DB, loginParam dto.LoginDTO) bool {
	User := model.User{}
	db.Where("username = ?", loginParam.UserName).First(&User)
	return User.PassWord == loginParam.Password
}
