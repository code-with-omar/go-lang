package main

import "fmt"

func main() {
	age := 12
	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 18:
		fmt.Println("Teen")
	default:
		fmt.Println("Adult")
	}
}