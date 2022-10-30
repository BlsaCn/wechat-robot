package constant

// JoinGroupXml 新人进群
const JoinGroupXml = "<?xml version=\"1.0\"?><msg><appmsg appid=\"\" sdkver=\"0\"><title>欢迎%s～</title><des> 邀请人：%s&#x000A;&#x000D;——————————</des><username /><action>view</action><type>5</type><showtype>0</showtype><content /><url>%s</url><lowurl /><forwardflag>0</forwardflag><dataurl /><lowdataurl /><contentattr>0</contentattr><streamvideo><streamvideourl /><streamvideototaltime>0</streamvideototaltime><streamvideotitle /><streamvideowording /><streamvideoweburl /><streamvideothumburl /><streamvideoaduxinfo /><streamvideopublishid /></streamvideo><canvasPageItem><canvasPageXml><![CDATA[]]></canvasPageXml></canvasPageItem><appattach><attachid /><cdnthumburl>3057020100044b304902010002045561986002032f802902044b833cb7020462ee8df3042462393263653863612d353235622d346532312d626434342d6430313231363863353235630204011800030201000405004c52ad00</cdnthumburl><cdnthumbmd5>12517f1c130a37468be531db3d88bc74</cdnthumbmd5><cdnthumblength>13611</cdnthumblength><cdnthumbheight>120</cdnthumbheight><cdnthumbwidth>120</cdnthumbwidth><cdnthumbaeskey>071a95fcef7a71895888e50a6869bca7</cdnthumbaeskey><aeskey>071a95fcef7a71895888e50a6869bca7</aeskey><encryver>1</encryver><fileext /><islargefilemsg>0</islargefilemsg></appattach><extinfo /><androidsource>3</androidsource><thumburl>%s</thumburl><websearch /></appmsg><fromusername>AV___-_-</fromusername><scene>0</scene><appinfo><version>1</version><appname></appname></appinfo><commenturl></commenturl></msg>"

// ShareVideoXml 分享解析的短视频无水印链接
const ShareVideoXml = `<?xml version="1.0"?>
	<msg>
		<appmsg appid="" sdkver="0">
			<title>%s|%s</title>
			<des>作者:%s
评论数:%d
点赞数:%d</des>
			<action>view</action>
			<type>5</type>
			<showtype>0</showtype>
			<content />
			<url>%s</url>
			<dataurl />
			<lowurl />
			<lowdataurl />
			<recorditem><![CDATA[]]></recorditem>
			<thumburl>%s</thumburl>
			<messageaction />
			<extinfo />
			<sourceusername />
			<sourcedisplayname />
			<commenturl />
			<appattach>
				<totallen>0</totallen>
				<attachid />
				<emoticonmd5 />
				<fileext />
				<aeskey />
			</appattach>
			<weappinfo>
				<pagepath />
				<username />
				<appid />
				<appservicetype>0</appservicetype>
			</weappinfo>
			<websearch />
		</appmsg>
		<fromusername>AV___-_-</fromusername>
		<scene>0</scene>
		<appinfo>
			<version>1</version>
			<appname></appname>
		</appinfo>
		<commenturl></commenturl>
	</msg>
	`
