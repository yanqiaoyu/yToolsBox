/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu?tab=repositories
 * @Date: 2021-09-12 16:14:37
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-09-12 16:14:38
 * @FilePath: /ytoolsbox-gin/service/MenusService.go
 */
package service

import (
	"main/model"

	"github.com/gin-gonic/gin"
)

func GetMenus(ctx *gin.Context) {
	data := model.Menus{
		Data: []model.MenusData{
			{Id: 1, AuthName: "首页", Path: "home", ChildMenus: []model.ChildMenus{}},
			{Id: 2, AuthName: "任务", Path: "alltask", ChildMenus: []model.ChildMenus{
				{Id: 201, AuthName: "普通任务", Path: "task"},
				{Id: 202, AuthName: "定时任务", Path: "crontask"},
			}},
			{Id: 3, AuthName: "工具盒", Path: "toolbox", ChildMenus: []model.ChildMenus{}},
			{Id: 4, AuthName: "全局配置", Path: "globalconfig", ChildMenus: []model.ChildMenus{
				{Id: 401, AuthName: "用户管理", Path: "users"},
				{Id: 402, AuthName: "权限管理", Path: "rights"},
			}},
			{Id: 5, AuthName: "关于", Path: "about", ChildMenus: []model.ChildMenus{}},
		},
		Meta: model.MenusMeta{
			Msg:         "suc",
			Status_code: 200,
		},
	}

	ctx.JSON(200, data)

}
