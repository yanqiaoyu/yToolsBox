package model

import "gorm.io/gorm"

type SecurityEvents struct {
	gorm.Model
	Name          string `json:"name" gorm:"column:name;unique"`
	Type          string `json:"type" gorm:"column:type"`
	Desc          string `json:"desc" gorm:"column:desc"`
	Level         string `json:"level" gorm:"column:level"`
	TriggerMethod string `json:"triggermethod" gorm:"column:triggermethod"`
}
