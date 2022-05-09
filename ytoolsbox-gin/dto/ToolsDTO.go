package dto

import (
	"main/model"
)

// 新增一个工具
type PostNewToolBasicInfoDTOReq model.Tool
type PostNewToolConfigInfoDTOReq model.ToolConfig

// 工具卡片信息
type BriefToolsInfoDTO struct {
	// 工具ID
	ID int `form:"id" json:"id" gorm:"column:id"`
	// 工具名称
	ToolName string `form:"toolName" json:"toolName" gorm:"column:toolName"`
	// 工具描述
	ToolDesc string `form:"toolDesc" json:"toolDesc" gorm:"column:toolDesc"`
	// 工具作者
	ToolAuthor string `form:"toolAuthor" json:"toolAuthor" gorm:"column:toolAuthor"`
	// 工具评分
	ToolRate float32 `form:"toolRate" json:"toolRate" gorm:"column:toolRate"`
	//  工具评分人数
	ToolRateCount int8 `form:"toolRateCount" json:"toolRateCount" gorm:"column:toolRateCount"`
	// 工具类型：脚本，容器
	ToolType string `form:"toolType" json:"toolType" gorm:"column:toolType"`
	// 工具使用说明
	ToolTutorial string `form:"toolTutorial" json:"toolTutorial" gorm:"column:toolTutorial"`
}

// 查询所有工具卡片信息的请求DTO
type GetAllToolsDTOReq struct {
	Query string `json:"query" form:"query" `
}

// 查询所有工具卡片信息的响应DTO
type GetAllToolsDTOResp struct {
	Total int                 `json:"total"`
	Tools []BriefToolsInfoDTO `json:"tools"`
}

// 查询特定工具配置信息的请求DTO(URI)
type GetSpecifiedToolConfigDTOReqURI struct {
	ToolID int16 `uri:"toolID" binding:"required"`
}

// 查询特定工具配置信息的请求DTO(QueryInfo)
type GetSpecifiedToolConfigDTOReqQuery struct {
	Query    string `json:"query" form:"query" `
	Pagenum  int    `json:"pagenum" form:"pagenum" binding:"required"`
	Pagesize int    `json:"pagesize" form:"pagesize" binding:"required"`
}

// 查询特定工具配置信息的响应DTO
type GetSpecifiedToolConfigDTOResp struct {
	// 工具ID
	ID         int16                    `form:"id" json:"id" gorm:"column:id"`
	Total      int                      `json:"total"`
	ToolConfig []map[string]interface{} `json:"toolsConfig" gorm:"foreignkey:toolID;association_autoupdate:false"`
}

// 查询某个工具的某个配置请求DTO
type GetSpecifiedToolConfigByConfigIDDTOReq struct {
	ToolID int16 `uri:"toolID" gorm:"column:toolID" binding:"required"`
	ID     uint  `uri:"configID" gorm:"column:id" binding:"required"`
}

// 查询某个工具的某个配置响应DTO
type GetSpecifiedToolConfigByConfigIDDTOResp struct {
	ToolConfig BriefToolConfigDTO `json:"toolConfig"`
}

// 更新某个工具的某个配置请求DTO(URI)
type PutSpecifiedToolConfigByConfigIDDTOReqURI struct {
	ToolID int16 `uri:"toolID" gorm:"column:toolID" binding:"required"`
	ID     uint  `uri:"configID" gorm:"column:id" binding:"required"`
}

// 更新某个工具的某个配置请求DTO(QueryInfo)
type PutSpecifiedToolConfigByConfigIDDTOReqQuery struct {
	ToolConfig BriefToolConfigDTO `json:"toolConfig"`
}

// 更新某个工具的使用说明
type PutSpecifiedToolTutorialByToolIDDTOReq struct {
	ToolID       int16  `form:"toolID" json:"toolID" gorm:"column:toolID"`
	ToolTutorial string `form:"toolTutorial" json:"toolTutorial" gorm:"column:toolTutorial"`
}

// 删除某个工具的某个配置请求DTO
type DeleteSpecifiedConfigDTOReq struct {
	ToolID int16 `uri:"toolID" gorm:"column:toolID" binding:"required"`
	// 配置ID
	ID uint `uri:"configID" gorm:"column:id" binding:"required"`
}

type BriefToolConfigDTO struct {
	// 配置ID
	ID int `form:"id" json:"id" gorm:"column:id"`
	// 工具ID,对应的是tools表里面的id列
	ToolID int `form:"toolID" json:"toolID" gorm:"column:toolID"`
	// 配置名称
	ToolConfigName string `form:"toolConfigName" json:"toolConfigName" gorm:"column:toolConfigName"`
	// 配置描述
	ToolConfigDesc string `form:"toolConfigDesc" json:"toolConfigDesc" gorm:"column:toolConfigDesc"`

	// 工具类型：脚本，容器
	ToolType string `form:"toolType" json:"toolType" gorm:"column:toolType" binding:"required"`
	// 镜像名称
	ToolDockerImageName string `form:"toolDockerImageName" json:"toolDockerImageName" gorm:"column:toolDockerImageName"`
	// 脚本工具的名称
	ToolScriptName string `form:"toolScriptName" json:"toolScriptName" gorm:"column:toolScriptName"`
	// 脚本在本地的存放路径
	ToolScriptLocalPath string `form:"toolScriptLocalPath" json:"toolScriptLocalPath" gorm:"column:toolScriptLocalPath"`
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
