package main

import (
	"html/template"
	"log"
	"os"
)

type menuItem struct {
	Name  string
	Price float64
}

type menu struct {
	Name  string
	Items []menuItem
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menulist := []menu{
		menu{
			Name: "Breakfast",
			Items: []menuItem{
				menuItem{
					Name:  "Bacon and Eggs",
					Price: 50.00,
				},
				menuItem{
					Name:  "Pancakes",
					Price: 30.00,
				},
			},
		},
		menu{
			Name: "Lunch",
			Items: []menuItem{
				menuItem{
					Name:  "Hamburger",
					Price: 75.00,
				},
				menuItem{
					Name:  "Sausages",
					Price: 50.00,
				},
			},
		},
		menu{
			Name: "Dinner",
			Items: []menuItem{
				menuItem{
					Name:  "Lasagna",
					Price: 100.00,
				},
				menuItem{
					Name:  "T-Bone Steak",
					Price: 250.00,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, menulist)
	if err != nil {
		log.Fatalln(err)
	}
}
