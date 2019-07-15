package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

var chat []string

func home(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		t, _ := template.ParseFiles("home.html")
		t.Execute(res, nil)
	}

	if req.Method == "POST" {
		req.ParseForm()
		fmt.Println(req.Form)
	}
}

func refresh(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	mess, _ := json.Marshal(chat)
	res.Write(mess)
	// fmt.Fprintf(res, mess)
}

func main() {
	chat = append(chat, "This is a message")
	chat = append(chat, "This is a another message")

	fmt.Println("Started Server")

	http.HandleFunc("/", home)
	http.HandleFunc("/refresh", refresh)
	http.ListenAndServe(":9000", nil)

}
