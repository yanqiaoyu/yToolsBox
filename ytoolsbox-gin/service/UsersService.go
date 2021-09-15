package service

import (
	"log"
	"main/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	query := ctx.Query("query")
	// 这行属性其实就是当前在第几页
	pagenum, _ := strconv.Atoi(ctx.Query("pagenum"))
	// 这行属性其实就是当前每页展示多少条数据
	pagesize, _ := strconv.Atoi(ctx.Query("pagesize"))

	param := ctx.Query("param")
	log.Print("\r\n query:", query, "\r\n pagenum: ", pagenum, "\r\n pagesize: ", pagesize, "\r\n param:", param)

	userList := []model.User{}
	T_or_F := true
	for i := 1; i <= 3; i++ {

		if i%3 == 0 {
			T_or_F = true
		} else {
			T_or_F = false
		}

		userList = append(userList, model.User{
			Id:         i,
			Username:   "yqy" + strconv.Itoa(i),
			Mobile:     "18616358651",
			Type:       1,
			Email:      "123@321.com",
			CreateTime: "2021-01-01 20:36:26.000Z",
			MgState:    T_or_F, // 当前用户的状态
			RoleName:   "管理员",
		})
	}
	// log.Print("\r\n userList: ", userList)

	ArrayStart := 0
	ArrayEnd := 0
	// 需要判断一下会不会溢出
	// 起点溢出情况
	if ((pagenum - 1) * pagesize) < len(userList) {
		ArrayStart = (pagenum - 1) * pagesize
	} else {
		ArrayStart = len(userList)
	}
	// 终点溢出判断
	if ((pagenum-1)*pagesize + pagesize) < len(userList) {
		ArrayEnd = (pagenum-1)*pagesize + pagesize
	} else {
		ArrayEnd = len(userList)
	}

	// 针对所有用户数组切片
	slice_userList := userList[ArrayStart:ArrayEnd]
	// log.Print("\r\n slice_userList: ", slice_userList)

	UserData := model.UsersData{Total: len(userList), Pagenum: pagenum, Users: slice_userList}
	Meta := model.Meta{Msg: "获取用户成功", Status_code: 200}

	ctx.JSON(200, gin.H{"data": UserData, "meta": Meta})
}
