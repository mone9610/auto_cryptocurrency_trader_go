package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"utils"

	_ "github.com/go-sql-driver/mysql"
)

const dbInfo = "root:root@tcp(mysql:3306)/auto_trader"

// TickerChild DB書き込み用のtickerデータの構造体
type tickerChild struct {
	high float64
	low  float64
	last float64
}

//TickerInfo トレンド分析用に利用するtickerデータの構造体
type TickerInfo struct {
	ID        int
	High      float64
	Last      float64
	Low       float64
	CreatedAt string
}

// // tickerHistoryの構造体を格納するための配列
// type TickerInfoArray []TickerInfo

//BuyOrder 買い注文の情報を格納するための構造体
type buyOrder struct {
	price        float64
	size         float64
	childOrderID string
}

// buyOrderInfo 買い注文の情報を取得するための構造体
type buyOrderInfo struct {
	id              int
	orderType       string
	price           float64
	size            float64
	childOrderID    string
	executionStatus int
	createdAt       string
	updatedAt       string
}

// 売り注文の情報を格納するための構造体
type sellOrder struct {
	price        float64
	size         float64
	childOrderID string
}

// sellOrderInfo 売り注文の情報を取得するための構造体
type sellOrderInfo struct {
	id              int
	orderType       string
	price           float64
	size            float64
	childOrderID    string
	executionStatus int
	createdAt       string
	updatedAt       string
}

// Mode 自動取引の状態を表すための構造体
type Mode struct {
	id        int
	tradeMode int
	createdAt string
	updatedAt string
}

// Conf アクセスキー、シークレットキーを設定するための構造体
type Conf struct {
	ID        int
	AccessKey string
	SecretKey string
	IsReady   int
}

//WriteTickerInfo tickerの最高価格、最低価格、最終価格をDBに書き込む関数
// 引数は最高価格、最低価格、最終価格とし、戻り値はboolとする
func WriteTickerInfo(high, low, last float64) bool {
	var data = new(tickerChild)
	*data = tickerChild{high, low, last}

	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		utils.LogUtil(err, 1)
		return false
	}
	defer db.Close()

	ins, err := db.Prepare("INSERT INTO ticker_info (high,last,low) VALUES(?,?,?)")
	if err != nil {
		utils.LogUtil(err, 1)
		return false
	}
	ins.Exec(data.high, data.last, data.low)
	defer ins.Close()
	return true
}

// CheckTradeMode 自動取引jobにて、モードを確認するための関数
// 引数なし、戻り値はモードとする
func CheckTradeMode() int {
	var m Mode
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	defer db.Close()

	// ToDo:全件検索は効率が悪いのであとで直す
	db.QueryRow("SELECT * FROM mode").Scan(&m.id, &m.tradeMode, &m.createdAt, &m.updatedAt)

	if err != nil {
		utils.LogUtil(err, 1)
	}
	return m.tradeMode
}

// ReadTickerInfo tikcer_infoテーブルから、最新の最高取引価格を取り出す関数
// 引数なし、戻り値は最新の取引価格(float64)
func ReadTickerInfo() float64 {
	var ti TickerInfo
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	defer db.Close()

	// ToDo:全件検索は効率が悪いのであとで直す
	db.QueryRow("SELECT * FROM ticker_info ORDER BY created_at DESC LIMIT 1").Scan(&ti.ID, &ti.High, &ti.Last, &ti.Low, &ti.CreatedAt)

	if err != nil {
		utils.LogUtil(err, 1)
	}
	return ti.High
}

// ReadIsReady アクセスキー、シークレットキーが設定済みか確認する関数
// 引数なし、isReadyを戻り値とする(未:0 済:1)
func ReadIsReady() (IsReady int) {
	var conf Conf
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	defer db.Close()

	db.QueryRow("SELECT * FROM conf").Scan(&conf.ID, &conf.AccessKey, &conf.SecretKey, &conf.IsReady)

	if err != nil {
		utils.LogUtil(err, 1)
	}
	return conf.IsReady
}

// ReadConf アクセスキー、シークレットキーをdbから読み取る関数
// 戻り値はaccessKeyとシークレットキー
func ReadConf() (AccessKey, SecretKey string) {
	var conf Conf
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.QueryRow("SELECT * FROM conf").Scan(&conf.ID, &conf.AccessKey, &conf.SecretKey, &conf.IsReady)

	if err != nil {
		panic(err.Error())
	}
	return conf.AccessKey, conf.SecretKey
}

// GETConf アクセスキー、シークレットキーをdbから読み取るための関数。
// confに対してGETメソッドを送った場合に利用する
func GETConf(w http.ResponseWriter, r *http.Request) {
	var conf Conf
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	defer db.Close()

	db.QueryRow("SELECT * FROM conf").Scan(&conf.ID, &conf.AccessKey, &conf.SecretKey, &conf.IsReady)

	if err != nil {
		utils.LogUtil(err, 1)
	}
	json.NewEncoder(w).Encode(&conf)
}

//PUTConf アクセスキーを更新するための関数。confに対してPUTメソッドを送った場合に利用する
func PUTConf(w http.ResponseWriter, r *http.Request) {
	var conf Conf
	json.NewDecoder(r.Body).Decode(&conf)
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	defer db.Close()
	fmt.Println(r.Body)

	ins, err := db.Prepare("UPDATE conf SET access_key=?,secret_key=?,is_ready=?")
	if err != nil {
		utils.LogUtil(err, 1)
	}
	ins.Exec(conf.AccessKey, conf.SecretKey, conf.IsReady)
	fmt.Println(conf)
	defer ins.Close()
	fmt.Println(ins)
}

// WriteBuyOrderInfo 買い注文履歴を保存するための関数
// 価格と数量とorder_idを引数とする。
func WriteBuyOrderInfo(price, size float64, childOrderID string) {
	data := new(buyOrder)
	*data = buyOrder{price, size, childOrderID}
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	defer db.Close()
	ins, err := db.Prepare("INSERT INTO buy_order (order_type,price,size,child_order_id,execution_status) VALUES('LIMIT',?,?,?,0)")
	if err != nil {
		utils.LogUtil(err, 1)
	}
	ins.Exec(data.price, data.size, data.childOrderID)
	defer ins.Close()
}

// UpdateTradeMode 取引モードを更新するための関数。
// モード番号を引数とする。返り値なし
func UpdateTradeMode(mode int) {
	m := mode
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	defer db.Close()
	ins, err := db.Prepare("UPDATE mode SET trade_mode=?")
	if err != nil {
		utils.LogUtil(err, 1)
	}
	ins.Exec(m)
	defer ins.Close()
}

// ReadBuyOrderInfo 買い注文履歴を読み取るための関数
// 引数はカラム番号。戻り値は、カラム番号にひもづく値をinterface型を戻り値とする。
func ReadBuyOrderInfo(columnNum int) interface{} {
	cn := columnNum
	var boi buyOrderInfo
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	defer db.Close()

	db.QueryRow("SELECT * FROM buy_order ORDER BY updated_at DESC LIMIT 1").Scan(&boi.id, &boi.orderType, &boi.price, &boi.size, &boi.childOrderID, &boi.executionStatus, &boi.createdAt, &boi.updatedAt)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	switch cn {
	case 3:
		value := boi.price
		return value
	case 5:
		value := boi.childOrderID
		return value
	default:
		return 0
	}
}

// ReadSellOrderInfo 売り注文履歴を読み取るための関数
// 引数はカラム番号。戻り値は、カラム番号にひもづく値をinterface型を戻り値とする。
// HACK: カラム番号で返り値を取得しているが、他にいい方法がないか検討
func ReadSellOrderInfo(columnNum int) interface{} {
	cn := columnNum
	var soi sellOrderInfo
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	defer db.Close()

	db.QueryRow("SELECT * FROM sell_order ORDER BY updated_at DESC LIMIT 1").Scan(&soi.id, &soi.orderType, &soi.price, &soi.size, &soi.childOrderID, &soi.executionStatus, &soi.createdAt, &soi.updatedAt)
	if err != nil {
		utils.LogUtil(err, 1)
	}
	fmt.Println(soi)
	switch cn {
	case 3:
		value := soi.price
		return value
	case 5:
		value := soi.childOrderID
		return value
	default:
		return 0
	}
}

// WriteSellOrderInfo 売り注文情報を書き込むための関数
// 引数はprice、size、order_idとする。
func WriteSellOrderInfo(price, size float64, childOrderID string) {
	data := new(sellOrder)
	*data = sellOrder{price, size, childOrderID}
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	ins, err := db.Prepare("INSERT INTO sell_order (order_type,price,size,child_order_id,execution_status) VALUES('LIMIT',?,?,?,0)")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(data.price, data.size, data.childOrderID)
	defer ins.Close()
}
