/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu
 * @Date: 2021-03-26 10:43:57
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-09-12 12:20:58
 * @FilePath: /ytoolsbox-gin/main.go
 */
package main

import (
	"main/common"
	"main/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 1.首先初始化读取配置文件
	InitConfig()
	// 2.从配置文件中拿到了配置，那么可以初始化数据库了
	common.InitDB()
	// 3.初始化一个服务器
	r := gin.Default()
	r.Use(middleware.Cors())
	// 4.收集所有的路由，统一管理
	r = CollectRouter(r)
	// 5.在特定的端口，运行起我们的服务器
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	// 获取当前的目录
	workdir, _ := os.Getwd()
	// 告诉viper配置文件的名称
	viper.SetConfigName("application")
	// 告诉viper配置文件的格式
	viper.SetConfigType("yml")
	// 告诉viper配置文件的位置
	viper.AddConfigPath(workdir + "/config")
	// 有了名字，后缀，位置，就能确定一个唯一的配置文件了
	err := viper.ReadInConfig()
	// 如果读取失败了，利用panic函数让程序主动崩溃
	if err != nil {
		panic(err)
	}
}
