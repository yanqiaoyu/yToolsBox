package dto

type DeleteSpecifiedCronTaskReq struct {
	TaskID uint `uri:"crontaskID" binding:"required"`
}
