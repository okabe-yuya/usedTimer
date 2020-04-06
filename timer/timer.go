package timer

import (
	"errors"
	"log"
	"time"
)

// 仕様について
// 1. 指定した周期毎に" "というメッセージをchannel経由で送信する
// 2. 指定した周期毎に時間切れとなり、errorを返して実行loopをbreakする
func UsedTimer(iterTimeInt, breakTimeInt int, ptc chan *Protocol) {
	// int型をtime.Duration型へ変換(second)
	iterTime := time.Duration(iterTimeInt) * time.Second
	breakTime := time.Duration(breakTimeInt) * time.Second

	// 周期毎にイベントを発生させるtickerを作成
	iterTicker := time.NewTicker(iterTime)
	breakTicker := time.NewTicker(breakTime)
	defer func() {
		iterTicker.Stop()
		breakTicker.Stop()
	}()

	for {
		select {
		// メッセージを送信する処理
		case <-iterTicker.C:
			ptc <- &Protocol{
				Type:    "MESSAGE",
				Message: "Golangはいいぞ",
			}
		case <-breakTicker.C:
			ptc <- &Protocol{
				Type:  "ERROR",
				Error: errors.New("[Error] Timer was broken"),
			}
			// breakだとスレッドがkillできないので注意
			return
		}
	}
}

// timerの動作管理を行うスケジューラー。再起動まで行う
func Scheduler(iterTimeInt, breakTimeInt, revivalNum int) {
	// 通信用のチャネルを作成
	ch := CreateProtocolChannel()
	// 何度も再起動させたいので簡略化のため無名関数を採用
	receiver := func() error {
		// 作成したUsedTimerからのメッセージを受け取るloopを作成
		for {
			select {
			case msg := <-ch:
				switch msg.Type {
				case "MESSAGE":
					log.Print(msg.Message)
				case "ERROR":
					log.Print(msg.Error.Error())
					return msg.Error
				}
			}
		}
	}

	// 監視のための処理を記述
	counter := 0
	for {
		// goroutineを用いてスレッドを作成
		go UsedTimer(iterTimeInt, breakTimeInt, ch)
		if err := receiver(); err != nil {
			log.Print("[Main] Restart timer thread. Please wait 3 seconds ...")
			time.Sleep(3 * time.Second)
			counter += 1
		}

		if counter >= revivalNum {
			log.Println("[Main] Reached revival limit. so stop worker ...")
			return
		}
	}
}
