package main

import (
	"curso_alura_golang/Go-crie-uma-aplicacao-web/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
