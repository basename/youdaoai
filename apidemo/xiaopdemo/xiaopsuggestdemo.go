package main

//
//import (
//	"demo/apidemo/utils"
//	"demo/apidemo/utils/authv3"
//	"fmt"
//)
//
//// 您的应用ID
//var appKey = ""
//
//// 您的应用密钥
//var appSecret = ""
//
//func main() {
//	// 添加请求参数
//	paramsMap := createRequestParams()
//	header := map[string][]string{
//		"Content-Type": {"application/x-www-form-urlencoded"},
//	}
//	// 添加鉴权相关参数
//	authv3.AddXiaopAuthParams(appKey, appSecret, paramsMap)
//	// 请求api服务
//	result := utils.DoPost("https://openapi.youdao.com/llmserver/plugin/suggest", header, paramsMap, "application/json")
//	// 打印返回结果
//	if result != nil {
//		fmt.Print(string(result))
//	}
//}
//
//func createRequestParams() map[string][]string {
//
//	/*
//		note: 将下列变量替换为需要请求的参数
//		取值参考文档: https://ai.youdao.com/DOCSIRMA/html/trans/api/xpls/index.html
//	*/
//	userId := "user_test"
//	query := "背诵静夜思"
//	answer := "床前明月光，疑是地上霜。举头望明月，低头思故乡。"
//
//	return map[string][]string{
//		"user_id": {userId},
//		"query":   {query},
//		"answer":  {answer},
//	}
//}
