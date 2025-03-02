package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	u1 := User{
		Name: "黄熠", Gender: "man", Age: 26,
	}
	m1 := map[string]interface{}{
		"Name": "peterhuang", "Gender": "man", "Age": 20,
	}
	hobbyList := []string{"basketball", "football", "double color ball"}
	t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server started failed, err:%v\n", err)
		return
	}
}
