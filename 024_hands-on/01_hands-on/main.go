package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err !=nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w,"dog.gohtml", nil)
}

func doggy (w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/toby.jpg", doggy)
	http.ListenAndServe(":8080", nil)
}
