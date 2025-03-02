package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./t.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	msg := "这是index页面"
	t.ExecuteTemplate(w, "./index.tmpl", msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	msg := "这是home页面"
	t.Execute(w, msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server started failed, err:%v\n", err)
		return
	}
}
