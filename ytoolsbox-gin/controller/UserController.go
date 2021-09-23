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
	"log"
	"main/common"
	"main/dao"
	"main/dto"
	"main/model"
	"main/response"
	"main/service"
	"main/util"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func Login(ctx *gin.Context) {
// 	db := common.GetDB()

// 	//获取参数
// 	phone := ctx.PostForm("Phone")
// 	password := ctx.PostForm("Password")

// 	// 数据验证
// 	if len(phone) != 11 {
// 		fmt.Println(len(phone))
// 		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
// 		return
// 	}

// 	if len(password) < 6 {
// 		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "账户长度必须大于六位")
// 		return
// 	}

// 	var user model.User
// 	db.Where("phone = ?", phone).First(&user)
// 	if user.ID == 0 {
// 		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
// 		return
// 	}

// 	// 判断密码是否正确
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 		response.Fail(ctx, "用户不存在", nil)
// 	}

// 	token, err := common.ReleaseToken(user)
// 	if err != nil {
// 		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常")
// 		log.Print("token generate error: ", err)
// 		return

// 	}

// 	response.Success(ctx, "密码正确", gin.H{"token": token})

// }

// func SignUp(ctx *gin.Context) {
// 	db := common.GetDB()

// 	//获取参数
// 	account := ctx.PostForm("account")
// 	phone := ctx.PostForm("Phone")
// 	password := ctx.PostForm("Password")

// 	//数据验证
// 	if len(phone) != 11 {
// 		fmt.Println(len(phone))
// 		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
// 		return
// 	}

// 	if len(password) < 6 {
// 		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "账户长度必须大于六位")
// 		return
// 	}

// 	if len(account) == 0 {
// 		account = util.GetRandomString2(6)
// 	}

// 	// 判断账号是否存在
// 	if isPhoneExist(db, phone) {
// 		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "该用户已经存在")
// 		return
// 	}

// 	// 准备好要插入的数据
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "加密失败")
// 	}

// 	newUser := model.User{
// 		Account:  account,
// 		Phone:    phone,
// 		Password: string(hashedPassword),
// 	}
// 	// newUser := model.User{
// 	// 	Name:     account,
// 	// 	Password: string(hashedPassword),
// 	// }

// 	// 插入数据
// 	db.Create(&newUser)
// 	response.Success(ctx, "注册成功", nil)

// }

// func isPhoneExist(db *gorm.DB, phone string) bool {
// 	var user model.User
// 	db.Where("phone = ?", phone).First(&user)
// 	if user.ID != 0 {
// 		return true
// 	}

// 	return false
// }

// // 判断登陆的账号是否合法
// func isUserExist() {

// }

// func Info(ctx *gin.Context) {
// 	user, _ := ctx.Get("user")
// 	response.Success(ctx, "", gin.H{"user": dto.ToUserDTO(user.(model.User))})
// }

func Login(ctx *gin.Context) {
	db := common.GetDB()
	loginParam := dto.LoginDTO{}

	// 模型绑定获取参数
	err := ctx.ShouldBind(&loginParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if !dao.IsUserExist(db, loginParam.UserName) {
		response.Fail(ctx, nil, gin.H{
			"status_code": 400,
			"message":     "账户不存在",
		})
		return
	}

	if loginParam.UserName != "admin" || loginParam.Password != "admin" {
		response.Fail(ctx, nil, gin.H{
			"status_code": 400,
			"message":     "账户或密码错误",
		})
		return
	}

	var msg struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	}

	msg.StatusCode = 200
	msg.Message = "登录成功"

	response.Success(ctx, gin.H{"token": "123456"}, util.Struct2MapViaJson(msg))
}

// 返回所有用户的方法
func GetAllUser(ctx *gin.Context) {
	db := common.GetDB()
	// 获取请求中的所有参数
	query, pagenum, pagesize, param := service.SplitGetAllUserParam(ctx)
	// 根据参数，从数据库中请求user条目
	userList, DefaultLength := dao.SelectAllUser(db, query, pagenum, pagesize, param)

	// 构造返回的结构体
	UserData := model.UsersData{Total: DefaultLength, Pagenum: pagenum, Users: userList}
	Meta := model.Meta{Msg: "获取用户成功", Status_code: 200}

	ctx.JSON(200, gin.H{"data": UserData, "meta": Meta})
}

// 返回特定用户的状态
func GetSpecifiedUser(ctx *gin.Context) {
	db := common.GetDB()
	userID, _ := strconv.Atoi(ctx.Param("userID"))
	// log.Println(userID)
	struct_userList := dao.SelectSpecifiedUser(db, userID)

	// 构造返回的结构体
	Meta := model.Meta{Msg: "获取用户成功", Status_code: 200}

	ctx.JSON(200, gin.H{"data": struct_userList, "meta": Meta})
}

// 更新用户的状态
func PutUserState(ctx *gin.Context) {
	db := common.GetDB()
	mgstate, userID := service.SplitPutUserStateParam(ctx)

	dao.UpdateUserState(db, mgstate, userID)
}

func PutUserInfo(ctx *gin.Context) {
	db := common.GetDB()
	userID, _ := strconv.Atoi(ctx.Param("userID"))

	email := ctx.PostForm("email")
	mobile := ctx.PostForm("mobile")
	log.Println(userID, email, mobile)
	dao.UpdateSpecifiedUser(db, userID, email, mobile)
}

// 新增用户
func PostNewUser(ctx *gin.Context) {
	db := common.GetDB()
	// 提取参数
	username, password, mobile, email, worknum := service.SplitPostNewUserParam(ctx)
	// 写入数据库
	newUser, result := dao.InsertNewUser(db, username, password, mobile, email, worknum)

	log.Println("newUser:", newUser, "result", result.Error, result.RowsAffected)
}

// 删除特定用户
func DeleteSpecifiedUser(ctx *gin.Context) {
	db := common.GetDB()
	userID, _ := strconv.Atoi(ctx.Param("userID"))
	// log.Println(userID)
	struct_userList := dao.DeleteSpecifiedUser(db, userID)

	// 构造返回的结构体
	Meta := model.Meta{Msg: "获取用户成功", Status_code: 200}

	ctx.JSON(200, gin.H{"data": struct_userList, "meta": Meta})
}
