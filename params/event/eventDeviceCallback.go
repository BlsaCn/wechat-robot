package event

import (
	"strings"
	"wechat-robot/service"
)

// DeviceCallback 设备回调事件
type DeviceCallback struct {
	Event   string `json:"Event"`
	Content struct {
		RobotWxid string `json:"robot_wxid"` // 机器人账号id
		Type      int    `json:"type"`       // 消息类型
		Msg       string `json:"msg"`        // 消息内容
		ToWxid    string `json:"to_wxid"`    // 接收用户ID
		ToName    string `json:"to_name"`    // 接收用户昵称
		MsgId     string `json:"msg_id"`     // 消息ID
	} `json:"content"`
}

func (e DeviceCallback) InitParams() {
	service.Type = e.Content.Type
	service.Msg = strings.TrimSpace(e.Content.Msg)
}
