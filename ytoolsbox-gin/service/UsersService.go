package service

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CalculateReturnMapLength(pagenum int, pagesize int, userList []map[string]interface{}) (int, int) {
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
	return ArrayStart, ArrayEnd
}

func SplitPutUserStateParam(ctx *gin.Context) (string, int) {
	mgstate := ctx.PostForm("mgstate")
	userID, _ := strconv.Atoi(ctx.PostForm("userID"))
	// log.Println(mgstate, userID)
	return mgstate, userID
}

func SplitPostNewUserParam(ctx *gin.Context) (string, string, string, string, string) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	mobile := ctx.PostForm("mobile")
	email := ctx.PostForm("email")
	worknum := ctx.PostForm("worknum")
	log.Println(username, password, mobile, email, worknum)
	return username, password, mobile, email, worknum
}
