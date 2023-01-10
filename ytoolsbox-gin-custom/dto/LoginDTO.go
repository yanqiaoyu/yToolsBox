package dto

// 登录的dto
type LoginDTO struct {
	UserName string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
}
