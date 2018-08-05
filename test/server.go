package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 単純な構造体
type Data1 struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {

	// ルーティング
	router := mux.NewRouter()

	// URL別設定
	router.HandleFunc("/sample1", func(w http.ResponseWriter, r *http.Request) {

		// 構造体を定義
		var data1 = Data1{}
		data1.Title = "sample1"
		data1.Message = "hello, sample1"
		data1.Status = 100

		// jsonエンコード
		outputJson, err := json.Marshal(&data1)
		if err != nil {
			panic(err)
		}

		// jsonヘッダーを出力
		w.Header().Set("Content-Type", "application/json")

		// jsonデータを出力
		fmt.Fprint(w, string(outputJson))

		// log
		log.Print(r.URL.Path)
	})

	// ハンドル割当
	http.Handle("/", router)

	// log
	log.Print("localhost:9001")

	// ポート
	http.ListenAndServe(":9001", nil)

}
