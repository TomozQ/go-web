package main

import (
	"net/http"
)

func main () {
	// http.ListenServe(アドレス, <<Handler>>)
	http.ListenAndServe("", http.NotFoundHandler())	// 空文字指定でlocalhost NotFoundHandler -> 404 page not found
}