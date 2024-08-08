package main

import "fmt"

/*
func *Название функции* (*входящие параметры*) (*возвращаемые значения*) {
	*Тело функции*

	return *возвращаемые значения*
}
*/

// main - основная функция
func main() {
	first, second := 1, 2

	Hello()
	HelloToOne("Egor")
	FullNameHello("Egor", "Onekeev")

	summa := Sum(first, second)
	fmt.Println(summa)

	summa, multiply := SumAndMultiply(first, second)
	fmt.Println(summa, multiply)

	_, multiply64 := namedSumAndMultiply(first, second)
	fmt.Println(multiply64)

	var multiplier func(x, y int) int

	multiplier = func(x, y int) int { return x * y }
	fmt.Println(multiplier(first, second))

	divider := func(x, y int) int { return x / y }
	fmt.Println(divider(second, first))

	sumFunc := func(x, y int) int {
		return x + y
	}

	subtractFunc := func(x, y int) int {
		return x - y
	}

	fmt.Println(Calculate(first, second, sumFunc))
	fmt.Println(Calculate(first, second, subtractFunc))

	dollar := 10

	getDollarValue := func() int {
		return dollar
	}

	fmt.Println(getDollarValue())
	dollar = 70

	fmt.Println(getDollarValue())
	divideBy2 := createDivider(2)
	divideBy10 := createDivider(10)

	fmt.Println(divideBy2(100))
	fmt.Println(divideBy10(100))
}

// Hello - приветствие
func Hello() {
	fmt.Println("Hello guys")
}

// HelloToOne - приветствие для одного
func HelloToOne(name string) {
	fmt.Printf("Zdarova %s\n", name)
}

// FullNameHello - приветствие с полным именем
func FullNameHello(name, surname string) {
	fmt.Printf("Hi %s %s\n", name, surname)
}

func Calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

func namedSumAndMultiply(first, second int) (sum int64, multiply int64) {
	sum = int64(first + second)
	multiply = int64(first) * int64(second)
	return // или return sum, multiply
}

func SumAndMultiply(first, second int) (int, int) {
	return first + second, first * second
}

func Sum(first, second int) int {
	sum := first + second
	return sum
}

func createDivider(divider int) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider
	}
	return dividerFunc
}
