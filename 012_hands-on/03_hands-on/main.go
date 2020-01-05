package main

import(
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name string
	Address string
	City string
	Zip string
	Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "Hotel California",
			Address: "42 Sunst Boulevard",
			City:    "Los Angeles",
			Zip:     "956212",
			Region:  "southern",
		},
		hotel{
			Name:    "HQ",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "souther",
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)

	}
}