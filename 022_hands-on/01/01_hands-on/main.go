package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/", def)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)

}

func dog (w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog running")

}

func def(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "default")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello Viral")
}
