package event

import "wechat-robot/service"

// PrivateChat 私聊消息事件
type PrivateChat struct {
	Event   string `json:"Event"` // 事件
	Content struct {
		RobotWxid string `json:"robot_wxid"` // 机器人账号id
		Type      int    `json:"type"`       // 1/文本消息 3/图片消息 34/语音消息  42/名片消息  43/视频 47/动态表情 48/地理位置  49/分享链接  2001/红包  2002/小程序  2003/群邀请  更多请参考常量表
		FromWxid  string `json:"from_wxid"`  // 来源用户ID
		FromName  string `json:"from_name"`  // 来源用户昵称
		Msg       string `json:"msg"`        // 消息内容
		MsgId     string `json:"msg_id"`     // 消息ID
	} `json:"content"`
}

func (e PrivateChat) InitParams() {
	service.Type = e.Content.Type
	service.FromWxid = e.Content.FromWxid
	service.FromName = e.Content.FromName
	service.Msg = e.Content.Msg
}
