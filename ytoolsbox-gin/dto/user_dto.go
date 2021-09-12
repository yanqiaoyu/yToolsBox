package dto

import "main/model"

type UserDTO struct {
	Name      string `json:"name"`
	TelePhone string `json:"telephone"`
}

func ToUserDTO(user model.User) UserDTO {
	return UserDTO{
		Name:      user.Account,
		TelePhone: user.Phone,
	}
}

// func ToUserDTO(user model.User) UserDTO {
// 	return UserDTO{
// 		Name: user.Name,
// 		// TelePhone: user.Phone,
// 	}
// }
