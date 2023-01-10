package dto

type GetAllSecurityEventsDTOReq struct {
	Query    string `json:"query" form:"query" `
	Pagenum  int    `json:"pagenum" form:"pagenum" binding:"required"`
	Pagesize int    `json:"pagesize" form:"pagesize" binding:"required"`
	Type     string `json:"type" form:"type"`
}

type GetAllSecurityEventsDTOResp struct {
	Total              int64
	SecurityEventsList []map[string]interface{}
}
