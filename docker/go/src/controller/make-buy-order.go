package controller

import (
	"model"
	"utils"
)

func MakeBuyOrder() {
	// tickerから高値を取得し。自らの残高情報をもとに、買い注文を発注する。
	high := model.ReadTickerInfo()
	buyOrderPrice := utils.RoundDown(model.CalculateBuyOrderPriceByAnyPer(high, 5), 0)
	available := model.GETBalance("JPY")
	size := model.CheckAvailableOrderSize(buyOrderPrice, available)
	childOrderId := model.POSTChildOrder(buyOrderPrice, size)
	model.WriteBuyOrderInfo(buyOrderPrice, size, childOrderId)
	model.UpdateTradeMode(1)
}
