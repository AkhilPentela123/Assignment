package main

import "fmt"

func Add(a *int, b *int) int {
	return *a + *b
}

func Sub(a *int, b *int) int {
	return *a - *b
}

func Mul(a *int, b *int) int {
	return *a * *b
}

func Div(a *int, b *int) (int, int) {
	return *a / *b, *a % *b
}

func main() {

	var a, b int

	fmt.Println("enter first number : ")
	fmt.Scanln(&a)
	fmt.Println("enter second number : ")
	fmt.Scanln(&b)

	fmt.Println("Addition:", Add(&a, &b))
	fmt.Println("Subtraction:", Sub(&a, &b))
	fmt.Println("Multiplication", Mul(&a, &b))
	fmt.Print("divide & modular:")
	fmt.Println(Div(&a, &b))

}
