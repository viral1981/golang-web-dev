package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name        string
	Description string
	Price       float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		meal{
			Meal: "Breakfast",
			Item: []item{
				item{
					Name:        "Weetabix",
					Description: "Cereal",
					Price:       5.23,
				},
				item{
					Name:        "Bacon Butty",
					Description: "Fry Up",
					Price:       3.00,
				},
				item{
					Name:        "Milk",
					Description: "Beverage",
					Price:       1.00,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Item: []item{
				item{
					Name:        "Hamburger",
					Description: "Burger",
					Price:       10.25,
				},
				item{
					Name:        "Chesse Toasted Sandwich",
					Description: "Sandwich",
					Price:       5.69,
				},
				item{
					Name:        "French Fries",
					Description: "Chips",
					Price:       0.99,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
