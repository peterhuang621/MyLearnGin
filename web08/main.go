package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Print("parse template failed, err:", err)
		return
	}
	name := "peterhuang"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Print("execute template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server started failed, err:%v\n", err)
		return
	}
}
