package job

import (
	"model"
	"utils"

	"github.com/carlescere/scheduler"
)

func WriteTickerJob() {
	job := func() {
		var high, last, low float64
		high, last, low = model.GETTicker()
		ins := model.WriteTickerInfo(high, last, low)
		if ins == false {
			utils.LogUtil("WriteTickerJob failed.", 1)
		}
	}
	scheduler.Every(10).Minutes().Run(job)
}
