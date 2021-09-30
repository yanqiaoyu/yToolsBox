package controller

import (
	"log"
	"main/dto"
	"main/util"

	"github.com/gin-gonic/gin"
)

// 新增一个工具
func PostNewTool(ctx *gin.Context) {
	// db := common.GetDB()
	PostNewToolReq := dto.PostNewToolDTOReq{}
	// 提取参数
	if util.ResolveParam(ctx, &PostNewToolReq) != nil {
		return
	}
	log.Println(PostNewToolReq)
	// 写入数据库
	// dao.InsertNewTool(db, PostNewToolReq)

}
