package main

import (
	"controller"
	"job"
)

func main() {
	// 無限ループでゴルーチンを起動する
	quit := make(chan bool)
	go job.WriteTickerJob()
	go job.AutoTradingJob()
	go controller.RESTAPI()
	// 永遠に返らない
	<-quit

}
