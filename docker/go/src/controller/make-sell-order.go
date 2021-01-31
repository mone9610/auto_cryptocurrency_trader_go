package controller

import (
	"model"
	"utils"
)

// MakeSellOrder 売り注文(指値注文・逆指値注文)を出す関数
func MakeSellOrder() {
	lastBuyOrderPrice := model.ReadBuyOrderInfo(3)
	limitPrice, stopPrice := model.CalculateSellOrderPriceByAnyPer(lastBuyOrderPrice.(float64), 5, 10)
	available := model.GETBalance("ETH")
	size := utils.RoundDown(model.CheckAvailableOrderSize(limitPrice, available), 3)
	parentOrderAcceptionID := model.POSTParentOrder(limitPrice, stopPrice, size)
	parentOrderID := model.GETParentOrderID(parentOrderAcceptionID)
	childOrderID := model.GETChildOrderID(parentOrderID)
	if childOrderID != "" {
		// ToDo:LIMIT注文しか書き込めていない。STOP注文の情報も書き込めるように変更する
		model.WriteSellOrderInfo(limitPrice, available, childOrderID)
		model.UpdateTradeMode(3)
	}
}
