package main

import (
	"demo/apidemo/utils"
	"demo/apidemo/utils/authv3"
	"fmt"
)

// 您的应用ID
var appKey = "60f40320720e64fc"

// 您的应用密钥
var appSecret = "lZzRmRPIbAULfDD6mSSnDvvGCUNQnUaq"

func main() {
	// 添加请求参数
	paramsMap := createRequestParams()
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv3.AddXiaopAuthParams(appKey, appSecret, paramsMap)
	// 请求api服务
	events := utils.DoPostBySSE("https://openapi.youdao.com/llmserver/ai/teacher/dialogue/chat", header, paramsMap)
	for event := range events {
		// 处理接收到的事件
		fmt.Println(event)

	}
}

func createRequestParams() map[string][]string {

	/*
		note: 将下列变量替换为需要请求的参数
		取值参考文档: https://ai.youdao.com/DOCSIRMA/html/trans/api/xpls/index.html
	*/
	userId := "user_test"
	taskName := "你好"
	chatInfo := "[{\"type\":\"text\",\"content\":\"你好，帮我算一下1+1=？\"}]"

	return map[string][]string{
		"user_id":   {userId},
		"task_name": {taskName},
		"chat_info": {chatInfo},
	}
}
