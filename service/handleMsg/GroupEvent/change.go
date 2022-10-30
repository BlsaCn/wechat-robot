package GroupEvent

import (
	"wechat-robot/request"
	"wechat-robot/service"
)

func ChangeNameHandle() {
	oldName := service.OldName
	newName := service.NewName
	userName := service.FromName
	request.ByText(service.FromGroup, "@"+userName+" 修改了群名称\n【旧名称】"+oldName+"\n【新名称】"+newName)
}
