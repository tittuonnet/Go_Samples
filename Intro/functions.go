package main

import "fmt"

/* Sum of two integer values) */
func sum(a, b int) int {
	return a + b
}

/* Multi return function */
func multiReturn(a, b int) (int, int) {
	return (a + b), (a * b)
}

func main() {
	fmt.Println("Sum of two integer values")
	fmt.Println(sum(5, 10))
	fmt.Println("Multi return function sample")

	sum, multiplication := multiReturn(5, 10)
	fmt.Println("Sum is ", sum)
	fmt.Println("Multiplication of ", multiplication)
}
