package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v\n", err)
		return
	}

	name := "peterhuang"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server started failed, err:%v\n", err)
		return
	}
}
