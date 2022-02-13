package main

import (
	"net/http"

	"github.com/EmersonTeles/golang_Store/routes"
)

func main(){
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}