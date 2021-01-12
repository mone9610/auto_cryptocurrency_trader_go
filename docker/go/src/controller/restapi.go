package controller

import (
	"fmt"
	"log"
	"model"
	"net/http"

	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// REST API　Serverとしてgoを起動するための関数
func RESTAPI() {
	fmt.Println("restapi start")
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/conf", model.GETConf).Methods("GET")
	//リクエストボディ: {"ID":1,"Access_Key":"hogehoge","Secret_Key":"hugahuga"}
	router.HandleFunc("/api/v1/conf", model.PUTConf).Methods("PUT")
	// ヘルスチェック用のパス
	router.HandleFunc("/health", healthCheck).Methods("GET")

	// CORS対応
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPut,
		},
		AllowedHeaders: []string{"*"},
	})
	handler := c.Handler(router)

	// Start Server
	log.Println("Listen Server ....")
	// 異常があった場合、処理を停止したいため、log.Fatal で囲む
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Health Check OK")
}
