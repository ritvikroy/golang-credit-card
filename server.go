package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	Port = ":8080"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", serveStatic)
	http.ListenAndServe(Port, nil)
}
