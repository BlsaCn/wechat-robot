package group

import (
	"encoding/json"
	"wechat-robot/request"
	"wechat-robot/service"
)

// MemberDetail 接口返回的数据返回
type MemberDetail struct {
	Code       int    `json:"Code"`
	Result     string `json:"Result"`
	ReturnJson struct {
		Wxid             string `json:"wxid"`
		WxNum            string `json:"wx_num"`
		Nickname         string `json:"nickname"`   // 昵称
		Headimgurl       string `json:"headimgurl"` // 头像
		Country          string `json:"country"`
		Province         string `json:"province"`
		City             string `json:"city"`
		Sex              int    `json:"sex"`
		Scene            int    `json:"scene"`
		Signature        string `json:"signature"`
		Backgroundimgurl string `json:"backgroundimgurl"`
		GroupWxid        string `json:"group_wxid"`
		GroupName        string `json:"group_name"`
	} `json:"ReturnJson"`
}

// Get 某个群成员详情
func (m *MemberDetail) Get(toWxid string) error {
	data := map[string]interface{}{
		"api":         "GetGroupMemberDetailInfo",
		"robot_wxid":  "wxid_i7r3xpsyks2241",
		"group_wxid":  service.FromGroup,
		"member_wxid": toWxid,
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
