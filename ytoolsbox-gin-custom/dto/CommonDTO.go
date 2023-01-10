package dto

type TestSSHDTO struct {
	// ssh链接信息
	IP       string `form:"ip" json:"ip"`
	Port     string `form:"port" json:"port"`
	Username string `form:"username" json:"username" `
	Password string `form:"password" json:"password"`
}
