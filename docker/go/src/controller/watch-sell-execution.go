package controller

import (
	"model"
)

// WatchSellExecution 最新の売り注文の約定状況を確認する。
// 約定していれば、modeを0に更新する。
func WatchSellExecution() {
	childOrderID := model.ReadSellOrderInfo(5)
	isExecuted := model.GETExecution(childOrderID.(string))
	if isExecuted {
		model.UpdateTradeMode(0)
	}
}
