/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu
 * @Date: 2021-06-22 14:50:51
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-06-22 18:28:05
 * @FilePath: \golang_web\router.go
 */

package main

import (
	"main/controller"
	"main/service"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	// r.POST("/api/auth/signup", controller.SignUp)
	r.POST("/api/auth/login", controller.Login)
	// r.GET("/api/auth/info", middleware.AuthMiddleWare(), controller.Info)
	r.GET("/api/auth/menus", service.GetMenus)
	r.GET("/api/auth/users", service.GetUsers)
	return r
}
