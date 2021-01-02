package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbInfo = "root:root@tcp(mysql:3306)/auto_trader"

// DB書き込み用のtickerデータの構造体
type TickerChild struct {
	High float64
	Low  float64
	Last float64
}

// トレンド分析用に利用するtickerデータの構造体
type TickerInfo struct {
	Id         int
	High       float64
	Last       float64
	Low        float64
	Created_at string
}

// tickerHistoryの構造体を格納するための配列
type TickerInfoArray []TickerInfo

// 買い注文の情報を格納するための構造体
type BuyOrder struct {
	price          float64
	size           float64
	child_order_id string
}

// 買い注文の情報を取得するための構造体
type BuyOrderInfo struct {
	id               int
	order_type       string
	price            float64
	size             float64
	child_order_id   string
	execution_status int
	created_at       string
	updated_at       string
}

// 売り注文の情報を格納するための構造体
type SellOrder struct {
	price          float64
	size           float64
	child_order_id string
}

type Mode struct {
	id         int
	trade_mode int
	created_at string
	updated_at string
}

type Conf struct {
	Id         int
	Access_Key string
	Secret_Key string
}

// 最高価格、最低価格、最終価格をDBに書き込む関数
// 引数は最高価格、最低価格、最終価格とし、戻り値はboolとする
func WriteTickerInfo(high, low, last float64) bool {
	var data = new(TickerChild)
	*data = TickerChild{high, low, last}
	fmt.Println(data)

	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err.Error())
		return false
	}
	defer db.Close()

	ins, err := db.Prepare("INSERT INTO ticker_info (high,last,low) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
		return false
	}
	ins.Exec(data.High, data.Last, data.Low)
	defer ins.Close()
	fmt.Println(ins)
	return true
}

// 自動取引jobにて、モードを確認するための関数
// 引数なし、戻り値はモードとする
func CheckTradeMode() int {
	var m Mode
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// ToDo:全件検索は効率が悪いのであとで直す
	db.QueryRow("SELECT * FROM mode").Scan(&m.id, &m.trade_mode, &m.created_at, &m.updated_at)

	if err != nil {
		panic(err.Error())
	}
	return m.trade_mode
}

// tikcer_infoテーブルから、最新の最高取引価格を取り出す関数
// 引数なし、戻り値は最新の取引価格(float64)
func ReadTickerInfo() float64 {
	var ti TickerInfo
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// ToDo:全件検索は効率が悪いのであとで直す
	db.QueryRow("SELECT * FROM ticker_info ORDER BY created_at DESC LIMIT 1").Scan(&ti.Id, &ti.High, &ti.Last, &ti.Low, &ti.Created_at)

	if err != nil {
		panic(err.Error())
	}
	return ti.High
}

// アクセスキー、シークレットキーをdbから読み取る
func ReadConf() (accessKey, secretKey string) {
	var conf Conf
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.QueryRow("SELECT * FROM conf").Scan(&conf.Id, &conf.Access_Key, &conf.Secret_Key)

	if err != nil {
		panic(err.Error())
	}
	return conf.Access_Key, conf.Secret_Key
}

// アクセスキー、シークレットキーをdbを更新する
// func UpdateConf() (w http.ResponseWriter, r *http.Request) {
// 	var conf Conf
// 	json.NewDecoder(r.Body).Decode(&conf)
// 	db, err := sql.Open("mysql", dbInfo)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	db.QueryRow("SELECT * FROM conf").Scan(&conf.Id, &conf.Access_Key, &conf.Secret_Key)

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	json.NewEncoder(w).Encode(article)
// 	return conf.Access_Key, conf.Secret_Key
// }

// 買い注文履歴を保存するための関数
// 価格と数量とorder_idを引数とする。
func WriteBuyOrderInfo(price, size float64, childOrderId string) {
	data := new(BuyOrder)
	*data = BuyOrder{price, size, childOrderId}
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	ins, err := db.Prepare("INSERT INTO buy_order (order_type,price,size,child_order_id,execution_status) VALUES('LIMIT',?,?,?,0)")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(data.price, data.size, data.child_order_id)
	defer ins.Close()
	fmt.Println(ins)
}

// modeを更新するための関数。
// モード番号を引数とする。返り値なし
func UpdateTradeMode(mode int) {
	m := mode
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	ins, err := db.Prepare("UPDATE mode SET trade_mode=?")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(m)
	defer ins.Close()
	fmt.Println(ins)
}

// 買い注文履歴を読み取るための関数
// 引数はカラム番号。戻り値は、カラム番号にひもづく値をinterface型を戻り値とする。
// HACK: カラム番号で返り値を取得しているが、他にいい方法がないか検討
func ReadBuyOrderInfo(columnNum int) interface{} {
	cn := columnNum
	fmt.Println(cn)
	var boi BuyOrderInfo
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.QueryRow("SELECT * FROM buy_order ORDER BY updated_at DESC LIMIT 1").Scan(&boi.id, &boi.order_type, &boi.price, &boi.size, &boi.child_order_id, &boi.execution_status, &boi.created_at, &boi.updated_at)
	if err != nil {
		panic(err.Error())
	}
	switch cn {
	case 3:
		value := boi.price
		return value
	case 5:
		value := boi.child_order_id
		return value
	default:
		return 0
	}
}

// 売り注文履歴を読み取るための関数
// 引数はカラム番号。戻り値は、カラム番号にひもづく値をinterface型を戻り値とする。
// HACK: カラム番号で返り値を取得しているが、他にいい方法がないか検討
func ReadSellOrderInfo(columnNum int) interface{} {
	cn := columnNum
	fmt.Println(cn)
	var boi BuyOrderInfo
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.QueryRow("SELECT * FROM sell_order ORDER BY updated_at DESC LIMIT 1").Scan(&boi.id, &boi.order_type, &boi.price, &boi.size, &boi.child_order_id, &boi.execution_status, &boi.created_at, &boi.updated_at)
	if err != nil {
		panic(err.Error())
	}
	switch cn {
	case 3:
		value := boi.price
		return value
	case 5:
		value := boi.child_order_id
		return value
	default:
		return 0
	}
}

// 売り注文情報を書き込むための関数
// 引数はprice、size、order_idとする。
func WriteSellOrderInfo(price, size float64, childOrderID string) {
	data := new(SellOrder)
	*data = SellOrder{price, size, childOrderID}
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	ins, err := db.Prepare("INSERT INTO sell_order (order_type,price,size,child_order_id,execution_status) VALUES('LIMIT',?,?,?,0)")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(data.price, data.size, data.child_order_id)
	defer ins.Close()
	fmt.Println(ins)
}
