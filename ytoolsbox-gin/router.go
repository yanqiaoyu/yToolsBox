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
	// 登录
	r.POST("/api/auth/login", controller.Login)
	// r.GET("/api/auth/info", middleware.AuthMiddleWare(), controller.Info)
	// 获取菜单信息
	r.GET("/api/auth/menus", service.GetMenus)
	// 获取所有用户的信息
	r.GET("/api/auth/users", controller.GetAllUser)
	// 获取特定用户的信息
	r.GET("/api/auth/users/:userID", controller.GetSpecifiedUser)
	// 更新特定用户的状态
	r.PUT("/api/auth/users/state", controller.PutUserState)
	// 更新特定用户的状态
	r.PUT("/api/auth/users/:userID", controller.PutUserInfo)
	// 新增用户
	r.POST("/api/auth/users", controller.PostNewUser)
	return r
}
