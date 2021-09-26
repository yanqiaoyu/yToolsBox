/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu
 * @Date: 2021-06-22 12:46:59
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-09-12 16:13:23
 * @FilePath: /ytoolsbox-gin/controller/UserController.go
 */

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

// 返回所有用户的方法
func GetAllUser(ctx *gin.Context) {
	db := common.GetDB()
	GetAllerUserReqParam := dto.GetAllUserDTOReq{}

	if util.ResolveParam(ctx, &GetAllerUserReqParam) != nil {
		return
	}

	// 根据参数，从数据库中请求user条目
	userList, DefaultLength := dao.SelectAllUser(db, GetAllerUserReqParam)

	// 构造返回的结构体
	UserData := dto.GetAllUserDTOResp{Total: DefaultLength, Pagenum: GetAllerUserReqParam.Pagenum, Users: userList}
	Meta := dto.SuccessResponseMeta{Message: "获取用户成功", StatusCode: 200}

	response.Success(ctx, util.Struct2MapViaJson(UserData), util.Struct2MapViaJson(Meta))
}

// 返回特定用户的状态
func GetSpecifiedUser(ctx *gin.Context) {
	db := common.GetDB()
	GetSpecifiedUserDTOReq := dto.GetSpecifiedUserDTOReq{}

	if util.ResolveURI(ctx, &GetSpecifiedUserDTOReq) != nil {
		return
	}

	struct_userList := dao.SelectSpecifiedUser(db, GetSpecifiedUserDTOReq.UserID)

	// 构造返回的结构体
	Meta := model.Meta{Msg: "获取用户状态成功", Status_code: 200}

	response.Success(ctx, util.Struct2MapViaJson(struct_userList), util.Struct2MapViaJson(Meta))
}

// 更新用户的状态
func PutUserState(ctx *gin.Context) {
	db := common.GetDB()
	PutUserStateDTOReq := dto.PutUserStateDTOReq{}

	if util.ResolveParam(ctx, &PutUserStateDTOReq) != nil {
		return
	}

	dao.UpdateUserState(db, PutUserStateDTOReq)
}

// 更新用户的信息
func PutUserInfo(ctx *gin.Context) {
	db := common.GetDB()

	// 这里用了2个结构体，是因为一个参数在url中，剩下的参数在form-data中
	PutUserInfoDTOReq := dto.PutUserInfoDTOReq{}
	GetSpecifiedUserDTOReq := dto.GetSpecifiedUserDTOReq{}

	if util.ResolveURI(ctx, &GetSpecifiedUserDTOReq) != nil {
		return
	}

	if util.ResolveParam(ctx, &PutUserInfoDTOReq) != nil {
		return
	}

	dao.UpdateSpecifiedUser(db, PutUserInfoDTOReq, GetSpecifiedUserDTOReq)
}

// 新增用户
func PostNewUser(ctx *gin.Context) {
	db := common.GetDB()
	PostNewUserReq := dto.PostNewUserReq{}
	// 提取参数
	if util.ResolveParam(ctx, &PostNewUserReq) != nil {
		return
	}
	// 写入数据库
	dao.InsertNewUser(db, PostNewUserReq)
}

// 删除特定用户
func DeleteSpecifiedUser(ctx *gin.Context) {
	db := common.GetDB()
	GetSpecifiedUserDTOReq := dto.GetSpecifiedUserDTOReq{}

	if util.ResolveURI(ctx, &GetSpecifiedUserDTOReq) != nil {
		return
	}
	struct_userList := dao.DeleteSpecifiedUser(db, int(GetSpecifiedUserDTOReq.UserID))

	// 构造返回的结构体
	Meta := model.Meta{Msg: "删除用户成功", Status_code: 200}
	response.Success(ctx, util.Struct2MapViaJson(struct_userList), util.Struct2MapViaJson(Meta))
}
