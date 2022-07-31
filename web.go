package main

import (
	"log"
	"net/http"
	"text/template"
	"github.com/gorilla/sessions"
)

// Temp is template structure
type Temps struct {
	notemp * template.Template
	indx * template.Template
	helo * template.Template
}

// CookieStoreを作成
var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-1234"))	// ([]byte("secret-key-12345")) → セッションの秘密キー

// Template for no-template
func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

// get target Template
func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("templates/" + fname + ".html", "templates/head.html", "templates/foot.html") // メインで使うテンプレートを第一引数に指定する
	return tmps
}

// index handler.
func index(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		Template  string
		Title 		string
		Message 	string
	}{
		Template: "index",
		Title: "Index",
		Message: "This is Top page.",
	}

	er := page("index").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// hello handler.
func hello(w http.ResponseWriter, rq *http.Request) {
	data := []string{
		"One", "Two", "Three",
	}

	item := struct {
		Title string
		Data []string
	}{
		Title: "Hello",
		Data: data,
	}

	er := page("hello").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}


func main () {

	// index handling
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request){
		index(w, rq)
	})

	// helo handling
	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request){
		hello(w, rq)
	})

	http.ListenAndServe("", nil)
}