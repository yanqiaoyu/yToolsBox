package dto

type DeleteSpecifiedcronTaskReq struct {
	CronTaskOriginID   uint `uri:"cronTaskOriginID" binding:"required"`
	CronTaskScheduleID uint `uri:"cronTaskScheduleID" binding:"required"`
}

type DeleteSpecifiedcronTaskResultReq struct {
	CronTaskResultID uint `uri:"cronTaskResultID" binding:"required"`
}

type PostNewcronTaskDTOReq struct {
	CronTaskName      string `form:"cronTaskName" json:"cronTaskName" gorm:"column:cronTaskName" binding:"required"`
	CronTaskDesc      string `form:"cronTaskDesc" json:"cronTaskDesc" gorm:"column:cronTaskDesc"`
	CronTaskFinalList string `form:"cronTaskFinalList" json:"cronTaskFinalList" gorm:"column:cronTaskFinalList" binding:"required" `
	CronTaskTime      string `form:"cronTaskTime" json:"cronTaskTime" gorm:"column:cronTaskTime" binding:"required" `
	CronRunAtOnce     bool   `form:"cronRunAtOnce" json:"cronRunAtOnce" gorm:"column:cronRunAtOnce"`
	// 级联选择器的所有信息
	CronTaskCascaderAllInfo string `form:"cronTaskCascaderAllInfo" json:"cronTaskCascaderAllInfo" gorm:"column:cronTaskCascaderAllInfo"`
	// 级联选择器的选中信息
	CronTaskCascaderSelectedInfo string `form:"cronTaskCascaderSelectedInfo" json:"cronTaskCascaderSelectedInfo" gorm:"column:cronTaskCascaderSelectedInfo"`
}

type GetAllCronTaskDTOReq struct {
	Query    string `json:"query" form:"query" `
	Pagenum  int    `json:"pagenum" form:"pagenum" binding:"required"`
	Pagesize int    `json:"pagesize" form:"pagesize" binding:"required"`
}

type GetAllCronTaskResultDTOReq struct {
	Query    string `json:"query" form:"query" `
	Pagenum  int    `json:"pagenum" form:"pagenum" binding:"required"`
	Pagesize int    `json:"pagesize" form:"pagesize" binding:"required"`
}

type GetAllCronTaskItemDTOResp struct {
	Total            int64                    `json:"total"`
	CronTaskItemList []map[string]interface{} `json:"cronTaskItemList"`
}

type GetSpecifiedCrontaskByScheduleIDDTOReq struct {
	CronTaskScheduleID uint `uri:"cronTaskScheduleID" binding:"required"`
}
