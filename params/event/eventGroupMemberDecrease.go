package event

import "wechat-robot/service"

// GroupMemberDecrease 群成员减少事件
type GroupMemberDecrease struct {
	Event   string `json:"Event"`
	Content struct {
		RobotWxid     string `json:"robot_wxid"`      // 机器人账号id
		FromGroup     string `json:"from_group"`      // 群号
		FromGroupName string `json:"from_group_name"` // 群名
		ToWxid        string `json:"to_wxid"`         // 被操作者wxid
		ToName        string `json:"to_name"`         // 被操作者昵称
		Time          string `json:"time"`            // 操作时间
	} `json:"content"`
}

func (e GroupMemberDecrease) InitParams() {
	service.FromGroup = e.Content.FromGroup
	service.FromGroupName = e.Content.FromGroupName
	service.ToName = e.Content.ToName
	service.FromGroupName = e.Content.FromGroupName
}
