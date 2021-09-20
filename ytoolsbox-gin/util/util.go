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
	"math/rand"
	"reflect"
	"time"
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
