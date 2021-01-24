package controller

import (
	"model"
)

// WatchBuyExecution 最新の買い注文の約定状況を確認する。
func WatchBuyExecution() {
	// 約定していれば、modeを2に更新する。
	ChildOrderID := model.ReadBuyOrderInfo(5)
	isExecuted := model.GETExecution(ChildOrderID.(string))
	if isExecuted {
		model.UpdateTradeMode(2)
	}
}
