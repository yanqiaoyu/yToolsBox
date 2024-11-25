package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/linkedin/goavro/v2"
	"log"
	"main/cons"
	"main/dspm"
	"strings"
	"time"
)

type ParamStruct struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

var SensitiveMap map[string][]string

var produceCodec *goavro.Codec

func init() {
	produceCodec = GetAvroSchemaStrCodec(cons.HttpSourceSchemaStr)
	json.Unmarshal(GetBytes(cons.SensitiveStr), &SensitiveMap)
}

func GetAvroSchemaStrCodec(schemaStr string) *goavro.Codec {
	// 创建 Avro 编解码器
	codec, err := goavro.NewCodec(schemaStr)
	if err != nil {
		log.Fatal(err)
	}
	return codec
}

func getStartTime(reduceDay int) int64 {
	hisDate := time.Now().AddDate(0, 0, -1*reduceDay)
	timeUnix := time.Date(hisDate.Year(), hisDate.Month(), hisDate.Day(), 0, 10, 0, 0, hisDate.Location()).Unix()
	return timeUnix
}
func GetBytes(str string) []byte {
	buffer := bytes.NewBufferString(str)
	byteSlice := buffer.Bytes()
	return byteSlice
}
func GetBaseData(ctx *gin.Context, recordTimestamp int64) map[string]interface{} {
	resultMap := make(map[string]interface{})
	err := json.Unmarshal(GetBytes(cons.HttpSourceLogStr), &resultMap)
	resultMap["recordTimestamp"] = recordTimestamp
	resultMap["insertTimestamp"] = recordTimestamp
	resultMap["uploadTimestamp"] = recordTimestamp
	if err != nil {
		log.Println("获取基础数据错误:", err)
	}
	uuId := uuid.New().String()
	resultMap["uuId"] = uuId

	resultMap["dstIp"] = cons.PocAddr
	uri := ctx.Request.RequestURI
	resultMap["uri"] = uri
	resultMap["url"] = uri

	reqHeader := strings.Replace(resultMap["requestHead"].(string), "/api/MarketingStatus/UpdateStatus", uri, -1)
	resultMap["requestHead"] = reqHeader
	srcIp := ctx.ClientIP()
	if srcIp == "::1" {
		srcIp = "192.168.1.1"
	}
	resultMap["srcIp"] = srcIp
	extensionsMap := make(map[string]interface{})
	extensionsMap["string"] = fmt.Sprintf("{\"xSrcIp\":\"%v\"}", srcIp)
	resultMap["extensions"] = extensionsMap

	return resultMap
}

func getSensitiveBody(num int) string {
	ResBodyMap := make(map[string]string)
	countNum := 0
	for key, valueSlice := range SensitiveMap {
		if num == countNum {
			break
		}
		for index, value := range valueSlice {
			sensitiveKey := fmt.Sprintf("%s_%v", key, index)
			ResBodyMap[sensitiveKey] = value
		}
		countNum++
	}
	var resBodySlice []ParamStruct
	for key, value := range ResBodyMap {
		paramStruct := ParamStruct{key, value}
		resBodySlice = append(resBodySlice, paramStruct)
	}
	marshal, _ := json.Marshal(resBodySlice)
	return string(marshal)
}

func sendLogs(logs []map[string]interface{}, restartFlink bool) {
	go func() {
		if restartFlink {
			dspm.RestartFlink()
		}
		var resultSlice []*pulsar.ProducerMessage
		for _, logMap := range logs {
			binaryData, err := produceCodec.BinaryFromNative(nil, logMap)
			if err != nil {
				panic(err)
			}
			producerMessage := &pulsar.ProducerMessage{Payload: binaryData}
			resultSlice = append(resultSlice, producerMessage)
		}

		client := dspm.CreateClient()
		producer := dspm.CreateProducer(client, cons.HttpSourceTopic)
		dspm.SendDataBatchToPulsar(producer, resultSlice)
		time.Sleep(10 * time.Second)
		defer client.Close()
		defer producer.Close()
	}()
}
