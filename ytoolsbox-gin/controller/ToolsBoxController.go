package controller

import (
	"main/common"
	"main/dao"
	"main/dto"
	"main/model"
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

	Meta := model.Meta{Msg: "新增工具成功", Status_code: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 查询所有工具
func GetAllTools(ctx *gin.Context) {
	db := common.GetDB()

	GetAllToolsDTOReq := dto.GetAllToolsDTOReq{}

	if util.ResolveParam(ctx, &GetAllToolsDTOReq) != nil {
		return
	}

	tools_List := dao.SelectAllTools(db, GetAllToolsDTOReq)

	// 构造返回的结构体
	ToolData := dto.GetAllToolsDTOResp{Total: len(tools_List), Tools: tools_List}
	Meta := dto.SuccessResponseMeta{Message: "获取工具成功", StatusCode: 200}

	response.Success(ctx, util.Struct2MapViaJson(ToolData), util.Struct2MapViaJson(Meta))
}

func GetSpecifiedToolConfig(ctx *gin.Context) {
	// db := common.GetDB()
	GetSpecifiedToolConfigReq := dto.GetSpecifiedToolConfigReq{}

	if util.ResolveURI(ctx, &GetSpecifiedToolConfigReq) != nil {
		return
	}
	// struct_userList := dao.DeleteSpecifiedUser(db, int(GetSpecifiedToolConfigReq.ToolID))

	// // 构造返回的结构体
	// Meta := model.Meta{Msg: "删除用户成功", Status_code: 200}
	// response.Success(ctx, util.Struct2MapViaJson(struct_userList), util.Struct2MapViaJson(Meta))
}
