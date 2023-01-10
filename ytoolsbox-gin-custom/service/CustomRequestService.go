package service

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"main/utils"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/gin-gonic/gin"
)

func JudgeWhetherContainsNeedResponseData(ctx *gin.Context) []byte {
	// 拿到URL中的加密数据
	needResponseData := ctx.Query("needResponseData")
	needResponseData = strings.ReplaceAll(needResponseData, " ", "+")

	if needResponseData == "" {
		log.Println("URL没有需要返回的数据")
		return []byte{}
	} else {
		log.Println("URL中的加密数据: ", needResponseData)
		needResponseBytes, _ := base64.StdEncoding.DecodeString(needResponseData)
		// 转换为字符串
		needResponseString := utils.BytesToString(needResponseBytes)
		log.Printf("base64解码结果为: %s", needResponseString)
		return needResponseBytes
	}

}

// 尝试Json
func TryJsonUnmarshal(needResponseBytes []byte) (map[string]interface{}, error) {
	// 尝试JSON反序列化的时候不需要make，被封装到Unmarshal中了
	var m map[string]interface{}

	err := json.Unmarshal(needResponseBytes, &m)
	if err != nil {
		log.Println("JSON反序列化失败: ", err)
		return nil, err
	}
	return m, nil
}

func TryXMLUnmarshal(needResponseBytes []byte) (string, error) {
	doc, err := xmlquery.Parse(strings.NewReader(utils.BytesToString(needResponseBytes)))
	if err != nil || doc.OutputXML(true) == "" {
		log.Println("XML反序列化失败: ", err)
		return "", err
	}
	log.Println(doc.OutputXML(true))
	return doc.OutputXML(true), nil
}

// func safety(d string) []byte {
// 	var buffer bytes.Buffer
// 	for _, c := range d {
// 		s := string(c)
// 		if c == 92 { // 92 is a backslash
// 			continue
// 		}
// 		if unicode.IsPrint(c) {
// 			buffer.WriteString(s)
// 		} else {
// 			buffer.WriteString(url.QueryEscape(s))
// 		}
// 	}
// 	return buffer.Bytes()
// }
