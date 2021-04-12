package main

import "fmt"

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

type Products []Product

func (product Product) Print() {
	fmt.Printf("Id = %v, Name = %v, Cost = %v, Units = %v, Category = %v\n", product.id, product.name, product.cost, product.units, product.category)
}

//Receive a pointer if the method is intended to change the state of the receiver

func (product *Product) ApplyDiscount(discount float32) {
	product.cost = product.cost * ((100 - discount) / 100)
}

/* func ApplyDiscount(product *Product, discount float32) {
	product.cost = product.cost * ((100 - discount) / 100)
} */

func main() {
	var pen = Product{100, "Pen", 10, 50, "stationary"}
	//Print(pen)
	pen.Print()

	products := Products{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "utencil"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
	}

	ken := Product{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"}
	fmt.Printf("Index of 'ken' - %v\n", products.Index(ken))
	fmt.Printf("Include 'ken' - %v\n", products.Include(ken))
	fmt.Printf("Any stationary products ? - %v\n", products.Any(func(p Product) bool {
		return p.category == "stationary"
	}))
	fmt.Printf("Are all stationary products ? - %v\n", products.All(func(p Product) bool {
		return p.category == "stationary"
	}))
	fmt.Printf("Filter stationary products  - %v\n", products.Filter(func(p Product) bool {
		return p.category == "stationary"
	}))

	/*  */
	//var penPtr = &Product{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"}
	fmt.Printf("Before discount, product = %v\n", pen)
	pen.ApplyDiscount(10)

	fmt.Printf("After discount, product = %v\n", pen)
}

func (products Products) Index(p Product) int {
	for idx, product := range products {
		if product == p {
			return idx
		}
	}
	return -1
}

func (products Products) Include(p Product) bool {
	return products.Index(p) != -1
}

func (products Products) Any(predicate func(p Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(p Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func (products Products) Filter(predicate func(p Product) bool) Products {
	result := []Product{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}
