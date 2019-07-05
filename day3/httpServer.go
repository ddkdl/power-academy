package main

import (
	"net/http"
	"text/template"
)

func home(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9000", nil)
}
