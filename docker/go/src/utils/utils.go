package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const logPath = "./error.log"

//NewRequest 独自headerとbodyでPOSTリクエストを送るための関数
// httpパッケージで対応できないリクエストはこの関数を使う
func NewRequest(method, url string, header map[string]string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		LogUtil(err, 1)
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	return req, nil
}

// SHA256でデジタル署名を作成するための関数
func MakeHMAC(text, secretKey string) string {
	key := secretKey
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(text))
	return hex.EncodeToString(mac.Sum(nil))
}

//MapToString Map型をstring型へ変換する関数
func MapToString(bytes map[string]interface{}) string {
	b, err := json.Marshal(bytes)
	if err != nil {
		LogUtil(err, 1)
		return ""
	}
	string := string(b)
	return string
}

// RoundDown 小数点切り捨て用関数
// 数値と切捨て位置を引数とし、切捨てた値を返り値とする
func RoundDown(num, places float64) float64 {
	shift := math.Pow(10, places)
	return math.Trunc(num*shift) / shift
}

// LogUtil ログ出力用関数
// メッセージを第一引数とし、ログレベルを第二引数とする。
// 第二引数が2の場合はFatalログとして、コンソールとファイルにメッセージを出力し、プロセスを終了する。
// 1の場合はエラーログとして、コンソールとファイルにメッセージを出力する。
// 0の場合はコンソールにのみ出力する。
func LogUtil(message interface{}, level int) {
	if level == 2 {
		logfile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic("cannnot open error.log:" + err.Error())
		}
		defer logfile.Close()

		log.SetOutput(io.MultiWriter(logfile, os.Stdout))

		log.SetFlags(log.Ldate | log.Ltime)
		log.Fatal("FATAL:", message)
	} else if level == 1 {
		logfile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic("cannnot open error.log:" + err.Error())
		}
		defer logfile.Close()
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))

		log.SetFlags(log.Ldate | log.Ltime)
		log.Println("ERROR:", message)
	} else {
		log.Println("INFO:", message)
	}

}
