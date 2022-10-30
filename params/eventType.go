package params

type BaseEvent[T any] struct {
	Event   string `json:"Event"`
	Content T      `json:"content"`
}

// EventDeviceCallback 设备回调事件
type EventDeviceCallback struct {
	RobotWxid string `json:"robot_wxid"` // 机器人账号id
	Type      int    `json:"type"`       // 消息类型
	Msg       string `json:"msg"`        // 消息内容
	ToWxid    string `json:"to_wxid"`    // 接收用户ID
	ToName    string `json:"to_name"`    // 接收用户昵称
	MsgId     string `json:"msg_id"`     // 消息ID
}

// EventGroupChat 群消息事件
type EventGroupChat struct {
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
}

// EventGroupMemberAdd 群成员增加事件 PS: 新人进群
type EventGroupMemberAdd struct {
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
}

// EventGroupMemberDecrease 群成员减少事件
type EventGroupMemberDecrease struct {
	RobotWxid     string `json:"robot_wxid"`      // 机器人账号id
	FromGroup     string `json:"from_group"`      // 群号
	FromGroupName string `json:"from_group_name"` // 群名
	ToWxid        string `json:"to_wxid"`         // 被操作者wxid
	ToName        string `json:"to_name"`         // 被操作者昵称
	Time          string `json:"time"`            // 操作时间
}

// EventGroupNameChange 群名称变动事件
type EventGroupNameChange struct {
	RobotWxid string `json:"robot_wxid"` // 机器人账号id
	FromGroup string `json:"from_group"` // 群号
	NewName   string `json:"new_name"`   // 新群名
	OldName   string `json:"old_name"`   // 旧群名
	FromName  string `json:"from_name"`  // 操作者昵称
}

// EventPrivateChat 私聊消息事件
type EventPrivateChat struct {
	RobotWxid string `json:"robot_wxid"` // 机器人账号id
	Type      int    `json:"type"`       // 1/文本消息 3/图片消息 34/语音消息  42/名片消息  43/视频 47/动态表情 48/地理位置  49/分享链接  2001/红包  2002/小程序  2003/群邀请  更多请参考常量表
	FromWxid  string `json:"from_wxid"`  // 来源用户ID
	FromName  string `json:"from_name"`  // 来源用户昵称
	Msg       string `json:"msg"`        // 消息内容
	MsgId     string `json:"msg_id"`     // 消息ID
}
