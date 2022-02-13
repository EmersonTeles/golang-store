package models

import "github.com/EmersonTeles/golang_Store/db"

type Product struct {
	Id int
	Name string
	Description string
	Price float64
	Quantity int
}

func SearchProducts() []Product{
	db := db.ConnectToPostgres()

	selectProducts, err := db.Query("select * from products")	
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	products := []Product{}
	for selectProducts.Next(){
		var id, quantity int
		var name,description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &quantity, &price)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		products = append(products, p)
	}
	defer db.Close()
	return products
}