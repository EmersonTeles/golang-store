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

	selectProducts, err := db.Query("select * from products order by id asc")	
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	products := []Product{}
	for selectProducts.Next(){
		var id, quantity int
		var name,description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)
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
func SearchProductById(id string) Product{
	db := db.ConnectToPostgres()

	selectProduct, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	for selectProduct.Next(){
		var id, quantity int
		var name,description string
		var price float64
		err = selectProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
	}
	return p
}
func InsertProduct(name string, description string , price float64 , quantity int) {
	db := db.ConnectToPostgres()

	insertProducts, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic("deu erro no models/InsertProduct l47: " + err.Error())
	}
	insertProducts.Exec(name, description, price, quantity)
	defer db.Close()
}	
func DeleteProduct(id string) {
	db := db.ConnectToPostgres()

	deleteProducts, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteProducts.Exec(id)
	defer db.Close()
}
func UpdateProduct(id string, name string, description string, price float64, quantity int) {
	db := db.ConnectToPostgres()

	updateProducts, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateProducts.Exec(name, description, price, quantity, id)
	defer db.Close()
}