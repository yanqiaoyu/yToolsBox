/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu
 * @Date: 2021-06-22 12:56:47
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-06-22 12:59:10
 * @FilePath: \golang_web\util\util.go
 */
package util

import (
	"encoding/json"
	"fmt"
	"main/dto"
	"main/response"
	"math/rand"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化生成一个种子
	rand.Seed(time.Now().UnixNano())
}

func GetRandomString2(n int) string {
	randBytes := make([]byte, n)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func Struct2MapViaReflect(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Struct2MapViaJson(obj interface{}) map[string]interface{} {
	data, _ := json.Marshal(&obj)
	m := make(map[string]interface{})
	json.Unmarshal(data, &m)
	return m
}

// 封装了一层模型绑定 Param
func ResolveParam(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBind(obj)
	if err != nil {
		msg := dto.FailResponseMeta{StatusCode: 400, Message: "必填字段为空"}
		response.Fail(ctx, nil, Struct2MapViaJson(msg))
		return err
	}
	return nil
}

// 封装了一层模型绑定 Query
func ResolveQuery(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindQuery(obj)
	if err != nil {
		msg := dto.FailResponseMeta{StatusCode: 400, Message: "必填字段为空"}
		response.Fail(ctx, nil, Struct2MapViaJson(msg))
		return err
	}
	return nil
}

// 封装了一层模型绑定 URI
func ResolveURI(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindUri(obj)
	if err != nil {
		msg := dto.FailResponseMeta{StatusCode: 400, Message: "必填字段为空"}
		response.Fail(ctx, nil, Struct2MapViaJson(msg))
		return err
	}
	return nil
}

// 创建文件夹
func CreateDir(folderPath string) (dirPath string) {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.MkdirAll(folderPath, 0777)
		// 再修改权限
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

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
