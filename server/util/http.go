package util

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err // 返回空字符串和错误
	}
	defer resp.Body.Close() // 确保关闭响应体流

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err // 读取响应体失败，返回错误
	}

	return string(body), nil // 将响应体转换为字符串并返回
}
