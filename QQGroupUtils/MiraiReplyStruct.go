package QQGroupUtils

type MiraiReplyStruct struct {
	RetCode int        `json:"retcode"`
	Message string     `json:"message"`
	Data    DataStruct `json:"data"`
}

type DataStruct struct {
	NickName string `json:"nickname"`
	Card     string `json:"card"`
	//其他内容忽略 并不需要
}
