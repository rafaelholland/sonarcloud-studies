package main

import "fmt"

func main() {
	fmt.Println(sum(2, 2))
}

func sum(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func multiplytwice(a int, b int) int {
	return a + b*2
}
