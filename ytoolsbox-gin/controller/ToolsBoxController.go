package controller

import (
	"main/common"
	"main/dao"
	"main/dto"
	"main/response"
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

// 查询所有工具
func GetAllTools(ctx *gin.Context) {

	db := common.GetDB()
	tools_List := dao.SelectAllTools(db)

	// 构造返回的结构体
	ToolData := dto.GetAllToolsDTOResp{Total: len(tools_List), Tools: tools_List}
	Meta := dto.SuccessResponseMeta{Message: "获取用户成功", StatusCode: 200}

	response.Success(ctx, util.Struct2MapViaJson(ToolData), util.Struct2MapViaJson(Meta))
}
