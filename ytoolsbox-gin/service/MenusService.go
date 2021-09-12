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
	"log"
	"main/model"

	"github.com/gin-gonic/gin"
)

func GetMenus(ctx *gin.Context) {
	log.Print("This is GetMenus Service")

	data := model.Menus{
		Data: []model.MenusData{
			{Id: 0, AuthName: "能效总览", Path: "dashboard", ChildMenus: []model.ChildMenus{}},
			{Id: 1, AuthName: "工具盒总览", Path: "toolbox", ChildMenus: []model.ChildMenus{}},
			{Id: 2, AuthName: "全局配置", Path: "globalconfig", ChildMenus: []model.ChildMenus{
				{Id: 201, AuthName: "用户管理", Path: "userconfig"},
				{Id: 202, AuthName: "系统管理", Path: "systemconfig"},
			}},
			{Id: 3, AuthName: "关于", Path: "about", ChildMenus: []model.ChildMenus{}},
		},
		Meta: model.MenusMeta{
			Msg:         "suc",
			Status_code: 200,
		},
	}

	ctx.JSON(200, data)

}
