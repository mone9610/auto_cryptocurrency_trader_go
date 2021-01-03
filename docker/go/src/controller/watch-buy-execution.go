package controller

import (
	"model"
)

func WatchBuyExecution() {
	// 最新の買い注文の約定状況を確認する。
	// 約定していれば、modeを2に更新する。
	childOrderId := model.ReadBuyOrderInfo(5)
	isExecuted := model.GETExecution(childOrderId.(string))
	if isExecuted {
		model.UpdateTradeMode(2)
	}
}
