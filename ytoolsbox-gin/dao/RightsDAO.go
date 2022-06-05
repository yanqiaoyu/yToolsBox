package dao

import (
	"main/dto"
	"main/model"
	"main/utils"

	"gorm.io/gorm"
)

func SelectAllRights(db *gorm.DB) []map[string]interface{} {
	struct_RightsList := []dto.RightsListDTO{}
	map_RightsList := []map[string]interface{}{}

	db.Order("id").Model(&model.Rights{}).Find(&struct_RightsList)

	// 把一个自定义结构体的array 转换成map的array
	// 这里用了json的方法 虽然效率低 但是解决了返回给前端大小写的问题
	for i := 0; i < len(struct_RightsList); i++ {
		map_item := utils.Struct2MapViaJson(struct_RightsList[i])
		map_RightsList = append(map_RightsList, map_item)
	}

	// log.Println(map_RightsList)
	return map_RightsList
}
