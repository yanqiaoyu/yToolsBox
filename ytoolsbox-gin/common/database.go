/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu
 * @Date: 2021-06-22 14:26:36
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-06-22 19:16:11
 * @FilePath: \golang_web\common\database.go
 */

package common

import (
	"fmt"
	"main/model"
	"strconv"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// 一系列的读取配置操作
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")

	args := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		username,
		password,
		database,
		port,
	)
	// 然后连接这个数据库
	db, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	if err != nil {
		panic("fail to connect to postgres, error" + err.Error())
	}

	InitAllTabls(db)

	DB = db
	return db

}

// 初始化所有表
func InitAllTabls(db *gorm.DB) {
	InitUserTabel(db)
}

// 初始化用户表
func InitUserTabel(db *gorm.DB) {
	UserList := []model.User{

		// 默认的超级管理员
		{
			UserName: "admin",
			Mobile:   "18578660000",
			Type:     1,
			Email:    "yqy1160058763@qq.com",
			MgState:  true, RoleName: "超级管理员",
			WorkNum: "10000颜桥宇", PassWord: "admin",
		},
		// 默认的访客
		{
			UserName: "guest",
			Mobile:   "18578660000",
			Type:     1,
			Email:    "yqy1160058763@qq.com",
			MgState:  true,
			RoleName: "访客",
			PassWord: "guest",
		},
	}

	// 填充一些测试数据，后续需要删掉
	T_or_F := true
	for i := 1; i <= 10; i++ {

		if i%3 == 0 {
			T_or_F = true
		} else {
			T_or_F = false
		}

		UserList = append(UserList, model.User{
			UserName: "测试用户" + strconv.Itoa(i),
			Mobile:   "18616350000",
			Type:     1,
			Email:    "123@321.com",
			MgState:  T_or_F, // 当前用户的状态
			RoleName: "管理员",
		})
	}

	db.AutoMigrate(&model.User{})
	db.Create(&UserList)
}

func GetDB() *gorm.DB {
	return DB
}
