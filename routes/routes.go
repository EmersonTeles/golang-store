package routes

import (
	"net/http"

	"github.com/EmersonTeles/golang_Store/controllers"
)

func LoadRoutes(){
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/produtos/novo", controllers.NewProduct)
}