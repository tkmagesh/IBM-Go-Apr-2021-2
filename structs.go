package main

import (
	"fmt"
)

//type MyNumber int

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

//composition over inheritance
/*
type PerishableProduct struct {
	product Product
	expiry  string
}
*/

type PerishableProduct struct {
	Product
	expiry string
}

func main() {
	//var pen = Product{100, "Pen", 10, 50, "Stationary"}
	var pen = Product{id: 100, name: "Pen", cost: 10, units: 50, category: "Stationary"}
	fmt.Println(pen)

	//var pencil = new(Product)
	var pencil = Product{}
	fmt.Println(pencil)

	var pointerToPen *Product = &pen
	fmt.Println(pen.name)
	fmt.Println(pointerToPen.name)

	/*
		var grapes = PerishableProduct{product: Product{id: 200, name: "grapes", cost: 70, units: 10, category: "fruits"}, expiry: "2 Days"}
		fmt.Println(grapes)

		fmt.Println(grapes.product.cost, grapes.product.units)
	*/

	var grapes = PerishableProduct{Product: Product{id: 200, name: "grapes", cost: 70, units: 10, category: "fruits"}, Dummy: Dummy{name: "Dummy"}, expiry: "2 Days"}
	fmt.Println(grapes)
	fmt.Println(grapes.Product.cost, grapes.Product.units)
	fmt.Println(grapes.cost, grapes.units)

}
