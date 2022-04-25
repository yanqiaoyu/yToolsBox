package model

import (
	"gorm.io/gorm"
)

type CronTasks struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      int
	UpdatedAt      int
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	ToolName       string         `form:"toolName" json:"toolName" gorm:"column:toolName"`
	ToolConfigName string         `form:"toolConfigName" json:"toolConfigName" gorm:"column:toolConfigName"`
	Progress       int            `form:"progress" json:"progress" gorm:"column:progress;default:0"`
	IsDone         bool           `form:"isDone" json:"isDone" gorm:"column:isDone"`
	ReturnContent  string         `form:"returnContent" json:"returnContent" gorm:"column:returnContent"`
}
