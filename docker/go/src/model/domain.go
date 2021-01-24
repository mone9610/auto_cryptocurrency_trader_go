package model

// CalculateBuyOrderPriceByAnyPer 指値で買い注文を出す際の価格を分析する関数
// 最高取引価格、any%を引数とし、指値注文での価格を戻り値とする
// any%ルールを適用した場合の価格とする
func CalculateBuyOrderPriceByAnyPer(high, rate float64) float64 {
	r := 1 - (rate / 100)
	buyOrderPrice := high * r
	return buyOrderPrice
}

// CalculateSellOrderPriceByAnyPer 売値の注文価格を決める関数
// 引数は買い注文価格とany%。戻り値は、売値の指値注文の価格、逆指値注文の価格とする。
func CalculateSellOrderPriceByAnyPer(lastBuyOrderPrice, limitRate, stopRate float64) (limitPrice, stopPrice float64) {
	lbop := lastBuyOrderPrice
	lR := 1 + (limitRate / 100)
	sR := 1 - (stopRate / 100)
	limitPrice = lbop * lR
	stopPrice = lbop * sR
	return limitPrice, stopPrice
}

// CheckAvailableOrderSize 注文可能数量をチェックするための関数
// 価格と利用可能数量を引数とし、数量を返り値とする
func CheckAvailableOrderSize(price, available float64) float64 {
	bop := price
	a := available
	size := a / bop
	return size
}
