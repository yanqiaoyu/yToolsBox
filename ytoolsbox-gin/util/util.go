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
	"fmt"
	"math/rand"
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
