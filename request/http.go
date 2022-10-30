package request

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"unsafe"
	"wechat-robot/config"
)

// HttpRequest 发送http请求给服务端，请求数据
func HttpRequest(data map[string]interface{}) (string, error) {
	data["token"] = config.Token
	b, err := json.Marshal(&data)
	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest("POST", config.HttpServer, reader)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	cline := http.Client{}
	resp, err := cline.Do(req)
	if err != nil {
		return "", err
	}

	// 获取返回结果
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	logrus.Println("返回数据：" + (*str))
	return *str, nil
}
