package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Verde musgo", 2999, 20},
		{"Bermuda", "Grená", 2599, 10},
		{"Meias", "Vermelha tomate", 2000, 5},
		{"Tênis", "Nike", 5999, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
