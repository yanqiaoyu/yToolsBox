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

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"

	// "gorm.io/driver/mysql"
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
	// charset := viper.GetString("datasource.charset")
	// 格式化好配置
	// args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
	// 	username,
	// 	password,

	// 	host,
	// 	port,

	// 	database,
	// 	charset)
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

	// 从model里面读取表结构，然后在数据库中初始化这个表
	db.AutoMigrate(&model.User{})

	r := db.Exec(`CREATE TABLE company( 
		ID INT PRIMARY KEY     NOT NULL,
		NAME           TEXT    NOT NULL,
		AGE            INT     NOT NULL,
		ADDRESS        CHAR(50),
		SALARY         REAL
	 )`)

	db.Exec("commit")

	fmt.Print(r)

	DB = db
	return db

}

func GetDB() *gorm.DB {
	return DB
}
