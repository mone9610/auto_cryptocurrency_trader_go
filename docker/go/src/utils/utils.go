package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

// 独自headerとbodyでPOSTリクエストを送るための関数
// httpパッケージで対応できないリクエストはこの関数を使う
func NewRequest(method, url string, header map[string]string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
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

// Map型をstring型へ変換する関数
func MapToString(bytes map[string]interface{}) string {
	b, err := json.Marshal(bytes)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return "error"
	}
	string := string(b)
	return string
}

// 小数点切り捨て用関数
// 数値と切捨て位置を引数とし、切捨てた値を返り値とする
func RoundDown(num, places float64) float64 {
	shift := math.Pow(10, places)
	return math.Trunc(num*shift) / shift
}
