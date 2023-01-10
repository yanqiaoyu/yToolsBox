package dao

import (
	"main/dto"
	"main/model"

	"main/utils"

	"gorm.io/gorm"
)

// func SelectAllUser(db *gorm.DB, query string, pagenum int, pagesize int, param string) ([]map[string]interface{}, int) {
func SelectAllUser(db *gorm.DB, obj dto.GetAllUserDTOReq) ([]map[string]interface{}, int) {
	struct_userList := []dto.UserDTO{}
	map_userList := []map[string]interface{}{}
	query := obj.Query
	pagenum := obj.Pagenum
	pagesize := obj.Pagesize

	// 不带Query，返回全部
	// 否则返回like搜索后的结果
	if query == "" {
		db.Order("id").Model(&model.User{}).Find(&struct_userList)
	} else {
		db.Order("id").Where("username LIKE ?", "%"+query+"%").Model(&model.User{}).Find(&struct_userList)
	}

	DefaultLength := len(struct_userList)

	// 把一个自定义结构体的array 转换成map的array
	// 这里用了json的方法 虽然效率低 但是解决了返回给前端大小写的问题
	for i := 0; i < len(struct_userList); i++ {
		map_item := utils.Struct2MapViaJson(struct_userList[i])
		map_userList = append(map_userList, map_item)
	}

	// 计算一下需要如何切割数组
	ArrayStart, ArrayEnd := utils.CalculateReturnMapLength(pagenum, pagesize, map_userList)
	// 返回切片后的结果
	return map_userList[ArrayStart:ArrayEnd], DefaultLength
}

func SelectSpecifiedUser(db *gorm.DB, userID int64) dto.UserDTO {
	struct_userList := dto.UserDTO{}
	db.Model(&model.User{}).Where("id = ?", userID).Find(&struct_userList)
	return struct_userList
}

func DeleteSpecifiedUser(db *gorm.DB, userID int) model.User {
	struct_user := model.User{}
	db.Delete(&struct_user, userID)
	return struct_user
}

func UpdateUserState(db *gorm.DB, PutUserStateDTOReq dto.PutUserStateDTOReq) {
	db.Model(&model.User{}).Where("id = ?", PutUserStateDTOReq.UserID).Update("mgstate", PutUserStateDTOReq.Mgstate)
}

func UpdateSpecifiedUser(db *gorm.DB, PutUserInfoDTOReq dto.PutUserInfoDTOReq, GetSpecifiedUserDTOReq dto.GetSpecifiedUserDTOReq) {
	db.Model(&model.User{}).Where("id = ?", GetSpecifiedUserDTOReq.UserID).Updates(map[string]interface{}{"email": PutUserInfoDTOReq.Email, "mobile": PutUserInfoDTOReq.Mobile})
}

func InsertNewUser(db *gorm.DB, PostNewUserReq dto.PostNewUserReq) (model.User, *gorm.DB) {
	newUser := model.User{UserName: PostNewUserReq.UserName, PassWord: PostNewUserReq.PassWord, Mobile: PostNewUserReq.Mobile, Email: PostNewUserReq.Email, WorkNum: PostNewUserReq.WorkNum}
	result := db.Create(&newUser)

	return newUser, result
}
