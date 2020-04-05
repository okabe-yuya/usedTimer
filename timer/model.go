package timer

// スレッド間で用いるchannel用のstruct定義
type Protocol struct {
	// 失敗時のmessageなのかを判定するために用意
	Type    string
	Message string
	Error   error
}

// 構造体を元にチャネルを作成
func CreateProtocolChannel() chan *Protocol {
	return make(chan *Protocol)
}
