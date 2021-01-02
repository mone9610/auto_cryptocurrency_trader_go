package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// REST API　Serverとしてgoを起動するための関数
func RESTAPI() {
	router := mux.NewRouter()

	// エンドポイント
	// router.HandleFunc("/api/v1/conf", model.UpdateConf).Methods("PUT")
	// router.HandleFunc("/articles", getArticles).Methods("GET")

	// Start Server
	log.Println("Listen Server ....")
	// 異常があった場合、処理を停止したいため、log.Fatal で囲む
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get all articles")
}
