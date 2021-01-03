package controller

import (
	"model"
)

func WatchSellExecution() {
	// 最新の売り注文の約定状況を確認する。
	// 約定していれば、modeを0に更新する。
	childOrderId := model.ReadSellOrderInfo(5)
	isExecuted := model.GETExecution(childOrderId.(string))
	if isExecuted {
		model.UpdateTradeMode(0)
	}
}
