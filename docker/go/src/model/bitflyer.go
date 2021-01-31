package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"utils"
)

const bitFlyerEndopoint = "https://api.bitflyer.com"

// TickerParent : bitFlyerAPIから取得するtickerデータの構造体
type TickerParent struct {
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"` //最低売価格
	BestAsk         float64 `json:"best_ask"` //最高買価格
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     int     `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   int     `json:"total_ask_depth"`
	MarketBidSize   int     `json:"market_bid_size"`
	MarketAskSize   int     `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"` //最終取引価格
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

type ParentOrderResponse struct {
	ParentOrderAcceptanceID string `json:"parent_order_acceptance_id"`
}

// ChildOrderResponse 子注文のIDを格納するための構造体
type ChildOrderResponse struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}

// BalanceResponse bitFlyerLightnintAPIから残高情報を取得するための構造体
type BalanceResponse struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Available    float64 `json:"available"`
}

type Executions struct {
	ID                     int     `json:"id"`
	ChildOrderID           string  `json:"child_order_id"`
	Side                   string  `json:"side"`
	Price                  float64 `json:"price"`
	Size                   float64 `json:"size"`
	Commission             float64 `json:"commission"`
	ExecDate               string  `json:"exec_date"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
}

type ParentOrder struct {
	ID            int    `json:"id"`
	ParentOrderID string `json:"parent_order_id"`
	OrderMethod   string `json:"order_method"`
	ExpireDate    string `json:"expire_date"`
	TimeInForce   string `json:"time_in_force"`
	Parameters    []struct {
		ProductCode   string  `json:"product_code"`
		ConditionType string  `json:"condition_type"`
		Side          string  `json:"side"`
		Price         float64 `json:"price"`
		Size          float64 `json:"size"`
		TriggerPrice  float64 `json:"trigger_price"`
		Offset        float64 `json:"offset"`
	} `json:"parameters"`
	ParentOrderAcceptanceID string `json:"parent_order_acceptance_id"`
}

type ChildOrder struct {
	ID                     int     `json:"id"`
	ChildOrderID           string  `json:"child_order_id"`
	ProductCode            string  `json:"product_code"`
	Side                   string  `json:"side"`
	ChildOrderType         string  `json:"child_order_type"`
	Price                  float64 `json:"price"`
	AveragePrice           float64 `json:"average_price"`
	Size                   float64 `json:"size"`
	ChildOrderState        string  `json:"child_order_state"`
	ExpireDate             string  `json:"expire_date"`
	ChildOrderDate         string  `json:"child_order_date"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
	OutstandingSize        float64 `json:"outstanding_size"`
	CancelSize             float64 `json:"cancel_size"`
	ExecutedSize           float64 `json:"executed_size"`
	TotalCommission        float64 `json:"total_commission"`
}

//Tickerから最新の最高価格、最終価格、最低価格を入手するための関数
// 引数なし、戻り値はtickerから取得した最高価格、最終取引価格、最低価格
func GETTicker() (hi, la, lo float64) {
	path := "/v1/getticker?product_code=ETH_JPY"
	url := bitFlyerEndopoint + path
	// リクエストを定義
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	// リクエストを送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	var ticker TickerParent
	json.Unmarshal([]byte(byteArray), &ticker)

	// 最高価格、最終価格、最低価格を取得する
	var high = ticker.BestAsk
	var last = ticker.Ltp
	var low = ticker.BestBid
	return high, last, low
}

// POSTParentOrder 親注文を出すための関数。基本的には売り注文のみで使う。
// 指値の売り注文価格と、逆指値の売り注文価格を引数とし、order_idを返り値とする。
func POSTParentOrder(limitPrice, stopPrice, size float64) string {
	// アクセスキー、シークレットキーをdbから読み取る
	accessKey, secretKey := ReadConf()
	// キー付きでsha256で署名を作成
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	method := "POST"
	path := "/v1/me/sendparentorder"

	// 親注文のパラメータを定義
	var paramArray []map[string]interface{}
	limitParam := map[string]interface{}{
		"product_code":   "ETH_JPY",
		"condition_type": "LIMIT",
		"side":           "SELL",
		"price":          limitPrice,
		"size":           size,
	}
	stopParam := map[string]interface{}{
		"product_code":   "ETH_JPY",
		"condition_type": "STOP",
		"side":           "SELL",
		"trigger_price":  stopPrice,
		"size":           size,
	}
	paramArray = append(paramArray, limitParam)
	paramArray = append(paramArray, stopParam)

	body := map[string]interface{}{
		"order_method":     "OCO",
		"minute_to_expire": 43200,
		"time_in_force":    "GTC",
		"parameters":       paramArray,
	}
	fmt.Println(body)
	sbody := utils.MapToString(body)
	text := timestamp + method + path + sbody
	sign := utils.MakeHMAC(text, secretKey)

	// リクエストヘッダを作成する
	header := map[string]string{
		"ACCESS-KEY":       accessKey,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}

	// リクエストを送信する
	url := bitFlyerEndopoint + path
	req, err := utils.NewRequest(method, url, header, []byte(sbody))
	if err != nil {
		utils.LogUtil(err, 1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	// レスポンスボディをstringへ変換して返り値とする
	byteArray, _ := ioutil.ReadAll(res.Body)
	sres := string(byteArray)
	jsonBytes := ([]byte)(sres)
	data := new(ParentOrderResponse)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		utils.LogUtil(err, 1)
		return ""
	}
	return string(data.ParentOrderAcceptanceID)
}

// POSTChildOrder 子注文を出すための関数。基本的には買い注文のみで利用する。
// 指値の買い注文価格を引数として、order_idを返り値とする
func POSTChildOrder(price, size float64) string {
	// アクセスキー、シークレットキーをdbから読み取る
	accessKey, secretKey := ReadConf()

	//sha256で署名を作成
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	method := "POST"
	path := "/v1/me/sendchildorder"
	body := map[string]interface{}{
		"product_code":     "ETH_JPY",
		"child_order_type": "LIMIT",
		"side":             "BUY",
		"price":            price,
		"size":             size,
		"minute_to_expire": 10000,
		"time_in_force":    "GTC",
	}
	sbody := utils.MapToString(body)
	text := timestamp + method + path + sbody
	sign := utils.MakeHMAC(text, secretKey)

	// リクエストヘッダを作成する
	header := map[string]string{
		// ToDo:ACCESS-KEYをDBに格納した上で変数に代入できるようにする
		"ACCESS-KEY":       accessKey,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}

	// リクエストを送信する
	url := bitFlyerEndopoint + path
	req, err := utils.NewRequest(method, url, header, []byte(sbody))
	if err != nil {
		utils.LogUtil(err, 1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.LogUtil(err, 1)
	}

	// レスポンスボディをstringへ変換して返り値とする
	byteArray, _ := ioutil.ReadAll(res.Body)
	sres := string(byteArray)
	jsonBytes := ([]byte)(sres)
	data := new(ChildOrderResponse)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		utils.LogUtil(err, 1)
		return ""
	}
	return string(data.ChildOrderAcceptanceID)
}

// GETExecution 約定履歴を取得するための関数。
// order_idを引数として、boolを返り値とする。
// ToDo:引数を指定した上で約定履歴を獲得できるようにする。
func GETExecution(childOrderID string) bool {
	coi := childOrderID
	if coi == "" {
		utils.LogUtil("childOrderID is null", 1)
		return false
	}

	// アクセスキー、シークレットキーをdbから読み取る
	accessKey, secretKey := ReadConf()

	// キー付きでsha256で署名を作成
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	method := "GET"
	path := "/v1/me/getexecutions?product_code=ETH_JPY&child_order_acceptance_id=" + coi
	text := timestamp + method + path
	sign := utils.MakeHMAC(text, secretKey)

	// リクエストヘッダを作成する
	header := map[string]string{
		// ToDo:ACCESS-KEYをDBに格納した上で変数に代入できるようにする
		"ACCESS-KEY":       accessKey,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}

	// リクエストを送信する
	url := bitFlyerEndopoint + path
	req, err := utils.NewRequest(method, url, header, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	// レスポンスボディを読み取り、件数が0件ならfalseを戻り値とする
	byteArray, err := ioutil.ReadAll(res.Body)
	var data []Executions
	err = json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if len(data) == 0 {
		fmt.Println("No data")
		return false
	}
	return true
}

// GETBalance 残高を取得するための関数
// クエリパラメータで通過の種類を引数とし、残高を返り値とする
func GETBalance(currencyCode string) (available float64) {
	cc := currencyCode

	// アクセスキー、シークレットキーをdbから読み取る
	accessKey, secretKey := ReadConf()
	// sha256で署名を作成
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	method := "GET"
	path := "/v1/me/getbalance"
	text := timestamp + method + path
	sign := utils.MakeHMAC(text, secretKey)

	// リクエストヘッダを作成する
	header := map[string]string{
		// ToDo:ACCESS-KEYをDBに格納した上で変数に代入できるようにする
		"ACCESS-KEY":       accessKey,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}

	// リクエストを送信する
	url := bitFlyerEndopoint + path
	req, err := utils.NewRequest(method, url, header, nil)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.LogUtil(err, 1)
	}

	// 引数をキーとして指定した通貨コードに紐づく利用可能残高を取得する
	byteArray, err := ioutil.ReadAll(res.Body)
	var data []BalanceResponse
	err = json.Unmarshal(byteArray, &data)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	for i := range data {
		if data[i].CurrencyCode == cc {
			return data[i].Amount
		} else {
			continue
		}
	}
	return 0
}

// GETParentOrderId 親注文受付Idをキーにして親注文Idを取得するための関数
// ParentOrderAcceptanceIDを引数として、ParentOrderIdを戻り値とする
func GETParentOrderID(parentOrderAcceptionID string) string {
	poai := parentOrderAcceptionID

	// アクセスキー、シークレットキーをdbから読み取る
	accessKey, secretKey := ReadConf()

	// キー付きでsha256で署名を作成
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	method := "GET"
	path := "/v1/me/getparentorder?parent_order_acceptance_id=" + poai
	text := timestamp + method + path
	sign := utils.MakeHMAC(text, secretKey)

	// リクエストヘッダを作成する
	header := map[string]string{
		// ToDo:ACCESS-KEYをDBに格納した上で変数に代入できるようにする
		"ACCESS-KEY":       accessKey,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}

	// リクエストを送信する
	url := bitFlyerEndopoint + path
	req, err := utils.NewRequest(method, url, header, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	// レスポンスボディを読み取り、親注文Idを取得する
	byteArray, err := ioutil.ReadAll(res.Body)
	var data ParentOrder
	err = json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return data.ParentOrderID
}

// GETChildOrderID 　親注文Idをキーにして子注文Idを取得するための関数
// ParentOrderIdを引数として、ChildOrderAcceptionIdを戻り値とする
func GETChildOrderID(parentOrderID string) string {
	poi := parentOrderID

	// 親注文IDがnullの場合、子注文IDを取得しない
	if poi == "" {
		utils.LogUtil("parentOrderID is null", 1)
		return ""
	}

	// アクセスキー、シークレットキーをdbから読み取る
	accessKey, secretKey := ReadConf()

	// キー付きでsha256で署名を作成
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	method := "GET"
	path := "/v1/me/getchildorders?product_code=ETH_JPY&parent_order_id=" + poi
	text := timestamp + method + path
	sign := utils.MakeHMAC(text, secretKey)

	// リクエストヘッダを作成する
	header := map[string]string{
		// ToDo:ACCESS-KEYをDBに格納した上で変数に代入できるようにする
		"ACCESS-KEY":       accessKey,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}

	// リクエストを送信する
	url := bitFlyerEndopoint + path
	req, err := utils.NewRequest(method, url, header, nil)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.LogUtil(err, 1)
	}

	// レスポンスボディを読み取り、件数が0件ならfalseを戻り値とする
	byteArray, err := ioutil.ReadAll(res.Body)
	var data []ChildOrder
	err = json.Unmarshal(byteArray, &data)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	if len(data) == 0 {
		utils.LogUtil("Can't get chiledOrderID", 1)
	}
	return data[0].ChildOrderAcceptanceID
}
