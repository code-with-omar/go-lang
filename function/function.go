package main

import "fmt"

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}
func calculateSum(n1 int, n2 int) int {
	return n1 + n2
}
func main() {
	num1 := 5
	num2 := 10
	add(num1, num2)
	sum := calculateSum(num1,num2)
	fmt.Println(sum)
}