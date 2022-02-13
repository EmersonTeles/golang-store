package routes

import (
	"net/http"

	"github.com/EmersonTeles/golang_Store/controllers"
)

func LoadRoutes(){
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/produtos/novo", controllers.NewProduct)
	http.HandleFunc("/produtos/editar", controllers.EditProduct)
	http.HandleFunc("/produtos/update", controllers.Update)
	http.HandleFunc("/produtos/insert", controllers.Insert)
}