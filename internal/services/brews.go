package services

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

type Brew struct {
	BrewMethod string
	BrewTime   string
	CoffeeIn   float32
	CoffeeOut  float32
}

func IndexHandler(w http.ResponseWriter, r *http.Request, htmlTemplates []string) {
	brews := map[string][]Brew{
		"Brews": {
			{BrewMethod: "Drip", BrewTime: "05:00", CoffeeIn: 15.6, CoffeeOut: 250},
			{BrewMethod: "Espresso", BrewTime: "00:35", CoffeeIn: 15.6, CoffeeOut: 30},
			{BrewMethod: "Espresso", BrewTime: "00:40", CoffeeIn: 15.6, CoffeeOut: 30},
		},
	}

	htmlTemplates = append(htmlTemplates, "web/brew.html")

	tmpl := template.Must(template.ParseFiles(htmlTemplates...))
	tmpl.Execute(w, brews)
}

func GetBrews(w http.ResponseWriter, r *http.Request) {
	brews := map[string][]Brew{
		"Brews": {
			{BrewMethod: "Drip", BrewTime: "05:00", CoffeeIn: 15.6, CoffeeOut: 250},
			{BrewMethod: "Espresso", BrewTime: "00:35", CoffeeIn: 15.6, CoffeeOut: 30},
			{BrewMethod: "Espresso", BrewTime: "00:40", CoffeeIn: 15.6, CoffeeOut: 30},
		},
	}

	tmpl := template.Must(template.ParseFiles("web/brew.html"))
	tmpl.Execute(w, brews)
}

func NewBrew(w http.ResponseWriter, r *http.Request) {
	brewMethod := r.PostFormValue("brewMethod")
	brewTime := r.PostFormValue("brewTime")
	coffeeIn, _ := strconv.ParseFloat(r.PostFormValue("coffeeIn"), 32)
	coffeeOut, _ := strconv.ParseFloat(r.PostFormValue("coffeeOut"), 32)

	fmt.Println(brewMethod)

	tmpl := template.Must(template.ParseFiles("web/brew.html"))
	tmpl.ExecuteTemplate(w, "singleBrew", Brew{BrewMethod: brewMethod, BrewTime: brewTime, CoffeeIn: float32(coffeeIn), CoffeeOut: float32(coffeeOut)})
}
