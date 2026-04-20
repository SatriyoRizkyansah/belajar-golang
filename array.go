package main

import (
	"fmt"
)

func main() {
	var names [3]string

	names[0] = "Rio"
	names[1] = "Satriyo"
	names[2] = "Pratama"

	fmt.Println("Names: ", names)

	var numbers = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Numbers: ", numbers)

	var fruits = [...]string{"Apple", "Banana", "Cherry"}

	fmt.Println("Fruits: ", fruits)	
}