package model

import "gorm.io/gorm"

type Rights struct {
	gorm.Model
	AuthName string `json:"authname" gorm:"column:authname;unique"`
	Level    int8   `json:"level" gorm:"column:level"`
	Pid      int    `json:"pid" gorm:"column:pid;default:0"`
	Path     string `json:"path" gorm:"column:path"`
}
