package main

import "fmt"

func main() {
	age := 19

	if age < 18 {
		println("Age is less than 18")
	}

	if isChild := isChildren(age); isChild == true {
		fmt.Println("You are young")
	} else {
		fmt.Println("Gogogo")
	}
}

func isChildren(age int) bool {
	return age < 18
}
