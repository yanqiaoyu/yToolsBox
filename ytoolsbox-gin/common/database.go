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
	"flag"
	"fmt"
	"main/model"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var compileMode, host string

	flag.StringVar(&compileMode, "m", "test", "运行模式")
	flag.Parse()

	// 区分生产环境和测试环境
	if compileMode == "production" {
		host = viper.GetString("datasource.productionhost")
	} else {
		host = viper.GetString("datasource.testhost")
	}
	// fmt.Println("host is", host)
	// 一系列的读取配置操作

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

	InitAllTables(db)

	DB = db
	return db

}

// 初始化所有表
func InitAllTables(db *gorm.DB) {
	InitUserTable(db)
	InitRightsTable(db)
	InitToolsTable(db)
	InitToolsConfigTable(db)
	InitTaskTable(db)
	InitCronTaskTable(db)
}

// 初始化工具基础信息表
func InitToolsTable(db *gorm.DB) {
	db.AutoMigrate(&model.Tool{})
}

// 初始化工具配置信息表
func InitToolsConfigTable(db *gorm.DB) {
	db.AutoMigrate(&model.ToolConfig{})
}

// 初始化用户表
func InitUserTable(db *gorm.DB) {
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

	db.AutoMigrate(&model.User{})
	db.Create(&UserList)
}

// 初始化权限表
func InitRightsTable(db *gorm.DB) {
	RightsList := []model.Rights{
		{AuthName: "首页", Level: 0, Pid: 0, Path: "home"},
		{AuthName: "任务", Level: 1, Pid: 0, Path: "dashboard"},
		{AuthName: "工具盒", Level: 1, Pid: 0, Path: "toolbox"},
		{AuthName: "全局配置", Level: 2, Pid: 0, Path: "config"},
		{AuthName: "用户管理", Level: 2, Pid: 4, Path: "users"},
		{AuthName: "权限管理", Level: 2, Pid: 4, Path: "rights"},
	}

	db.AutoMigrate(&model.Rights{})
	db.Create(&RightsList)
}

// 初始化任务列表
func InitTaskTable(db *gorm.DB) {
	db.AutoMigrate(&model.Tasks{})
}

// 初始化定时任务列表
func InitCronTaskTable(db *gorm.DB) {
	db.AutoMigrate(&model.CronTasks{})
}

func GetDB() *gorm.DB {
	return DB
}
