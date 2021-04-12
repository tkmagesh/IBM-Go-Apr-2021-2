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

	var grapes = PerishableProduct{Product: Product{id: 200, name: "grapes", cost: 70, units: 10, category: "fruits"}, expiry: "2 Days"}
	fmt.Println(grapes)
	fmt.Println(grapes.Product.cost, grapes.Product.units)
	fmt.Println(grapes.cost, grapes.units)

	products := []Product{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "utencil"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
	}

	ken := Product{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"}
	fmt.Printf("Index of 'ken' - %v\n", Index(products, ken))
	fmt.Printf("Include 'ken' - %v\n", Include(products, ken))
	fmt.Printf("Any stationary products ? - %v\n", Any(products, func(p Product) bool {
		return p.category == "stationary"
	}))
	fmt.Printf("Are all stationary products ? - %v\n", All(products, func(p Product) bool {
		return p.category == "stationary"
	}))
	fmt.Printf("Filter stationary products  - %v\n", Filter(products, func(p Product) bool {
		return p.category == "stationary"
	}))
}

func Index(products []Product, p Product) int {
	for idx, product := range products {
		if product == p {
			return idx
		}
	}
	return -1
}

func Include(products []Product, p Product) bool {
	return Index(products, p) != -1
}

func Any(products []Product, predicate func(p Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func All(products []Product, predicate func(p Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func Filter(products []Product, predicate func(p Product) bool) []Product {
	result := []Product{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

/*
Index() => index of a product in the products slice
Include() => if the given product in the given slice
Any() => if there is any product matching the criteria (user defined)
All() => if all the products matches the given criteria
Filter() => return the products matching the given criteria


*/
