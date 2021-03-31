package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func main() {
	a := 10
	b := 50

	resultado := soma(a, b)

	fmt.Printf("O resultado da soma de %d + %d é: %d\n", a, b, resultado)
}
