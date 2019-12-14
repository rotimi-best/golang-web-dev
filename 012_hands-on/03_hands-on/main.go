package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type region struct {
	Region string
	Hotels []hotel
}

type regions []region

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotelRegions := regions{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Loxy Co",
					Address: "Saint Luis street",
					City:    "Los Angelis",
					Zip:     "109",
					Region:  "southern",
				},
				hotel{
					Name:    "Sunrise master",
					Address: "Blink avenue street",
					City:    "London",
					Zip:     "609",
					Region:  "southern",
				},
			},
		},
		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "Too shady",
					Address: "Collardo main road, 12 street",
					City:    "Manhattan",
					Zip:     "521",
					Region:  "nothern",
				},
				hotel{
					Name:    "Just like home",
					Address: "44 Kevin Hart road",
					City:    "Florida",
					Zip:     "444",
					Region:  "nothern",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotelRegions)
	if err != nil {
		log.Fatalln(err)
	}
}
