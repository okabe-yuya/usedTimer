package main

import (
	"log"
	"used_timer/timer"
	// "used_timer/timer"
)

func main() {
	log.Print("[Main] Start timer ...")
	// 2秒毎にlogを出し、起動から10秒経過した時にbreak
	timer.Scheduler(2, 10, 5)
}
