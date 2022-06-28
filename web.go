package main

import (
	"net/http"
)

func main () {
	// http.FileServer( <<FileSystem>> )
	http.ListenAndServe("", http.FileServer(http.Dir(".")))	// httpパッケージのDir関数
}