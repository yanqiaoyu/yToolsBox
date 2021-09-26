/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu
 * @Date: 2021-06-22 12:45:43
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-09-12 16:10:46
 * @FilePath: /ytoolsbox-gin/model/model.go
 */

package model

import "gorm.io/gorm"

// 在这里定义好表的结构
type User struct {
	gorm.Model
	// 用户名需要唯一
	UserName string `json:"username" gorm:"column:username;unique"`
	Mobile   string `json:"mobile" gorm:"column:mobile"`
	Type     int    `json:"type" gorm:"column:type"`
	Email    string `json:"email" gorm:"column:email"`
	MgState  bool   `json:"mgstate" gorm:"column:mgstate"`
	RoleName string `json:"role" gorm:"column:role"`
	WorkNum  string `json:"worknum" gorm:"column:worknum"`
	PassWord string `json:"password" gorm:"column:password"`
	RoleID   string `json:"roleid" gorm:"column:roleid"`
}
