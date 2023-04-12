package main

import (
	"assignment/app"
	"fmt"
)

func main() {

	var a, b int
	fmt.Println("Enter a Value: ")
	fmt.Scanln(&a)
	fmt.Println("Enter b Value: ")
	fmt.Scanln(&b)

	fmt.Printf("Sum of %v and %v is : %v\n", a, b, app.Sum(&a, &b))
	fmt.Printf("Sub of %v and %v is : %v\n", a, b, app.Sub(&a, &b))
	fmt.Printf("Mul of %v and %v is : %v\n", a, b, app.Mult(&a, &b))

	quotient, remainder := app.Divi(&a, &b)

	fmt.Printf("Quotient of %v divided by %v is : %v\n", a, b, quotient)
	fmt.Printf("Remainder of %v divided by %v is : %v\n", a, b, remainder)

}
