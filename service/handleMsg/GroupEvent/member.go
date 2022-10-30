package GroupEvent

import (
	"fmt"
	"wechat-robot/constant"
	"wechat-robot/request"
	"wechat-robot/service"
	"wechat-robot/service/group"
)

// AddHandle 进群事件
func AddHandle() {
	// 获取邀请者信息
	inviterWxid := service.InviterWxid
	var inviterDetail group.MemberDetail
	err := inviterDetail.Get(inviterWxid)
	var inviterName string
	if err == nil {
		inviterName = inviterDetail.ReturnJson.Nickname
	} else {
		inviterName = service.InviterUsername
	}

	// 获取新人信息
	var detail group.MemberDetail
	err = detail.Get(service.GuestWxid)
	if err != nil {
		request.ByText(service.FromGroup, "@"+service.GuestUsername+"欢迎进群")
		return
	}

	// 发送链接
	img := detail.ReturnJson.Headimgurl
	nickname := detail.ReturnJson.Nickname
	xml := fmt.Sprintf(constant.JoinGroupXml, nickname, inviterName, img, img)
	request.ByLinkMsg(service.FromGroup, xml)
}

// DecreaseHandle 退群事件
func DecreaseHandle() {
	userName := service.ToName
	groupName := service.FromGroupName
	request.ByText(service.FromGroup, "@"+userName+"退出了【"+groupName+"】")
}
