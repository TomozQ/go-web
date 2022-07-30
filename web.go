package main

import (
	"log"
	"net/http"
	"text/template"
)

func main () {
	html := `
		<html><body>
		<h1>HELLO</h1>
		<p>This is sample message.</p>
		</body></html>
	`

	// Newでテンプレートを作成、Parseでhtmlの内容をパースしたTemplateを作成
	tf, er := template.New("index").Parse(html)
	if er != nil {
		// Fatal → ログとして値を出力するのに使用される関数エラー時にエラー内容がターミナルに出力される。
		log.Fatal(er)
	}

	hh := func(w http.ResponseWriter, rq *http.Request) {
		// <<template>>.Excuteでテンプレートをレンダリング出力
		er = tf.Execute(w, nil)
		if er != nil {
			log.Fatal(er)
		}
	}

	http.Handle("/hello", http.HandlerFunc(hh))

	http.ListenAndServe("", nil)
}