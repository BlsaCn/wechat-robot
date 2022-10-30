package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wechat-robot/constant"
	"wechat-robot/params"
	"wechat-robot/service"
	"wechat-robot/service/handleMsg/GroupEvent"
	"wechat-robot/service/handleMsg/friendMsg"
	"wechat-robot/service/handleMsg/groupMsg"
)

func Notice(c *gin.Context) {
	err := params.InitEvent(c)
	if err != nil {
		fmt.Println("出现错误:", err)
		return
	}

	// 群聊
	if service.Event == constant.EventGroupChat {
		groupMsg.Handle()
		return
	}

	// ---------------------------- 事件
	// 群成员减少事件
	if service.Event == constant.EventGroupMemberDecrease {
		GroupEvent.DecreaseHandle()
		return
	}
	// 群成员增加事件
	if service.Event == constant.EventGroupMemberAdd {
		GroupEvent.AddHandle()
		return
	}
	// 群名称变动事件
	if service.Event == constant.EventGroupNameChange {
		GroupEvent.ChangeNameHandle()
		return
	}
	// 私聊
	if service.Event == constant.EventPrivateChat {
		friendMsg.Handle()
		return
	}
}
