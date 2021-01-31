package job

import (
	"controller"
	"model"
	"utils"

	"github.com/carlescere/scheduler"
)

// AutoTradingJob 仮想通貨を自動取引するための関数
func AutoTradingJob() {
	job := func() {
		tradeMode := model.CheckTradeMode()
		isReady := model.ReadIsReady()
		if isReady == 1 {
			if tradeMode == 0 {
				utils.LogUtil("Mode is 0,now func MakeBuyOrder is going to be executed.", 0)
				controller.MakeBuyOrder()
			} else if tradeMode == 1 {
				utils.LogUtil("Mode is 1,now func WatchBuyExecution is going to be executed.", 0)
				controller.WatchBuyExecution()
			} else if tradeMode == 2 {
				utils.LogUtil("Mode is 2,now func MakeSellOrder is going to be executed.", 0)
				controller.MakeSellOrder()
			} else if tradeMode == 3 {
				utils.LogUtil("Mode is 3,now func WatchSellExecution is going to be executed.", 0)
				controller.WatchSellExecution()
			}
		} else {
			utils.LogUtil("Your access_key and secret_key are not ready,open port:80 and input your keys. ", 1)
		}
	}
	scheduler.Every(10).Seconds().Run(job)
}
