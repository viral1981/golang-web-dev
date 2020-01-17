package main

import (
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			//HttpOnly: true,
			//Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}
