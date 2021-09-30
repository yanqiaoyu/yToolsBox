package controller

import (
	"main/common"
	"main/dao"
	"main/dto"
	"main/util"

	"github.com/gin-gonic/gin"
)

// 新增一个工具
func PostNewTool(ctx *gin.Context) {
	db := common.GetDB()
	PostNewToolReq := dto.PostNewToolDTOReq{}
	// 提取参数
	if util.ResolveParam(ctx, &PostNewToolReq) != nil {
		return
	}

	// 写入数据库
	dao.InsertNewTool(db, PostNewToolReq)

}
