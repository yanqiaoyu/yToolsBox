package dao

import (
	"fmt"
	"log"
	"main/model"
	"main/utils"
	"strings"

	"gorm.io/gorm"
)

// 随机泄漏3条数据
func SelectRandomLeakgeData(db *gorm.DB) []map[string]interface{} {
	leakageData := []model.Demo_table{}
	map_leakageData := []map[string]interface{}{}

	// 这里从3张表里, 各随机查询一条数据
	for i := 1; i <= 3; i++ {
		SQL := fmt.Sprintf("select * from demo%d_table limit 1 offset %d", i, utils.GetAnRandomInt(1, 99))
		db.Debug().Raw(SQL).Scan(&leakageData)
		map_item := utils.Struct2MapViaJson(leakageData[0])
		// 去掉查询结果,首尾的空格,这个问题是建库时引起的,懒得重新build镜像了,在这里处理了
		for k, v := range map_item {
			map_item[k] = strings.Trim(fmt.Sprintf("%v", v), " ")
		}
		map_leakageData = append(map_leakageData, map_item)
	}

	log.Println("泄露数据: ", map_leakageData)
	return map_leakageData
}
