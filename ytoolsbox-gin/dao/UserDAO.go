package dao

import (
	"log"
	"main/dto"
	"main/model"
	"main/service"
	"main/util"

	"gorm.io/gorm"
)

func SelectAllUser(db *gorm.DB, query string, pagenum int, pagesize int, param string) []map[string]interface{} {
	struct_userList := []dto.UserDTO{}
	map_userList := []map[string]interface{}{}
	db.Model(&model.User{}).Find(&struct_userList)

	// 把一个自定义结构体的array 转换成map的array
	// 这里用了json的方法 虽然效率低 但是解决了返回给前端大小写的问题
	for i := 0; i < len(struct_userList); i++ {
		map_item := util.Struct2MapViaJson(struct_userList[i])
		log.Println(map_item)
		map_userList = append(map_userList, map_item)
	}

	// 计算一下需要如何切割数组
	ArrayStart, ArrayEnd := service.CalculateReturnMapLength(pagenum, pagesize, map_userList)
	// 返回切片后的结果
	return map_userList[ArrayStart:ArrayEnd]
}
