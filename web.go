package main

import (
	"net/http"
)

func main () {
	// `` →  改行を含むテキストリテラルを記述するのに使われる。
	// HTMLソースコードのstringをWriteで書き出せば、それがそのままwebブラウザでHTMLとして認識され、表示される。
	msg := `
			<html><body>
			<h1>Hello</h1>
			<p>This is GO-server!!</p>
			</body></html>
	`

	hh := func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte(msg))
	}

	http.Handle("/hello", http.HandlerFunc(hh))

	http.ListenAndServe("", nil)
}