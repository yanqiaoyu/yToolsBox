package model

import "gorm.io/gorm"

type CronTasksResult struct {
	ID                 uint `gorm:"primaryKey"`
	CreatedAt          int
	UpdatedAt          int
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	CronTaskID         uint           `form:"cronTaskID" json:"cronTaskID" gorm:"column:cronTaskID"`
	CronTaskScheduleID uint           `form:"cronTaskScheduleID" json:"cronTaskScheduleID" gorm:"column:cronTaskScheduleID"`
	CronTaskName       string         `form:"cronTaskName" json:"cronTaskName" gorm:"column:cronTaskName"`
	CronTaskDesc       string         `form:"cronTaskDesc" json:"cronTaskDesc" gorm:"column:cronTaskDesc"`
	ToolName           string         `form:"toolName" json:"toolName" gorm:"column:toolName"`
	ToolConfigName     string         `form:"toolConfigName" json:"toolConfigName" gorm:"column:toolConfigName"`
	Progress           int            `form:"progress" json:"progress" gorm:"column:progress;default:0"`
	IsDone             bool           `form:"isDone" json:"isDone" gorm:"column:isDone"`
	ReturnContent      string         `form:"returnContent" json:"returnContent" gorm:"column:returnContent"`
}
