package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

/*
 html/template の文法は text/template と一緒なので
 https://golang.org/pkg/text/template/
 を参照する
*/

// テンプレートに単一変数を渡す
func simpleHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/sample.html"))

	message := "Hello, World"

	err := t.ExecuteTemplate(w, "sample.html", message)
	if err != nil {
		log.Fatalf("can't execute template: %v", err)
	}
}

// テンプレートにマップを渡す
func mappedDataHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/map.html"))

	m := map[string]string{
		"hoge": "foo",
		"fuga": "bar",
	}

	err := t.ExecuteTemplate(w, "map.html", m)
	if err != nil {
		log.Fatalf("can't execute template: %v", err)
	}
}

// テンプレートに構造体を渡す
func structDataHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/struct.html"))

	s := struct {
		Name string
		Age  int
		Job  string
	}{"Mr. Anderson", 28, "Messiah"}

	err := t.ExecuteTemplate(w, "struct.html", s)
	if err != nil {
		log.Fatalf("can't execute template: %v", err)
	}
}

// テンプレート中に使用する独自関数を定義できる
func funcMapHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"htmlunescape": func(text string) template.HTML { return template.HTML(text) },
		"now":          func() int64 { return time.Now().Unix() },
	}

	t := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/funcMap.html"))

	m := map[string]string{
		"hoge": "<a href='http://google.com'>google.com</a>",
		"fuga": "bar",
	}

	err := t.ExecuteTemplate(w, "funcMap.html", m)
	if err != nil {
		log.Fatalf("can't execute template: %v", err)
	}
}

// 名前付きテンプレート
func namedTemplateHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/named.html"))

	m := map[string]string{
		"hoge": "foo",
		"fuga": "bar",
	}

	err := t.ExecuteTemplate(w, "Morpheus", m) // execute by defined name which define in named.html
	if err != nil {
		log.Fatalf("can't execute template: %v", err)
	}
}

// テンプレートの挿入
func includeTemplateHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/main.html", "templates/header.html"))

	m := struct {
		Title   string
		Content string
	}{"This is title.", "oooooo xooxooo ooo xo"}

	// TODO: include 先の header.html の {{ .Title }} に変数を送る方法はあるのか？
	err := t.ExecuteTemplate(w, "main.html", m)
	if err != nil {
		log.Fatalf("can't execute template: %v", err)
	}
}

func main() {
	http.HandleFunc("/", simpleHandler)
	http.HandleFunc("/map", mappedDataHandler)
	http.HandleFunc("/struct", structDataHandler)
	http.HandleFunc("/funcMap", funcMapHandler)
	http.HandleFunc("/namedTemplate", namedTemplateHandler)
	http.HandleFunc("/includeTemplate", includeTemplateHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("can't listen or serve http: %v", err)
	}
}
