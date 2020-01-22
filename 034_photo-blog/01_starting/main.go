package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/expire", expire)
	// add route to server pictures
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	//process form submission
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()
		//create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		// create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		//copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)
		c = appendValue(w, c, fname)
	}
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//HttpOnly: true,
			//Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {

	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c.MaxAge = 30
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
