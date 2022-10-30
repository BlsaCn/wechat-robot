package common

import (
	"strings"
	"wechat-robot/package/kugou"
	"wechat-robot/request"
	"wechat-robot/service"
)

func TextMsg() {
	if strings.HasPrefix(service.Msg, "点歌") || strings.HasPrefix(service.Msg, "来首") {
		service.Replied = true
		r := []rune(service.Msg)
		keyword := string(r[2:])
		if keyword == "" {
			request.ByText(service.FromGroup, "指令错误，如：点歌孤勇者")
			return
		}
		res, msg := kugou.Search(keyword)
		if len(msg) > 0 {
			request.ByText(service.FromGroup, msg)
			return
		}

		request.ByMusicLink(service.FromGroup, res)
		return
	}
}
