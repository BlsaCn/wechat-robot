package service

// Replied 已经回复过
var Replied = false

// Event 事件名
var Event string

var Type int

// FromGroup 群号
var FromGroup string

// FromGroupName 群名称
var FromGroupName string

// FromWxid 发消息的群成员id
var FromWxid string

// FromName 发消息的群成员昵称
var FromName string

// Msg 消息内容
var Msg string

// EventGroupMemberAdd -------------

// InviterWxid 邀请者账号id
var InviterWxid string

// InviterUsername 邀请者昵称
var InviterUsername string

// GuestWxid 新人账号id
var GuestWxid string

// GuestUsername 新人昵称
var GuestUsername string

// EventGroupMemberDecrease --------------

// ToName 被操作者昵称
var ToName string

// EventGroupNameChange

// NewName 新群名
var NewName string

// OldName 旧群名
var OldName string
