package main

import "fmt"

func main() {
	// сайд эффект
	// признак пустого значения

	num := 3
	square(num)
	fmt.Println(num)

	squarePointer(&num)
	fmt.Println(num)

	// empty value flag

	var wallet1 *int
	fmt.Println(hasWallet(wallet1))

	wallet2 := 0
	fmt.Println(hasWallet(&wallet2))

	wallet3 := 100
	fmt.Println(hasWallet(&wallet3))

}

func square(num int) {
	num *= num
}

func squarePointer(num *int) {
	*num *= *num
}

func hasWallet(money *int) bool {
	return money != nil
}
