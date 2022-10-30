package params

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"wechat-robot/params/event"
	"wechat-robot/service"
)

// eventParams 事件，不同事件不同参数
type eventParams struct {
	Event string `json:"Event"` // 事件
}

var eventMap = map[string]EventInterface{
	"EventGroupChat":           event.GroupChat{},
	"EventDeviceCallback":      event.DeviceCallback{},
	"EventGroupMemberAdd":      event.GroupMemberAdd{},
	"EventGroupMemberDecrease": event.GroupMemberDecrease{},
	"EventGroupNameChange":     event.GroupNameChange{},
	"EventPrivateChat":         event.PrivateChat{},
}

func InitEvent(c *gin.Context) error {
	p := &eventParams{}
	err := c.ShouldBindBodyWith(&p, binding.JSON)
	if err != nil {
		return err
	}

	service.Event = p.Event

	// 对应事件的结构体
	var EventStruct = eventMap[p.Event]
	switch eventInstance := EventStruct.(type) {
	case event.DeviceCallback:
	case event.GroupMemberAdd:
		_ = c.ShouldBindBodyWith(&eventInstance, binding.JSON)
		eventInstance.InitParams()
	case event.GroupMemberDecrease:
		_ = c.ShouldBindBodyWith(&eventInstance, binding.JSON)
		eventInstance.InitParams()
	case event.GroupNameChange:
		_ = c.ShouldBindBodyWith(&eventInstance, binding.JSON)
		eventInstance.InitParams()
	case event.PrivateChat:
		_ = c.ShouldBindBodyWith(&eventInstance, binding.JSON)
		eventInstance.InitParams()
	case event.GroupChat:
		_ = c.ShouldBindBodyWith(&eventInstance, binding.JSON)
		eventInstance.InitParams()
	default:
		return errors.New(p.Event + "事件不存在")
	}
	if service.Msg == "" || service.FromGroup == "" {
		return errors.New("类型异常")
	}
	return nil
}
