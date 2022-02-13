package main

import (
	"html/template"
	"net/http"

	"github.com/EmersonTeles/golang_Store/models"
)

var temp = template.Must(template.ParseGlob ("templates/*.html"))

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
func index(w http.ResponseWriter, r *http.Request){
	products := models.SearchProducts()
	temp.ExecuteTemplate(w, "Index", products)
}