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
	Account  string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"varchar(20);not null;unique"`
	Password string `gorm:"size 255;not null"`
}
