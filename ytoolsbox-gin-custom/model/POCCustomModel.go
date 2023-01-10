package model

import "gorm.io/gorm"

type POCConfig struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// 工具盒信息
	ToolBoxAddress     string `form:"toolBoxAddress" json:"toolBoxAddress" gorm:"column:toolBoxAddress"`
	ToolBoxSSHPort     string `form:"toolBoxSSHPort" json:"toolBoxSSHPort" gorm:"column:toolBoxSSHPort"`
	ToolBoxSSHUserName string `form:"toolBoxSSHUserName" json:"toolBoxSSHUserName" gorm:"column:toolBoxSSHUserName"`
	ToolBoxSSHPassword string `form:"toolBoxSSHPassword" json:"toolBoxSSHPassword" gorm:"column:toolBoxSSHPassword"`

	// 大脑信息
	DSCAddress     string `form:"dscAddress" json:"dscAddress" gorm:"column:dscAddress"`
	DSCSSHPort     string `form:"dscSSHPort" json:"dscSSHPort" gorm:"column:dscSSHPort"`
	DSCSSHUserName string `form:"dscSSHUserName" json:"dscSSHUserName" gorm:"column:dscSSHUserName"`
	DSCPassword    string `form:"dscPassword" json:"dscPassword" gorm:"column:dscPassword"`
	DP_Token       string `form:"dp_Token" json:"dp_Token" gorm:"column:dp_Token"`
	// 新增前端账号密码
	DSCWebUserName string `form:"dscWebUserName" json:"dscWebUserName" gorm:"column:dscWebUserName"`
	DSCWebPassword string `form:"dscWebPassword" json:"dscWebPassword" gorm:"column:dscWebPassword"`
}

type RiskAndVulnerability struct {
	gorm.Model
	Name          string `json:"name" gorm:"column:name;unique"`
	Type          string `json:"type" gorm:"column:type"`
	Desc          string `json:"desc" gorm:"column:desc"`
	Level         string `json:"level" gorm:"column:level"`
	TriggerMethod string `json:"triggermethod" gorm:"column:triggermethod"`
}

type RiskAndVulnerabilityLog struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Type      string         `form:"type" json:"type" gorm:"column:type"`
	Name      string         `form:"name" json:"name" gorm:"column:name"`
	URL       string         `form:"url" json:"url" gorm:"column:url"`
}

type AgentInstallConfig struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// 安装类型
	Type string `form:"type" json:"type" gorm:"column:type"`
	// ssh链接信息
	IP       string `form:"ip" json:"ip" gorm:"column:ip"`
	Port     string `form:"port" json:"port" gorm:"column:port"`
	Username string `form:"username" json:"username" gorm:"column:username"`
	Password string `form:"password" json:"password" gorm:"column:password"`
}
