package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// ProcessPlayerName 处理玩家名称
func ProcessPlayerName(playerName string) string {
	runes := []rune(playerName)

	if len(runes) > 9 {
		return string(runes[0]) + strings.Repeat("*", 7) + string(runes[len(runes)-1])
	} else if len(runes) == 2 {
		return string(runes[0]) + "*" + string(runes[1])
	} else if len(runes) > 2 {
		return string(runes[0]) + strings.Repeat("*", len(runes)-2) + string(runes[len(runes)-1])
	}
	return playerName
}

// HTTPGet 发送GET请求并解析JSON响应
func HTTPGet(url string) (map[string]interface{}, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HTTPGetWithHeaders 发送带自定义headers的GET请求并解析JSON响应
func HTTPGetWithHeaders(url string, headers map[string]string) (map[string]interface{}, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 添加headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
