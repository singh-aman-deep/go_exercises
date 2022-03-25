package main

import "fmt"

func main() {
	fmt.Println("total of 3 and 5 is", sum(3, 5))
	fmt.Println("Multiplication of 3 and 5 is", mul(3, 5))
}

func sum(param ...int) int {
	total := 0
	for _, v := range param {
		total = total + v
	}
	return total
}

func mul(param ...int) int {
	total := 1
	for _, v := range param {
		total = total * v
	}
	return total
}
