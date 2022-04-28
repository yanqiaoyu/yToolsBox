package model

import (
	"gorm.io/gorm"
)

type CronTasks struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// 定时任务名称
	TaskName string `form:"taskName" json:"taskName" gorm:"column:taskName"`
	// 定时任务描述
	TaskDesc string `form:"taskDesc" json:"taskDesc" gorm:"column:taskDesc"`
	// 使用的工具ID
	ToolID int `form:"toolID" json:"toolID" gorm:"column:toolID"`
	// 使用的工具配置ID
	ToolConfigID int `form:"toolConfigID" json:"toolConfigID" gorm:"column:toolConfigID"`
	// 时间表达式, 用于确定间隔执行的时间
	CronExpr string `form:"cronExpr" json:"cronExpr" gorm:"column:cronExpr"`

	// Progress      int    `form:"progress" json:"progress" gorm:"column:progress;default:0"`
	// IsDone        bool   `form:"isDone" json:"isDone" gorm:"column:isDone"`
	// ReturnContent string `form:"returnContent" json:"returnContent" gorm:"column:returnContent"`
}
