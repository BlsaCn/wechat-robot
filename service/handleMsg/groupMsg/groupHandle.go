package groupMsg

import (
	"fmt"
	"github.com/BlsaCn/parse-video-url-go/parse"
	"strconv"
	"strings"
	"wechat-robot/constant"
	"wechat-robot/request"
	"wechat-robot/service"
	"wechat-robot/service/group"
	"wechat-robot/service/handleMsg/common"
	"wechat-robot/utils"
)

func Handle() {
	if service.Type == constant.System {
		return
	}

	if service.Type == constant.Text {
		textMsg()
		return
	}
}

// 撤回消息
func revokedMsg() {
	// params.Msg
	//	$revokedMsg = $this->msg['revoked_msg'];
	//	// 文本消息
	//	if ($revokedMsg['msg_type'] === 1) {
	//		$msg = " 撤回了一条消息：" . PHP_EOL . $this->msg['revoked_msg']['content'];
	//		return SendGroupMsgAndAt::getInstance()->send(
	//		$msg,
	//		$this->fromWxid,
	//		$this->fromWxid,
	//		$this->finalFromWxid,
	//		$this->finalFromName
	//		);
	//	}
	//
	//	// 动态表情
	//	if ($revokedMsg['msg_type'] === 47) {
	//		// $param = [
	//		//     "event" => "SendEmojiMsg",
	//		//     "robot_wxid" => 'wxid_i7r3xpsyks2241',
	//		//     "to_wxid" => '19401949239@chatroom',
	//		//     "msg" => [
	//		//         "url" => $this->msg['revoked_msg']['content'],
	//		//         "name" => '202111191551197602.gif'
	//		//     ]
	//		// ];
	//		return SendEmojiMsg::getInstance()->send(
	//		$this->msg['revoked_msg']['content'],
	//		$this->fromWxid,
	//		$this->fromWxid,
	//	);
	//	}
}

// 文本消息
func textMsg() {
	common.TextMsg()
	if service.Replied {
		return
	}

	if strings.HasPrefix(service.Msg, "群成员列表") {
		service.Replied = true
		var m group.Member
		err := m.List()
		if err != nil {
			request.ToManagerByErr(err)
			return
		}

		msg := "本群一共" + strconv.Itoa(m.ReturnJson.Count) + "人\n"
		manCount := 0
		womanCount := 0
		unknownCount := 0
		memberStr := ""
		for k, item := range m.ReturnJson.MemberList {
			nickname := ""
			if len(item.GroupNickname) > 0 {
				nickname = item.GroupNickname
			} else {
				nickname = item.Nickname
			}
			if item.Sex == 1 {
				manCount += 1
			} else if item.Sex == 0 {
				unknownCount += 1
			} else {
				womanCount += 1
			}
			memberStr += strconv.Itoa(k+1) + "、" + nickname + "\n"
		}
		msg += "男性有" + strconv.Itoa(manCount) + "人\n"
		msg += "女性有" + strconv.Itoa(womanCount) + "人\n"
		msg += "未知性别有" + strconv.Itoa(unknownCount) + "人\n"
		msg += memberStr

		request.ByText(service.FromGroup, msg)
		return
	}

	if utils.IncludeUrl(service.Msg) {
		videoInfo, err := parse.ByMsg(service.Msg)
		if err == nil {
			xmlStr := fmt.Sprintf(
				constant.ShareVideoXml, "@"+service.FromName, videoInfo.Title, videoInfo.AuthorName,
				videoInfo.CommentNum, videoInfo.LikeNum, videoInfo.VideoUrl, videoInfo.CoverUrl,
			)
			request.ByLinkMsg(service.FromGroup, xmlStr)
			return
		}
	}
}
