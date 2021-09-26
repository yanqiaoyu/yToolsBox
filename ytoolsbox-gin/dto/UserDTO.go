package dto

import "time"

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

type GetAllUserDTOReq struct {
	Query    string `json:"query" form:"query" `
	Pagenum  int    `json:"pagenum" form:"pagenum" binding:"required"`
	Pagesize int    `json:"pagesize" form:"pagesize" binding:"required"`
}

type GetAllUserDTOResp struct {
	Total   int                      `json:"total"`
	Pagenum int                      `json:"pagenum"`
	Users   []map[string]interface{} `json:"users"`
}

type GetSpecifiedUserDTOReq struct {
	UserID int64 `uri:"userID" binding:"required"`
}

type PutUserStateDTOReq struct {
	Mgstate string `form:"mgstate" binding:"required"`
	UserID  int64  `form:"userID" binding:"required"`
}

type PutUserInfoDTOReq struct {
	Email  string `form:"email"`
	Mobile string `form:"mobile"`
}
