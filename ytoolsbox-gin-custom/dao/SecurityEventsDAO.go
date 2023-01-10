package dao

import (
	"main/dto"
	"main/model"
	"main/utils"

	"gorm.io/gorm"
)

// 查询所有脆弱性和风险的条目
func SelectAllSecurityEventsDAO(db *gorm.DB, obj dto.GetAllSecurityEventsDTOReq) ([]map[string]interface{}, int) {
	SecurityEventsItem := []model.SecurityEvents{}
	query := obj.Query
	pagenum := obj.Pagenum
	pagesize := obj.Pagesize
	myType := obj.Type
	map_SecurityEventsList := []map[string]interface{}{}

	// 不带Query，返回全部
	// 否则返回like搜索后的结果
	if query == "" {
		if myType == "" {
			// 按照时间升序
			db.Order("created_at desc").Model(&model.SecurityEvents{}).Find(&SecurityEventsItem)
		} else {
			// 需要关注Type
			db.Order("created_at desc").Model(&model.SecurityEvents{}).Where("\"type\" = ?", myType).Find(&SecurityEventsItem)
		}

	} else {
		if myType == "" {
			// 按照时间升序
			db.Debug().Order("created_at desc").Where("\"name\" LIKE ?", "%"+query+"%").Model(&model.SecurityEvents{}).Find(&SecurityEventsItem)
		} else {
			db.Debug().Order("created_at desc").Where("\"name\" LIKE ?", "%"+query+"%").Where("\"type\" = ?", myType).Model(&model.SecurityEvents{}).Find(&SecurityEventsItem)
		}

	}

	DefaultLength := len(SecurityEventsItem)

	// 把一个自定义结构体的array 转换成map的array
	// 这里用了json的方法 虽然效率低 但是解决了返回给前端大小写的问题
	for i := 0; i < DefaultLength; i++ {
		map_item := utils.Struct2MapViaJson(SecurityEventsItem[i])
		map_SecurityEventsList = append(map_SecurityEventsList, map_item)
	}

	// 计算一下需要如何切割数组
	ArrayStart, ArrayEnd := utils.CalculateReturnMapLength(pagenum, pagesize, map_SecurityEventsList)
	// 返回切片后的结果
	return map_SecurityEventsList[ArrayStart:ArrayEnd], DefaultLength
}
