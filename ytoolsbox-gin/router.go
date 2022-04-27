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
	/* 用路由组重新归纳了一下路由 */
	v1 := r.Group(URL_Prefix)
	{
		// r.POST(URL_Prefix + "/signup", controller.SignUp)
		// 登录
		v1.POST("/login", controller.Login)
		// r.GET(URL_Prefix + "/info", middleware.AuthMiddleWare(), controller.Info)
		// 获取菜单信息
		v1.GET("/menus", service.GetMenus)
		// 获取所有用户的信息
		v1.GET("/users", controller.GetAllUser)
		// 获取特定用户的信息
		v1.GET("/users/:userID", controller.GetSpecifiedUser)
		// 更新特定用户的状态
		v1.PUT("/users/state", controller.PutUserState)
		// 更新特定用户的信息
		v1.PUT("/users/:userID", controller.PutUserInfo)
		// 新增用户
		v1.POST("/users", controller.PostNewUser)
		// 删除用户
		v1.DELETE("/users/:userID", controller.DeleteSpecifiedUser)

		// 获取权限
		v1.GET("/rights", controller.GetRights)

		// 添加新工具
		v1.POST("/tools", controller.PostNewTool)
		// 查询所有工具
		v1.GET("/tools", controller.GetAllTools)
		// 删除所有工具
		v1.DELETE("/tools", controller.DeleteAllTools)
		// 查询某个工具的所有配置
		v1.GET("/tools/config/:toolID", controller.GetSpecifiedToolConfig)
		// 查询某个工具的某个配置
		v1.GET("/tools/config/:toolID/:configID", controller.GetSpecifiedToolConfigByConfigID)
		// 更新某个工具的某个配置
		v1.PUT("/tools/config/:toolID/:configID", controller.PutSpecifiedToolConfigByConfigID)
		// 为某个工具新增配置
		v1.POST("/tools/config/:toolID", controller.PostNewConfig)
		// 删除某个工具下的某个配置
		v1.DELETE("/tools/config/:toolID/:configID", controller.DeleteSpecifiedConfig)
		// 上传脚本文件
		v1.POST("/upload", controller.PostScriptFile)

		/***
			以下是任务相关的路由表
		***/

		// 新建一个任务
		v1.POST("/tasks", controller.PostNewTask)
		// 查询Cascader里面的信息
		v1.GET("/tasks/cascader", controller.GetCascader)
		// 查询所有的TaskItem(任务进度)
		v1.GET("/tasks", controller.GetTaskItem)
		// 清空所有任务
		v1.DELETE("/tasks", controller.DeleteAllTask)
		// 删除特定任务
		v1.DELETE("/tasks/:taskID", controller.DeleteSpecifiedTask)
		// 重新开始执行一个任务
		v1.POST("/tasks/restart", controller.PostRestartTask)

		/***
			以下是定时任务相关的路由表
		***/

		// 新建一个定时任务
		v1.POST("/crontasks", controller.PostNewCronTask)
		// 清除所有定时任务
		v1.DELETE("/crontasks", controller.DeleteAllCronTask)

	}
	return r
}
