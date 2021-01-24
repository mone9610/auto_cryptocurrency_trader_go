package controller

import (
	"log"
	"model"
	"net/http"
	"utils"

	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// RESTAPI RESTAPI Serverとしてgoを起動するための関数
func RESTAPI() {
	utils.LogUtil("RESTAPI start", 0)
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

	// 異常があった場合、処理を停止したいため、log.Fatal で囲む
	// log.Fatal(http.ListenAndServe(":8000", handler))
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	utils.LogUtil("HealthCheck ok", 0)
}
