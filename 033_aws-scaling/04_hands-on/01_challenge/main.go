package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"io"
	"net/http"
	"os"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "vpadmin@vp-mysql:T00goodtobetrue@tcp(vp-mysql.mysql.database.azure.com:3306)/test02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/instance", instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "index page")
	check(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	//query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}

	fmt.Fprintln(w, s)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func instance (w http.ResponseWriter, req *http.Request){
	name, err := os.Hostname()
	check(err)
	_, err = io.WriteString(w, name)
	check(err)
}
