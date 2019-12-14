package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Breakfast string
	Lunch     string
	Dinner    string
}

type restaurant struct {
	Name string
	Menu menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	nigerianRestaurant := restaurant{
		Name: "Mr Bigs",
		Menu: menu{
			Breakfast: "Oat and Bread",
			Lunch:     "Fried rice and chicken",
			Dinner:    "Pounded yam and turkey",
		},
	}

	err := tpl.Execute(os.Stdout, nigerianRestaurant)
	if err != nil {
		log.Fatalln(err)
	}
}
