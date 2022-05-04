package model

import (
	"gorm.io/gorm"
)

type CronTasks struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// taskID 这个ID是schedule返回的ID
	CronTaskScheduleID int `form:"cronTaskScheduleID" json:"cronTaskScheduleID" gorm:"column:cronTaskScheduleID"`
	// 定时任务名称
	CronTaskName string `form:"cronTaskName" json:"cronTaskName" gorm:"column:cronTaskName"`
	// 定时任务描述
	CronTaskDesc string `form:"cronTaskDesc" json:"cronTaskDesc" gorm:"column:cronTaskDesc"`
	// 工具以及配置列表
	CronTaskFinalList string `form:"cronTaskFinalList" json:"cronTaskFinalList" gorm:"column:cronTaskFinalList"`
	// 时间表达式
	CronTaskTime string `form:"cronTaskTime" json:"cronTaskTime" gorm:"column:cronTaskTime"`
	// 是否立即执行一次
	CronRunAtOnce bool `form:"cronRunAtOnce" json:"cronRunAtOnce" gorm:"column:cronRunAtOnce"`
	// 级联选择器的所有信息
	CronTaskCascaderAllInfo string `form:"cronTaskCascaderAllInfo" json:"cronTaskCascaderAllInfo" gorm:"column:cronTaskCascaderAllInfo"`
	// 级联选择器的选中信息
	CronTaskCascaderSelectedInfo string `form:"cronTaskCascaderSelectedInfo" json:"cronTaskCascaderSelectedInfo" gorm:"column:cronTaskCascaderSelectedInfo"`

	// Progress      int    `form:"progress" json:"progress" gorm:"column:progress;default:0"`
	// IsDone        bool   `form:"isDone" json:"isDone" gorm:"column:isDone"`
	// ReturnContent string `form:"returnContent" json:"returnContent" gorm:"column:returnContent"`
}
