package main

import "fmt"

// Product يمثل أبسط وحدة بيانات في المتجر
type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	p := Product{ID: 1, Name: "Laptop", Price: 1500.0}
	fmt.Printf("تم إنشاء منتج: %+v\n", p)
}
