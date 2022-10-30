package event

import (
	"strings"
	"wechat-robot/service"
)

// GroupChat 群消息事件
type GroupChat struct {
	Event   string `json:"Event"` // 事件（易语言模板的子程序名）
	Content struct {
		RobotWxid     string `json:"robot_wxid"`      // 机器人账号id
		Type          int    `json:"type"`            // 消息类型	见：下方
		FromGroup     string `json:"from_group"`      // 来源群号
		FromGroupName string `json:"from_group_name"` // 来源群名称
		FromWxid      string `json:"from_wxid"`       // 具体发消息的群成员id
		FromName      string `json:"from_name"`       // 具体发消息的群成员昵称
		Msg           string `json:"msg"`             // 消息内容
		// 附带JSON属性（群消息有艾特人员时，返回被艾特信息）
		MsgSource struct {
			Atuserlist []struct {
				Wxid         string `json:"wxid"`
				Nickname     string `json:"nickname"`
				PositionFrom int    `json:"position_from"`
				PositionTo   int    `json:"position_to"`
			} `json:"atuserlist"`
		} `json:"msg_source"`

		MsgId string `json:"msg_id"` // 消息ID
	} `json:"content"`
}

func (e GroupChat) InitParams() {
	service.Type = e.Content.Type
	service.FromGroup = e.Content.FromGroup
	service.FromGroupName = e.Content.FromGroupName
	service.FromWxid = e.Content.FromWxid
	service.FromName = e.Content.FromName
	service.Msg = strings.TrimSpace(e.Content.Msg)
}
