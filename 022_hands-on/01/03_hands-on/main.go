package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", def)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog running")

}

func def(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "default")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("viral.gohtml")
	tpl.ExecuteTemplate(w, "viral.gohtml", "Viral")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
}
