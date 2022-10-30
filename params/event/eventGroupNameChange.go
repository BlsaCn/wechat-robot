package event

import (
	"wechat-robot/service"
)

// GroupNameChange 群名称变动事件
type GroupNameChange struct {
	Event   string `json:"Event"`
	Content struct {
		RobotWxid string `json:"robot_wxid"` // 机器人账号id
		FromGroup string `json:"from_group"` // 群号
		NewName   string `json:"new_name"`   // 新群名
		OldName   string `json:"old_name"`   // 旧群名
		FromName  string `json:"from_name"`  // 操作者昵称
	} `json:"content"`
}

func (e GroupNameChange) InitParams() {
	service.FromGroup = e.Content.FromGroup
	service.FromName = e.Content.FromName
	service.NewName = e.Content.NewName
	service.OldName = e.Content.OldName
}
