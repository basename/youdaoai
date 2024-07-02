package main

import (
	"demo/apidemo/utils"
	"demo/apidemo/utils/authv3"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

type Message struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

func main() {

	// 打开一个已存在的Excel文件
	f, err := excelize.OpenFile("./example.xlsx")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	// 假设数据从第二行开始（第一行是标题行）
	// 遍历每一行
	for row := 2; row <= 300; row++ {
		// 读取第二列的值
		cellValue, err := f.GetCellValue("Sheet1", fmt.Sprintf("B%d", row))
		if err != nil || len(cellValue) <= 0 {
			print(err)
			break
			// 如果读取错误，可能是到达文件末尾
			log.Fatalf("Failed to get cell value: %v", err)
		}

		//// 对读取的值进行处理（这里是将字符串转换为大写）
		//processedValue := strings.ToUpper(cellValue)

		//添加请求参数
		questionStr := cellValue

		answerStr := getXPAnswer(questionStr)
		fmt.Println(cellValue, answerStr)

		// 将处理后的值写入到第三列
		err = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), answerStr)
		if err != nil {
			log.Fatalf("Failed to set cell value: %v", err)
		}
	}

	// 保存修改后的文件
	err = f.SaveAs("modified_example.xlsx")
	if err != nil {
		log.Fatalf("Failed to save file: %v", err)
	}

	fmt.Println("Excel file processed successfully.")
}

func getXPAnswer(questionStr string) string {
	// 您的应用ID
	var appKey = "60f40320720e64fc"

	// 您的应用密钥
	var appSecret = "lZzRmRPIbAULfDD6mSSnDvvGCUNQnUaq"

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
