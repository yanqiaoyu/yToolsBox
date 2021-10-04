package model

import "gorm.io/gorm"

// 工具的基本信息
type Tool struct {
	gorm.Model
	// 工具名称
	ToolName string `form:"toolName" json:"toolName" gorm:"column:toolName;unique" binding:"required"`
	// 工具描述
	ToolDesc string `form:"toolDesc" json:"toolDesc" gorm:"column:toolDesc;default:'这个工具很神秘,没有留下介绍'"`
	// 工具作者
	ToolAuthor string `form:"toolAuthor" json:"toolAuthor" gorm:"column:toolAuthor" binding:"required"`
	// 工具作者联系方式
	ToolAuthorMobile string `form:"toolAuthorMobile" json:"toolAuthorMobile" gorm:"column:toolAuthorMobile"`
	// 工具评分
	ToolRate float32 `form:"toolRate" json:"toolRate" gorm:"column:toolRate"`
	//  工具评分人数
	ToolRateCount int8 `form:"toolRateCount" json:"toolRateCount" gorm:"column:toolRateCount"`
}

// 工具的配置信息
type ToolConfig struct {
	gorm.Model
	// 工具ID,对应的是tools表里面的id列
	ToolID int `form:"toolID" json:"toolID" gorm:"column:toolID"`
	// 配置名称
	ToolConfigName string `form:"toolConfigName" json:"toolConfigName" gorm:"column:toolConfigName;default:'默认配置'"`
	// 配置描述
	ToolConfigDesc string `form:"toolConfigDesc" json:"toolConfigDesc" gorm:"column:toolConfigDesc;default:'根据添加工具时的信息生产的默认配置'"`

	// 工具类型：脚本，容器
	ToolType string `form:"toolType" json:"toolType" gorm:"column:toolType" binding:"required"`
	// 镜像名称
	ToolDockerImageName string `form:"toolDockerImageName" json:"toolDockerImageName" gorm:"column:toolDockerImageName"`
	// 脚本工具的名称
	ToolScriptName string `form:"toolScriptName" json:"toolScriptName" gorm:"column:toolScriptName"`
	// 脚本的存放路径
	ToolScriptPath string `form:"toolScriptPath" json:"toolScriptPath" gorm:"column:toolScriptPath"`
	// 工具运行需要的参数
	ToolOptions string `form:"toolOptions" json:"toolOptions" gorm:"column:toolOptions"`
	// 工具最终执行指令
	ToolRunCMD string `form:"toolRunCMD" json:"toolRunCMD" gorm:"column:toolRunCMD"`
	// 工具执行位置
	ToolExecuteLocation string `form:"toolExecuteLocation" json:"toolExecuteLocation" gorm:"column:toolExecuteLocation" binding:"required"`
	// 远程ssh IP
	ToolRemoteIP string `form:"toolRemoteIP" json:"toolRemoteIP" gorm:"column:toolRemoteIP"`
	// 远程ssh Port
	ToolRemoteSSH_Port string `form:"toolRemoteSSH_Port" json:"toolRemoteSSH_Port" gorm:"column:toolRemoteSSH_Port"`
	// 远程ssh Account
	ToolRemoteSSH_Account string `form:"toolRemoteSSH_Account" json:"toolRemoteSSH_Account" gorm:"column:toolRemoteSSH_Account"`
	// 远程ssh Password
	ToolRemoteSSH_Password string `form:"toolRemoteSSH_Password" json:"toolRemoteSSH_Password" gorm:"column:toolRemoteSSH_Password"`
	// Python版本
	ToolPythonVersion string `form:"toolPythonVersion" json:"toolPythonVersion" gorm:"column:toolPythonVersion"`
	// Shell版本
	ToolShellVersion string `form:"toolShellVersion" json:"toolShellVersion" gorm:"column:toolShellVersion"`
}
