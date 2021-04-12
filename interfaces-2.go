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

/* implementation of "Interface" interface in sort package */
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].id < products[j].id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
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
	sort.Sort(products)
	fmt.Printf("After sorting:\n%v", products)
}
