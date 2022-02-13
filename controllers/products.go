package controllers

import (
	"html/template"
	"net/http"

	"github.com/EmersonTeles/golang_Store/models"
)


var temp = template.Must(template.ParseGlob ("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request){
	products := models.SearchProducts()
	temp.ExecuteTemplate(w, "Index", products)
}
func NewProduct(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "New", nil)
}