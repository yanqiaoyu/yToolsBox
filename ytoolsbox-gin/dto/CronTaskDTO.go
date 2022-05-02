package dto

type DeleteSpecifiedcronTaskReq struct {
	TaskID uint `uri:"crontaskID" binding:"required"`
}

type PostNewcronTaskDTOReq struct {
	CronTaskName      string `form:"cronTaskName" json:"cronTaskName" gorm:"column:cronTaskName" binding:"required"`
	CronTaskDesc      string `form:"cronTaskDesc" json:"cronTaskDesc" gorm:"column:cronTaskDesc"`
	CronTaskFinalList string `form:"cronTaskFinalList" json:"cronTaskFinalList" gorm:"column:cronTaskFinalList" binding:"required" `
	CronTaskTime      string `form:"cronTaskTime" json:"cronTaskTime" gorm:"column:cronTaskTime" binding:"required" `
	CronRunAtOnce     bool   `form:"cronRunAtOnce" json:"cronRunAtOnce" gorm:"column:cronRunAtOnce"`
}

type GetAllCronTaskDTOReq struct {
	Query    string `json:"query" form:"query" `
	Pagenum  int    `json:"pagenum" form:"pagenum" binding:"required"`
	Pagesize int    `json:"pagesize" form:"pagesize" binding:"required"`
}

type GetAllCronTaskItemDTOResp struct {
	Total            int64                    `json:"total"`
	CronTaskItemList []map[string]interface{} `json:"cronTaskItemList"`
}
