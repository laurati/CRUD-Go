package main

import (
	"laurati/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

//go mod init na pasta laurati
//go mod tidy
