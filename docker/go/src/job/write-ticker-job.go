package job

import (
	"fmt"
	"model"

	"github.com/carlescere/scheduler"
)

func WriteTickerJob() {
	job := func() {
		fmt.Println("WriteTickerJob Start")
		var high, last, low float64
		high, last, low = model.GETTicker()
		fmt.Println(high, last, low)
		ins := model.WriteTickerInfo(high, last, low)
		// ToDo:エラーログ等をファイルに書き出す想定
		if ins == false {
			fmt.Println("WriteTickerInfo failed")
		}
	}
	scheduler.Every(10).Minutes().Run(job)
}
