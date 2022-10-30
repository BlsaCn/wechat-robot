package controller

import (
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	// const xml2 = "<?xml version=\"1.0\"?><msg><appmsg appid=\"\" sdkver=\"0\"><title>欢迎%s～</title><des> 邀请人：%s&#x000A;&#x000D;———————————☆玩得开心☆</des><username /><action>view</action><type>5</type><showtype>0</showtype><content /><url>%s</url><lowurl /><forwardflag>0</forwardflag><dataurl /><lowdataurl /><contentattr>0</contentattr><streamvideo><streamvideourl /><streamvideototaltime>0</streamvideototaltime><streamvideotitle /><streamvideowording /><streamvideoweburl /><streamvideothumburl /><streamvideoaduxinfo /><streamvideopublishid /></streamvideo><canvasPageItem><canvasPageXml><![CDATA[]]></canvasPageXml></canvasPageItem><appattach><attachid /><cdnthumburl>3057020100044b304902010002045561986002032f802902044b833cb7020462ee8df3042462393263653863612d353235622d346532312d626434342d6430313231363863353235630204011800030201000405004c52ad00</cdnthumburl><cdnthumbmd5>12517f1c130a37468be531db3d88bc74</cdnthumbmd5><cdnthumblength>13611</cdnthumblength><cdnthumbheight>120</cdnthumbheight><cdnthumbwidth>120</cdnthumbwidth><cdnthumbaeskey>071a95fcef7a71895888e50a6869bca7</cdnthumbaeskey><aeskey>071a95fcef7a71895888e50a6869bca7</aeskey><encryver>1</encryver><fileext /><islargefilemsg>0</islargefilemsg></appattach><extinfo /><androidsource>3</androidsource><thumburl>%s</thumburl><websearch /></appmsg><fromusername>AV___-_-</fromusername><scene>0</scene><appinfo><version>1</version><appname></appname></appinfo><commenturl></commenturl></msg>"
	// xml := fmt.Sprintf(xml2, "新人", "邀请人", "http://wx.qlogo.cn/mmhead/ver_1/iaTkdrME9wOa0VJIlcqs0DzMfObqbClfZbQHs2wcfEhPw86asYAoFRQPozxhFtDgny1pnOxIicicbLbaV1XNaweVYkoOtw3MbjlsVrJo5hqP4A/0", "http://wx.qlogo.cn/mmhead/ver_1/iaTkdrME9wOa0VJIlcqs0DzMfObqbClfZbQHs2wcfEhPw86asYAoFRQPozxhFtDgny1pnOxIicicbLbaV1XNaweVYkoOtw3MbjlsVrJo5hqP4A/0")
	// request.ByLinkMsg("19401949239@chatroom", xml)

	// const xml3 = "<?xml version=\"1.0\"?><msg><appmsg appid=\"\" sdkver=\"0\"><title>@段九。|一直变好才有底气留住爱我的你们#风里有了秋的味道 #生活碎片 #纯情女高</title><des>作者:程Yooooo 评论数:14095点赞数:562524</des><action>view</action><type>5</type><showtype>0</showtype><content /><url>http://v6-cdn-tos.ppxvod.com/389358e464bb4eafde2ebc54ba5deeac/63580d54/video/tos/cn/tos-cn-ve-0076/65827b50c65941a4a6eb3809aaea621a/?a=1319&ch=0&cr=0&dr=3&cd=0%7C0%7C0%7C0&cv=1&br=1588&bt=1588&cs=0&ds=3&eid=2048&ft=FIJbvNN6V~6wbLMFq8dzJLeOYZlcm2zdd2bLN5ljtiZm&mime_type=video_mp4&qs=0&rc=Zjw6OWc2Nzk2OzlkNGdlZ0BpanFla2Q6ZnJyZzMzNGYzM0AtMzEwYWIzNjYxYF5jYWJiYSNuMWlhcjRfLTJgLS1kMWFzcw%3D%3D&l=202210252322340102080650441512EAAC</url><dataurl /><lowurl /><lowdataurl /><recorditem><![CDATA[]]></recorditem><thumburl>https://p26.douyinpic.com/aweme/1080x1080/aweme-avatar/mosaic-legacy_31991000995363770746a.jpeg?from=116350172</thumburl><messageaction /><extinfo /><sourceusername /><sourcedisplayname /><commenturl /><appattach><totallen>0</totallen><attachid /><emoticonmd5 /><fileext /><aeskey /></appattach><weappinfo><pagepath /><username /><appid /><appservicetype>0</appservicetype></weappinfo><websearch /></appmsg><fromusername>AV___-_-</fromusername><scene>0</scene><appinfo><version>1</version><appname></appname></appinfo><commenturl></commenturl></msg>"
	// request.ByLinkMsg("19401949239@chatroom", xml3)

	// xmlStr := fmt.Sprintf(
	// 	constant.ShareVideoXml, "@段九。", "标题标题标题标题", "作者作者",
	// 	11, 22,
	// 	// "http://v6-cdn-tos.ppxvod.com/389358e464bb4eafde2ebc54ba5deeac/63580d54/video/tos/cn/tos-cn-ve-0076/65827b50c65941a4a6eb3809aaea621a/?a=1319&ch=0&cr=0&dr=3&cd=0%7C0%7C0%7C0&cv=1&br=1588&bt=1588&cs=0&ds=3&eid=2048&ft=FIJbvNN6V~6wbLMFq8dzJLeOYZlcm2zdd2bLN5ljtiZm&mime_type=video_mp4&qs=0&rc=Zjw6OWc2Nzk2OzlkNGdlZ0BpanFla2Q6ZnJyZzMzNGYzM0AtMzEwYWIzNjYxYF5jYWJiYSNuMWlhcjRfLTJgLS1kMWFzcw%3D%3D&l=202210252322340102080650441512EAAC",
	// 	"https://cn-gddg-ct-01-10.bilivideo.com/upgcxcode/72/42/873574272/873574272-1-208.mp4?e=ig8euxZM2rNcNbh1nWdVhwdlhzRHhwdVhoNvNC8BqJIzNbfq9rVEuxTEnE8L5F6VnEsSTx0vkX8fqJeYTj_lta53NCM=&uipk=5&nbs=1&deadline=1667137654&gen=playurlv2&os=bcache&oi=1921313500&trid=0000cc0119c48e194f35986b5fb6d853e2b1T&mid=0&platform=html5&upsig=deea14ac9415a8eff661d56e131c893f&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,mid,platform&cdnid=61310&bvc=vod&nettype=0&bw=319998&orderid=0,1&logo=80000000",
	// 	"https://p26.douyinpic.com/aweme/1080x1080/aweme-avatar/mosaic-legacy_31991000995363770746a.jpeg?from=116350172",
	// )
	// // 机器人测试群2
	// request.ByLinkMsg("19401949239@chatroom", xmlStr)
}
