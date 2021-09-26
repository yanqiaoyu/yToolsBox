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
	URL_Prefix := "/api/auth"
	// r.POST(URL_Prefix + "/signup", controller.SignUp)
	// 登录
	r.POST(URL_Prefix+"/login", controller.Login)
	// r.GET(URL_Prefix + "/info", middleware.AuthMiddleWare(), controller.Info)
	// 获取菜单信息
	r.GET(URL_Prefix+"/menus", service.GetMenus)
	// 获取所有用户的信息
	r.GET(URL_Prefix+"/users", controller.GetAllUser)
	// 获取特定用户的信息
	r.GET(URL_Prefix+"/users/:userID", controller.GetSpecifiedUser)
	// 更新特定用户的状态
	r.PUT(URL_Prefix+"/users/state", controller.PutUserState)
	// 更新特定用户的状态
	r.PUT(URL_Prefix+"/users/:userID", controller.PutUserInfo)
	// 新增用户
	r.POST(URL_Prefix+"/users", controller.PostNewUser)
	// 删除用户
	r.DELETE(URL_Prefix+"/users/:userID", controller.DeleteSpecifiedUser)

	// 获取权限
	r.GET(URL_Prefix+"/rights", controller.GetRights)
	return r
}
