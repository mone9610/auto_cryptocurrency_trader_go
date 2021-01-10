package controller

import (
	"fmt"
	"log"
	"model"
	"net/http"

	"github.com/gorilla/mux"
)

// REST API　Serverとしてgoを起動するための関数
func RESTAPI() {
	fmt.Println("restapi start")
	router := mux.NewRouter()

	// エンドポイント
	// 以下の形式でリクエストを受け取る。
	// work: {"ID":1,"Access_Key":"hogehoge","Secret_Key":"hugahuga"}
	router.HandleFunc("/api/v1/conf", model.GETConf).Methods("GET")
	router.HandleFunc("/api/v1/conf", model.PUTConf).Methods("PUT")

	// Start Server
	log.Println("Listen Server ....")
	// 異常があった場合、処理を停止したいため、log.Fatal で囲む
	log.Fatal(http.ListenAndServe(":8000", router))
}
