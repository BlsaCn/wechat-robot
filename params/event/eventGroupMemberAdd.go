package event

import (
	"wechat-robot/service"
)

// GroupMemberAdd 群成员增加事件 PS: 新人进群
type GroupMemberAdd struct {
	Event   string `json:"Event"`
	Content struct {
		RobotWxid     string `json:"robot_wxid"`      // 机器人账号id
		FromGroup     string `json:"from_group"`      // 群号
		FromGroupName string `json:"from_group_name"` // 群名
		Guest         struct {
			Wxid     string `json:"wxid"`     // 新人账号id
			Username string `json:"username"` // 新人昵称
		} `json:"guest"`
		Inviter struct {
			Wxid     string `json:"wxid"`     // 邀请者账号id
			Username string `json:"username"` // 邀请者昵称
		} `json:"inviter"`
	} `json:"content"`
}

func (e GroupMemberAdd) InitParams() {
	service.FromGroup = e.Content.FromGroup
	service.FromGroupName = e.Content.FromGroupName
	service.InviterWxid = e.Content.Inviter.Wxid
	service.InviterUsername = e.Content.Inviter.Username
	service.GuestWxid = e.Content.Guest.Wxid
	service.GuestUsername = e.Content.Guest.Username
}
