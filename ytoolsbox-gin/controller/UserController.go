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
	"main/dto"
	"main/model"
	"main/response"

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

// 判断登陆的账号是否合法
func isUserExist() {

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, "", gin.H{"user": dto.ToUserDTO(user.(model.User))})
}

func Login(ctx *gin.Context) {
	// 正常来讲，需要进数据库
	// 但是目前我们先不入库，直接返回

	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	// log.Print("Login Name is:\n", name, "Login Password is:\n", password)

	// 账户名密码错误
	if password != "admin" || name != "admin" {
		data := map[string]interface{}{
			"data": map[string]interface{}{
				"status_code": 401,
				"message":     "账号密码错误",
			},
		}
		ctx.JSON(200, data)
		return
	}

	data := map[string]interface{}{
		"data": map[string]interface{}{
			"status_code": 200,
			"message":     "登录成功",
			"token":       "123456",
		},
	}

	ctx.JSON(200, data)

}
