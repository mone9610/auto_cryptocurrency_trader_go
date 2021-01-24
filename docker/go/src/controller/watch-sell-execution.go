package controller

import (
	"model"
)

// WatchSellExecution 最新の売り注文の約定状況を確認する。
func WatchSellExecution() {
	// 約定していれば、modeを0に更新する。
	childOrderID := model.ReadSellOrderInfo(5)
	isExecuted := model.GETExecution(childOrderID.(string))
	if isExecuted {
		model.UpdateTradeMode(0)
	}
}
