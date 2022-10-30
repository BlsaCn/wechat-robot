package request

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"wechat-robot/config"
	"wechat-robot/respon"
)

// ToManager 发送消息给管理员
func ToManager(msg string) {
	ByText(config.Manager, msg)
}

// ToManagerByErr 发送错误消息给管理员
func ToManagerByErr(err interface{}) {
	msg := ""
	switch v := err.(type) {
	case string:
		msg = v
		break
	case error:
		msg = v.Error()
		break
	}
	ByText(config.Manager, msg)
}

// ByMusicLink 发送一条可播放的歌曲链接
func ByMusicLink(toWxid string, musicReq respon.ResponseData) {
	data := map[string]interface{}{
		"api":        "SendMusicLinkMsg",
		"robot_wxid": config.Robot,
		"to_wxid":    toWxid,
		"title":      musicReq.Title,
		"desc":       musicReq.Desc,
		"url":        musicReq.URL,
		"dataurl":    musicReq.Dataurl,
		"thumburl":   musicReq.Thumburl,
	}
	send(data)
}

// ByText 发送文本消息
func ByText(toWxid, msg string) {
	data := map[string]interface{}{
		"api":        "SendTextMsg",
		"robot_wxid": config.Robot,
		"to_wxid":    toWxid,
		"msg":        msg,
	}
	send(data)
}

// ByLinkMsg 发送链接消息
func ByLinkMsg(toWxid, xml string) {
	data := map[string]interface{}{
		"api":        "SendLinkMsg",
		"robot_wxid": config.Robot,
		"to_wxid":    toWxid,
		"xml":        xml,
	}
	send(data)
}

// 发送动作
func send(data gin.H) {
	data["token"] = config.Token
	b, err := json.Marshal(&data)
	if err != nil {
		logrus.Println("json.Marshal", data, err)
		return
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest("POST", config.HttpServer, reader)
	if err != nil {
		logrus.Println("http.NewRequest", err)
		return
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	cline := http.Client{}
	_, err = cline.Do(req)
	if err != nil {
		logrus.Println("cline.Do", err)
	}
}
