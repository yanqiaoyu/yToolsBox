package model

import "gorm.io/gorm"

type Tool struct {
	gorm.Model
	// 工具类型：脚本，容器
	ToolType string `form:"toolType" json:"toolType" binding:"required" gorm:"column:toolType"`
	// 镜像名称
	ToolDockerImageName string `form:"toolDockerImageName" json:"ToolDockerImageName" binding:"required" gorm:"column:toolDockerImageName"`
	// 工具名称
	ToolName string `form:"toolName" json:"toolName" binding:"required" gorm:"column:toolName"`
	// 工具运行需要的参数
	ToolOptions string `form:"toolOptions" json:"toolOptions" gorm:"column:toolOptions"`
	// 工具最终执行指令
	ToolRunCMD string `form:"toolRunCMD" json:"toolRunCMD" gorm:"column:toolRunCMD"`
	// 工具描述
	ToolDesc string `form:"toolDesc" json:"toolDesc" gorm:"column:toolDesc"`
	// 工具执行位置
	ToolExecuteLocation string `form:"toolExecuteLocation" json:"toolExecuteLocation" binding:"required" gorm:"column:toolExecuteLocation"`
	// 远程ssh IP
	ToolRemoteIP string `form:"toolRemoteIP" json:"toolRemoteIP" gorm:"column:toolRemoteIP"`
	// 远程ssh Port
	ToolRemoteSSH_Port string `form:"toolRemoteSSH_Port" json:"toolRemoteSSH_Port" gorm:"column:toolRemoteSSH_Port"`
	// 远程ssh Account
	ToolRemoteSSH_Account string `form:"toolRemoteSSH_Account" json:"toolRemoteSSH_Account" gorm:"column:toolRemoteSSH_Account"`
	// 远程ssh Password
	ToolRemoteSSH_Password string `form:"toolRemoteSSH_Password" json:"toolRemoteSSH_Password" gorm:"column:toolRemoteSSH_Password"`

	// 工具作者
	ToolAuthor string `form:"toolAuthor" json:"toolAuthor" gorm:"column:toolAuthor"`
	// 作者联系方式
	ToolAuthorMobile string `form:"toolAuthorMobile" json:"toolAuthorMobile" gorm:"column:toolAuthorMobile"`
}
