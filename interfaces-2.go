package main

import (
	"fmt"
	"sort"
)

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

func (product Product) String() string {
	return fmt.Sprintf("Id = %v, Name = %v, Cost = %v, Units = %v, Category = %v\n", product.id, product.name, product.cost, product.units, product.category)
}

type Products []Product

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%v", product)
	}
	return result
}

type ProductComparer func(p1 Product, p2 Product) bool

type ProductComparers map[string]ProductComparer

//generalize the below implementation using reflection
var productComparers ProductComparers = ProductComparers{
	"id": func(p1 Product, p2 Product) bool {
		return p1.id < p2.id
	},
	"name": func(p1 Product, p2 Product) bool {
		return p1.name < p2.name
	},
	"cost": func(p1 Product, p2 Product) bool {
		return p1.cost < p2.cost
	},
	"units": func(p1 Product, p2 Product) bool {
		return p1.units < p2.units
	},
	"category": func(p1 Product, p2 Product) bool {
		return p1.category < p2.category
	},
}

var currentProductComparer ProductComparer = productComparers["id"]

/* implementation of "Interface" interface in sort package */
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return currentProductComparer(products[i], products[j])
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Sort(attrName string) {
	currentProductComparer = productComparers[attrName]
	sort.Sort(products)
}

func main() {
	products := Products{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "utencil"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
	}
	fmt.Printf("Products\n%v", products)
	products.Sort("id")
	fmt.Printf("After sorting by id:\n%v", products)
	products.Sort("name")
	fmt.Printf("After sorting by name:\n%v", products)
	products.Sort("cost")
	fmt.Printf("After sorting by cost:\n%v", products)
	products.Sort("units")
	fmt.Printf("After sorting by units:\n%v", products)
	products.Sort("category")
	fmt.Printf("After sorting by category:\n%v", products)
}
