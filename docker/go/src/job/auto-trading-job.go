package job

import (
	"actions"
	"fmt"
	"model"

	"github.com/carlescere/scheduler"
)

func AutoTradingJob() {
	job := func() {
		tradeMode := model.CheckTradeMode()
		if tradeMode == 0 {
			fmt.Println("Mode is 0,now func MakeBuyOrder is going to be executed.")
			actions.MakeBuyOrder()
		} else if tradeMode == 1 {
			fmt.Println("Mode is 1,now func WatchBuyExecution is going to be executed.")
			actions.WatchBuyExecution()
		} else if tradeMode == 2 {
			fmt.Println("Mode is 2,now func makeSellOrder is going to be executed.")
			actions.MakeSellOrder()
		} else if tradeMode == 3 {
			fmt.Println("Mode is 3,now func makeSellOrder is going to be executed.")
			actions.WatchSellExecution()
		}
	}
	scheduler.Every(10).Seconds().Run(job)
}
