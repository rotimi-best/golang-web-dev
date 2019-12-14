package main

import (
	"log"
	"os"
	"text/template"
)

type details struct {
	Name, Ingredients string
	Price             float64
}

type meal struct {
	Name    string
	Details details
}

type menu []meal

type restaurant struct {
	Name, WorkingHour, AllowedCurrency string
	Menu                               menu
	NoOfEmployees                      int64
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	myRestaurants := restaurants{
		restaurant{
			Name:            "ABC foods",
			WorkingHour:     "10am - 6pm",
			AllowedCurrency: "UAH",
			Menu: menu{
				meal{
					Name: "BreakFast",
					Details: details{
						Name:        "Chicken soup",
						Ingredients: "Chicken, Spices, Magi cube",
						Price:       100.95,
					},
				},
				meal{
					Name: "Lunch",
					Details: details{
						Name:        "Ewa goyin and eggs",
						Ingredients: "Beans, eggs, spices, tomatoes",
						Price:       80.91,
					},
				},
				meal{
					Name: "Dinner",
					Details: details{
						Name:        "Semolina and okro soup",
						Ingredients: "Semolina, Okro, onions, magi cube",
						Price:       150.0,
					},
				},
			},
			NoOfEmployees: 50,
		},
		restaurant{
			Name:            "XYZ Delight",
			WorkingHour:     "6am - 11:59pm",
			AllowedCurrency: "USD",
			Menu: menu{
				meal{
					Name: "BreakFast",
					Details: details{
						Name:        "Something light",
						Ingredients: "Light ingredients",
						Price:       20.99,
					},
				},
				meal{
					Name: "Lunch",
					Details: details{
						Name:        "Something average",
						Ingredients: "Average ingredients",
						Price:       30.11,
					},
				},
				meal{
					Name: "Dinner",
					Details: details{
						Name:        "Strong meal",
						Ingredients: "Strong ingredients",
						Price:       50.0,
					},
				},
			},
			NoOfEmployees: 100,
		},
	}

	err := tpl.Execute(os.Stdout, myRestaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
