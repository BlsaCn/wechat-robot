package group

import (
	"encoding/json"
	"wechat-robot/request"
	"wechat-robot/service"
)

// Member 接口返回的数据返回
type Member struct {
	Code       int    `json:"Code"`
	Result     string `json:"Result"`
	ReturnJson struct {
		GroupWxid     string `json:"group_wxid"` // 群id
		GroupName     string `json:"group_name"` // 群名称
		Count         int    `json:"count"`      // 总人数
		OwnerWxid     string `json:"owner_wxid"`
		OwnerNickname string `json:"owner_nickname"`
		MemberList    []struct {
			WxNum         string `json:"wx_num"`             // 微信号
			Avatar        string `json:"avatar"`             // 头像
			City          string `json:"city,omitempty"`     // 城市
			Country       string `json:"country"`            // 国家
			GroupNickname string `json:"group_nickname"`     // 在群的昵称
			Nickname      string `json:"nickname,omitempty"` // 昵称
			Province      string `json:"province"`           // 省份
			Remark        string `json:"remark"`             // 备注
			Sex           int    `json:"sex"`                // 性别
			Wxid          string `json:"wxid"`               // 微信id
		} `json:"member_list"`
	} `json:"ReturnJson"`
}

// List 群成员列表
func (m *Member) List() error {
	data := map[string]interface{}{
		"api":        "GetGroupMember",
		"robot_wxid": "wxid_i7r3xpsyks2241",
		"group_wxid": service.FromGroup,
		"is_refresh": "1",
	}
	str, err := request.HttpRequest(data)
	if err != nil {
		request.ToManagerByErr(err)
		return err
	}
	err = json.Unmarshal([]byte(str), &m)
	if err != nil {
		return err
	}
	if m.Code != 0 {
		request.ToManagerByErr(m.Result)
	}
	return nil
}
