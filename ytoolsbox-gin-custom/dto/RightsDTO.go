package dto

type RightsListDTO struct {
	ID       int64  `json:"id" gorm:"colume:id"`
	AuthName string `json:"authname" gorm:"column:authname"`
	Level    int8   `json:"level" gorm:"column:level"`
	Pid      int    `json:"pid" gorm:"column:pid;default:0"`
	Path     string `json:"path" gorm:"column:path"`
}

type GetAllRightsResp struct {
	Total      int64                    `json:"total"`
	RightsList []map[string]interface{} `json:"rightslist"`
}
