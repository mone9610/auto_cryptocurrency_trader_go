package controller

import (
	"model"
	"time"
	"utils"
)

// MakeSellOrder 売り注文(指値注文・逆指値注文)を出す関数
func MakeSellOrder() {
	lastBuyOrderPrice := model.ReadBuyOrderInfo(3)
	limitPrice, stopPrice := model.CalculateSellOrderPriceByAnyPer(lastBuyOrderPrice.(float64), 5, 10)
	limitPrice = utils.RoundDown(limitPrice, 0)
	stopPrice = utils.RoundDown(stopPrice, 0)
	available := utils.RoundDown(model.GETBalance("ETH"), 3)
	parentOrderAcceptionID := model.POSTParentOrder(limitPrice, stopPrice, available)
	parentOrderID := model.GETParentOrderID(parentOrderAcceptionID)
	// 親注文を出してすぐだと子注文IDが取得できないのでsleepする
	time.Sleep(time.Second * 2)
	childOrderID := model.GETChildOrderID(parentOrderID)
	if childOrderID != "" {
		// ToDo:LIMIT注文しか書き込めていない。STOP注文の情報も書き込めるように変更する
		model.WriteSellOrderInfo(limitPrice, available, childOrderID)
		model.UpdateTradeMode(3)
	}
}
