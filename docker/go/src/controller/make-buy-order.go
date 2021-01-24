package controller

import (
	"model"
	"utils"
)

// MakeBuyOrder 買い注文を出すための関数
func MakeBuyOrder() {
	// tickerから高値を取得し。自らの残高情報をもとに、買い注文を発注する。
	high := model.ReadTickerInfo()
	buyOrderPrice := utils.RoundDown(model.CalculateBuyOrderPriceByAnyPer(high, 5), 0)
	available := model.GETBalance("JPY")
	size := utils.RoundDown(model.CheckAvailableOrderSize(buyOrderPrice, available), 3)
	childOrderID := model.POSTChildOrder(buyOrderPrice, size)
	if childOrderID != "" {
		model.WriteBuyOrderInfo(buyOrderPrice, size, childOrderID)
		model.UpdateTradeMode(1)
	}
}
