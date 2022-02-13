package controllers

import (
	"html/template"
	"net/http"
	"strconv"

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
func EditProduct(w http.ResponseWriter, r *http.Request){
	productId := r.URL.Query().Get("id")
	product := models.SearchProductById(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}
func Insert(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvertToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}

		quantityConvertToInt, err := strconv.Atoi(quantity)
		if err != nil {
			panic(err.Error())
		}
		models.InsertProduct(name, description, priceConvertToFloat, quantityConvertToInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
func Delete(w http.ResponseWriter, r *http.Request){
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
func Update(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvertToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}

		quantityConvertToInt, err := strconv.Atoi(quantity)
		if err != nil {
			panic(err.Error())
		}
		models.UpdateProduct(id, name, description, priceConvertToFloat, quantityConvertToInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}