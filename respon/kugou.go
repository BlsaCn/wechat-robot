package respon

type ResponseData struct {
	Dataurl  string `json:"dataurl"`  // mp3地址
	Desc     string `json:"desc"`     // 内容
	Thumburl string `json:"thumburl"` // http图片地址
	Title    string `json:"title"`    // 标题
	URL      string `json:"url"`      // 链接地址
}
