package main

import (
	"net/http"
)

func main () {
	// 実行する処理をまとめた関数をhhに用意
	hh := func(w http.ResponseWriter, rq *http.Request){
		// ResponseWriterの「Write」メソッドを使って、レスポンスにテキストを出力している
		w.Write([]byte("Hello, This is GO-server!!"))	// 引数には出力する値をbyte配列として用意 ← クライアントに送られて表示される。
	}
	// 作成した関数は、HandleFuncを使い、/helloアドレスに割り付けれられる。
	// '/hello' でhh関数を実行する
	http.Handle("/hello", http.HandlerFunc(hh))

	// サーバー実行
	http.ListenAndServe("", nil)
}