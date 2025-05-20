package main

import "fmt"
func welcome (){
	fmt.Println("Welcome to the application")
}
func userName() string{
	fmt.Println("Please enter your name: ")
	var name string
	fmt.Scanln(&name)
	return name
}
func getNumbers()(int,int){
var num1, num2 int
	fmt.Println("Enter first number")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number")
	fmt.Scanln(&num2)
	return num1,num2
}
func calculationSum(num1 int, num2 int) int{
	sum := num1+num2
	return sum
}

func displayResult(name string,sum int){
	fmt.Println("Hello", name)
	fmt.Println("Summation =", sum)
}
func goodbyeMessage(){
	fmt.Println("Thank you for using the application")
}

func main() {
	welcome()
	name := userName()
	num1, num2 := getNumbers()
	sum := calculationSum(num1,num2)
	displayResult(name,sum)
	goodbyeMessage()
}