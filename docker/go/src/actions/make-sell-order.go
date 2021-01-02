package actions

import (
	"model"
	"utils"
)

func MakeSellOrder() {
	lastBuyOrderPrice := model.ReadBuyOrderInfo(3)
	limitPrice, stopPrice := model.CalculateSellOrderPriceByAnyPer(lastBuyOrderPrice.(float64), 5, 10)
	available := utils.RoundDown(model.GETBalance("ETH"), 2)
	parentOrderAcceptionId := model.POSTParentOrder(limitPrice, stopPrice, available)
	parentOrderId := model.GETParentOrderId(parentOrderAcceptionId)
	childOrderId := model.GETChildOrderId(parentOrderId)
	// ToDo:LIMIT注文しか書き込めていない。STOP注文の情報も書き込めるように変更する
	model.WriteSellOrderInfo(limitPrice, available, childOrderId)
	model.UpdateTradeMode(3)
}
