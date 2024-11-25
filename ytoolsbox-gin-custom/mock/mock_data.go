package mock

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/cons"
	"main/dspm"
	"time"
)

//1.清理ck表的数据
//2.清理pg表的数据
//3.清理pulsar的数据
//4.重启flink
//5.构造json文件
//6.发送数据到pulsar

// PersonGetAppDataMoreThanBaseLine 个人访问任一应用的数据量超过历史基线
func PersonGetAppDataMoreThanBaseLine(ctx *gin.Context) {
	dspm.ClearCkModelData()
	dspm.ClearPgModelData("")
	var starTime int64
	var resultMapSlice []map[string]interface{}
	for i := 0; i < 9; i++ {
		starTime = getStartTime(8 - i)
		num := 2
		logNum := 30
		if i > 6 {
			num = 7
			logNum = 100
		}
		resBody := getSensitiveBody(num)
		for j := 0; j < logNum; j++ {
			baseData := GetBaseData(ctx, starTime)
			baseData["responseBody"] = resBody
			resultMapSlice = append(resultMapSlice, baseData)
		}
	}
	sendLogs(resultMapSlice, false)
	dspm.ClearPgModelData("")
}

// IPReturnsLargeAmountSensitiveDataInShortTime ip短时间下载了大量的数据
func IPReturnsLargeAmountSensitiveDataInShortTime(ctx *gin.Context) string {
	var starTime int64 = time.Now().Unix()
	var resultMapSlice []map[string]interface{}
	resBody := getSensitiveBody(7)
	for i := 1001; i > 0; i-- {
		if i == 1 {
			starTime = starTime + 3600*2
		}
		baseData := GetBaseData(ctx, starTime)
		baseData["responseBody"] = resBody
		resultMapSlice = append(resultMapSlice, baseData)
	}
	sendLogs(resultMapSlice, true)
	return resBody
}

// ParameterIteration 参数遍历
func ParameterIteration(ctx *gin.Context) string {
	mid := ""
	dspm.ClearCkModelData()
	dspm.ClearPgModelData(mid)

	var starTime int64 = time.Now().Unix()
	var resultMapSlice []map[string]interface{}
	resBodyResult := "{\"data\":\"登录失败\"}"
	for i := 1001; i > 0; i-- {
		if i == 1 {
			starTime = starTime + 3600*2
		}
		reqBody := fmt.Sprintf("{\"name\":\"admin\",\"password\":\"admin\",\"check_code\":\"%v\"}", i)
		baseData := GetBaseData(ctx, starTime)
		baseData["requestBody"] = reqBody
		if i < 5 {
			resBodyResult = "{\"data\":\"登录成功\"}"
		}
		baseData["responseBody"] = resBodyResult

		resultMapSlice = append(resultMapSlice, baseData)
		if len(resultMapSlice) == 5000 {
			resultMapSlice = resultMapSlice[:0]
		}
	}
	sendLogs(resultMapSlice, false)
	return resBodyResult
}

// AccountShare 账号共享
func AccountShare(ctx *gin.Context) string {
	dspm.ClearCkModelData()
	var starTime int64 = time.Now().Unix()
	var resultMapSlice []map[string]interface{}
	resBodyResult := "{\"data\":\"登录成功\"}"
	data, _ := ctx.GetRawData()
	jsonString := string(data)

	for i := 1001; i > 0; i-- {
		if i == 1 {
			starTime = starTime + 3600*2
		}
		baseData := GetBaseData(ctx, starTime)
		baseData["srcIp"] = fmt.Sprintf("192.168.1.%d", i%50)
		baseData["requestBody"] = jsonString
		resultMapSlice = append(resultMapSlice, baseData)
	}
	sendLogs(resultMapSlice, true)
	return resBodyResult
}

// AppAccountMoreThanBaseLine 应用的账号数超过历史基线
func AppAccountMoreThanBaseLine(ctx *gin.Context) string {
	dspm.ClearCkModelData()
	dspm.ClearPgModelData("")
	var starTime int64 = time.Now().Unix()
	var resultMapSlice []map[string]interface{}
	resBodyResult := "{\"data\":\"访问成功\"}"

	for i := 0; i < 9; i++ {
		starTime = getStartTime(8 - i)
		for j := 0; j < 50; j++ {
			baseData := GetBaseData(ctx, starTime)
			baseData["requestBody"] = fmt.Sprintf("{\"name\":\"%s\","+
				"\"password\":\"randpassword@123456\",\"check_code\":\"9876\"}", cons.AccountSlice[0])
			if i > 6 {
				baseData["requestBody"] = fmt.Sprintf("{\"name\":\"%s\","+
					"\"password\":\"randpassword@123456\",\"check_code\":\"9876\"}", cons.AccountSlice[j%10])
			}
			resultMapSlice = append(resultMapSlice, baseData)
		}
	}
	sendLogs(resultMapSlice, false)
	return resBodyResult
}
