package kugou

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"wechat-robot/db"
	"wechat-robot/request"
	"wechat-robot/respon"
)

type List struct {
	Error   string `json:"error"`
	Errcode int    `json:"errcode"`
	Status  int    `json:"status"`
	Data    struct {
		Info []struct {
			Singername string `json:"singername"` // 歌手名称
			Songname   string `json:"songname"`   // 歌曲名称
			Hash       string `json:"hash"`       // 哈希值
			AlbumId    string `json:"album_id"`   // 专辑id
			AlbumName  string `json:"album_name"` // 专辑名称
		} `json:"info"`
		Timestamp int `json:"timestamp"` // 时间戳
	} `json:"data"`
}

type Detail struct {
	Status   int    `json:"status"`
	ErrCode  int    `json:"err_code"`
	ShowTips string `json:"show_tips"`
	Data     struct {
		AuthorName    string `json:"author_name"`     // 歌手名称
		SongName      string `json:"song_name"`       // 歌曲名称
		Img           string `json:"img"`             // 图片
		PlayUrl       string `json:"play_url"`        // 播放url
		PlayBackupUrl string `json:"play_backup_url"` // 备份播放url

	} `json:"data"`
}

// SearchUrl 酷狗搜索歌曲
// const SearchUrl = "https://complexsearch.kugou.com/v2/search/song?callback=callback123&keyword=%s&page=1&pagesize=30&bitrate=0&isfuzzy=0&inputtype=0&platform=WebFilter&userid=0&clientver=2000&iscorrection=1&privilege_filter=0&filter=10&token=&srcappid=2919&clienttime=1659763719742&mid=1659763719742&uuid=1659763719742&dfid=-&signature=43B6F7C7E9C6CEC82690B017C9263567"
const SearchUrl = "https://msearchcdn.kugou.com/api/v3/search/song?showtype=14&highlight=em&pagesize=1&tag_aggr=1&tagtype=全部&plat=0&sver=5&keyword=%s&correct=1&api_ver=1&version=9108&page=1&area_code=1&tag=1&with_res_tag=1"

// DetailUrl 酷狗音乐详情
const DetailUrl = "https://wwwapi.kugou.com/yy/index.php?r=play/getdata&callback=&hash=%s&album_id=%s"

func Search(keyword string) (res respon.ResponseData, msg string) {
	cacheKey := "MusicSearch:" + keyword
	var list = List{}
	cache, err := db.CacheGet(cacheKey)
	if err != nil {
		var err error
		list, err = list.search(keyword, cacheKey)
		if err != nil {
			list = List{}
			request.ToManagerByErr(err)
		}
	} else {
		_ = json.Unmarshal([]byte(cache), &list)
	}

	if len(list.Data.Info) == 0 {
		return res, "没有搜到歌曲"
	}

	detail := new(Detail)
	detail.get(list.Data.Info[0].Hash, list.Data.Info[0].AlbumId)
	if detail.ErrCode != 0 {
		return res, detail.ShowTips
	}

	if len(detail.Data.PlayUrl) > 0 {
		res.Dataurl = detail.Data.PlayUrl
	} else {
		res.Dataurl = detail.Data.PlayBackupUrl
	}
	res.Desc = detail.Data.AuthorName
	res.Thumburl = detail.Data.Img
	res.Title = detail.Data.SongName
	res.URL = " https://www.kugou.com/song/#hash=" + list.Data.Info[0].Hash
	return res, ""
}

func (d *Detail) get(hash, albumId string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	var req *http.Request
	req, _ = http.NewRequest("GET", fmt.Sprintf(DetailUrl, hash, albumId), nil)
	req.AddCookie(&http.Cookie{Name: "kg_mid", Value: "42cdb5832780da473e67886bf775a335", HttpOnly: true})
	req.AddCookie(&http.Cookie{Name: "kg_dfid", Value: "1txf1h3gMmna29eW3z2Fc5eV", HttpOnly: true})
	res, err := client.Do(req)
	if err != nil {
		request.ToManagerByErr(err)
		return
	}
	defer res.Body.Close()

	// 读取内容
	b, err := io.ReadAll(res.Body)
	_ = json.Unmarshal(b, d)
}

func (list List) search(keyword, cacheKey string) (List, error) {
	// 发送get请求
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(fmt.Sprintf(SearchUrl, keyword))
	if err != nil {
		return list, err
	}
	defer res.Body.Close()

	// 读取内容
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return list, err
	}

	// 返回来有多余的字符，需去除
	resultStr := string(b)
	if strings.HasPrefix(resultStr, "<!--KG_TAG_RES_START-->") {
		s := strings.TrimSpace(resultStr)
		r := []rune(s)
		resultStr = string(r[23:])
	}
	if strings.HasSuffix(resultStr, "<!--KG_TAG_RES_END-->") {
		s := strings.TrimSpace(resultStr)
		length := len([]rune(s))
		r := []rune(s)
		resultStr = string(r[:length-21])
	}

	// json字符串转成结构体
	err = json.Unmarshal([]byte(resultStr), &list)
	if err != nil {
		return list, err
	}
	if list.Errcode != 0 {
		return list, errors.New(list.Error)
	}

	db.CacheSave(cacheKey, resultStr)
	return list, nil
}
