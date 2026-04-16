package main

import "fmt"

func main() {
	name := "Rio"
	age := 20
	isMarried := false

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Is Married: ", isMarried)

	name = "Rio Satriyo"
	age = 21
	isMarried = true

	fmt.Println("Updated Name: ", name)
	fmt.Println("Updated Age: ", age)
	fmt.Println("Updated Is Married: ", isMarried)

	var (
	country = "Indonesia"
	city    = "Jakarta"
	)

	fmt.Println("Country: ", country)
	fmt.Println("City: ", city)
}