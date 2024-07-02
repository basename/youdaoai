package main

import (
	"demo/apidemo/utils"
	"demo/apidemo/utils/authv3"
	"encoding/json"
	"fmt"
	"strings"
)

// 您的应用ID
var appKey = "60f40320720e64fc"

// 您的应用密钥
var appSecret = "lZzRmRPIbAULfDD6mSSnDvvGCUNQnUaq"

type Message struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

func main() {
	// 添加请求参数
	questionStr := "1+1=?"

	answerStr := getXPAnswer(questionStr)

	fmt.Println(answerStr)
}

func getXPAnswer(questionStr string) string {
	paramsMap := createRequestParams(questionStr)
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv3.AddXiaopAuthParams(appKey, appSecret, paramsMap)
	// 请求api服务
	events := utils.DoPostBySSE("https://openapi.youdao.com/llmserver/ai/teacher/dialogue/chat", header, paramsMap)

	answer := ""
	for event := range events {
		// 处理接收到的事件
		fmt.Println(event)
		parts := strings.SplitN(event, ":", 2)
		if len(parts) != 2 {
			//fmt.Println("无效的输入格式")
			continue
		}
		jsonData := parts[1]

		var msg Message
		err := json.Unmarshal([]byte(jsonData), &msg)
		if err != nil {
			//fmt.Println("解析JSON出错:", err)
			continue
		}
		fmt.Sprintf("msg:%+v", msg)
		answer += msg.Content
	}
	return answer
}
func createRequestParams(question string) map[string][]string {

	/*
		note: 将下列变量替换为需要请求的参数
		取值参考文档: https://ai.youdao.com/DOCSIRMA/html/trans/api/xpls/index.html
	*/
	userId := "user_test"
	taskName := "你好"

	chatInfo := []Message{}

	tempInfo := Message{
		Type:    "text",
		Content: question,
	}

	chatInfo = append(chatInfo, tempInfo)

	chatInfoStr, _ := json.Marshal(chatInfo)
	inputInfo := string(chatInfoStr)
	return map[string][]string{
		"user_id":   {userId},
		"task_name": {taskName},
		"chat_info": {inputInfo},
	}
}
