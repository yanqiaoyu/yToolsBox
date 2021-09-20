package dto

import "time"

// type UserDTO struct {
// 	Name      string `json:"name"`
// 	TelePhone string `json:"telephone"`
// }

type UserDTO struct {
	ID         string    `json:"id" gorm:"column:id"`
	UserName   string    `json:"username" gorm:"column:username;unique"`
	Mobile     string    `json:"mobile" gorm:"column:mobile"`
	Type       int       `json:"type" gorm:"column:type"`
	Email      string    `json:"email" gorm:"column:email"`
	MgState    bool      `json:"mgstate" gorm:"column:mgstate"`
	RoleName   string    `json:"role" gorm:"column:role"`
	CreateTime time.Time `json:"createtime" gorm:"column:created_at"`
}

// func ToUserDTO(user model.User) UserDTO {
// 	return UserDTO{
// 		Name:      user.Account,
// 		TelePhone: user.Phone,
// 	}
// }
