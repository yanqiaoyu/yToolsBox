package controller

import (
	"main/common"
	"main/dao"
	"main/dto"
	"main/model"
	"main/response"
	"main/service"
	"main/util"

	"github.com/gin-gonic/gin"
)

// 新增一个工具
func PostNewTool(ctx *gin.Context) {
	// 提取工具的基本信息
	PostNewToolBasicInfoDTOReq := dto.PostNewToolBasicInfoDTOReq{}
	// 提取参数
	if util.ResolveParam(ctx, &PostNewToolBasicInfoDTOReq) != nil {
		return
	}

	// 提取工具的配置信息
	PostNewToolConfigInfoDTOReq := dto.PostNewToolConfigInfoDTOReq{}
	// 提取参数
	if util.ResolveParam(ctx, &PostNewToolConfigInfoDTOReq) != nil {
		return
	}

	// log.Print(util.Struct2MapViaJson(PostNewToolBasicInfoDTOReq), util.Struct2MapViaJson(PostNewToolConfigInfoDTOReq))

	db := common.GetDB()

	// 写入基本信息
	result := dao.InsertNewToolBasicInfo(db, &PostNewToolBasicInfoDTOReq)
	if result.Error != nil {
		// 再调用一次Error(),转换为字符串
		if result.Error.Error() == "ERROR: duplicate key value violates unique constraint \"tools_toolName_key\" (SQLSTATE 23505)" {
			msg := dto.FailResponseMeta{StatusCode: 400, Message: "不允许重复的工具名称"}
			response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
			return
		}
	}

	// 写入配置信息
	PostNewToolConfigInfoDTOReq.ToolID = int(PostNewToolBasicInfoDTOReq.ID)
	result = dao.InsertNewToolConfigInfo(db, &PostNewToolConfigInfoDTOReq)
	if result.Error != nil {
		// 再调用一次Error(),转换为字符串
		msg := dto.FailResponseMeta{StatusCode: 400, Message: "工具配置信息写入错误"}
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
		return
	}

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

// 查询某个工具的所有配置
func GetSpecifiedToolConfig(ctx *gin.Context) {
	db := common.GetDB()
	GetSpecifiedToolConfigDTOReq := dto.GetSpecifiedToolConfigDTOReq{}

	if util.ResolveURI(ctx, &GetSpecifiedToolConfigDTOReq) != nil {
		return
	}
	configList := dao.SelectSpecifiedToolConfig(db, int(GetSpecifiedToolConfigDTOReq.ToolID))
	// log.Println(configList)

	// 构造返回的结构体
	ToolData := dto.GetSpecifiedToolConfigDTOResp{Total: len(configList), ToolConfig: configList, ID: GetSpecifiedToolConfigDTOReq.ToolID}
	Meta := model.Meta{Msg: "查询配置成功", Status_code: 200}
	response.Success(ctx, util.Struct2MapViaJson(ToolData), util.Struct2MapViaJson(Meta))
}

// 上传脚本文件
func PostScriptFile(ctx *gin.Context) {
	db := common.GetDB()
	toolName, FileDST, err := service.SaveScriptFile(ctx)
	if err != nil {
		msg := dto.FailResponseMeta{StatusCode: 400, Message: "上传文件失败"}
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
	}

	dao.UpdateToolConfigScriptLocalPath(db, toolName, FileDST)

	Meta := model.Meta{Msg: "上传文件成功", Status_code: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 新增配置
func PostNewConfig(ctx *gin.Context) {
	
}