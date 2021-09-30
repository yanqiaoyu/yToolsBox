package dto

type PostNewToolDTOReq struct {
	// 工具类型：脚本，容器
	ToolType string `form:"toolType" json:"toolType" binding:"required"`
	// 镜像名称
	ToolDockerImageName string `form:"toolDockerImageName" json:"toolDockerImageName" binding:"required"`
	// 工具名称
	ToolName string `form:"toolName" json:"toolName" binding:"required"`
	// 工具运行需要的参数
	ToolOptions string `form:"toolOptions" json:"toolOptions"`
	// 工具最终执行指令
	ToolRunCMD string `form:"toolRunCMD" json:"toolRunCMD"`
	// 工具描述
	ToolDesc string `form:"toolDesc" json:"toolDesc"`
	// 工具执行位置
	ToolExecuteLocation string `form:"toolExecuteLocation" json:"toolExecuteLocation" binding:"required"`
	// 远程ssh IP
	ToolRemoteIP string `form:"toolRemoteIP" json:"toolRemoteIP"`
	// 远程ssh Port
	ToolRemoteSSH_Port string `form:"toolRemoteSSH_Port" json:"toolRemoteSSH_Port"`
	// 远程ssh Account
	ToolRemoteSSH_Account string `form:"toolRemoteSSH_Account" json:"toolRemoteSSH_Account"`
	// 远程ssh Password
	ToolRemoteSSH_Password string `form:"toolRemoteSSH_Password" json:"toolRemoteSSH_Password"`

	// 工具作者
	ToolAuthor string `form:"toolAuthor" json:"toolAuthor" binding:"required"`
	// 作者联系方式
	ToolAuthorMobile string `form:"toolAuthorMobile" json:"toolAuthorMobile"`
}
