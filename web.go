package main

import (
	"log"
	"net/http"
	"text/template"
)

func main () {
	tf, er := template.ParseFiles("templates/hello.html")
	// errorが発生して読み込めなかった場合には、stringリテラルを使ってテンプレートを生成する。
	if er != nil {
		tf, _ = template.New("index").Parse("<html><body><h1>NO TEMPLATE.</h1></body></html>")
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