package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const (
		min = 1
		max = 5
	)

	// rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел (устарел с Go 1.20)
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	value := rnd.Intn(max-min) + min

	// // 1. switch обычный
	fmt.Printf("\n(1)switch обычный:\n")
	switch value {
	case 1:
		fmt.Printf("%d\n\n", value)
	case 2, 3:
		fmt.Printf("2 or 3\n\n")
	case getFour():
		fmt.Printf("%d\n\n", value)
	default:
		fmt.Printf("defaul case is showed")
	}

	// 2. локальная переменная
	fmt.Println("(2) switch локальная переменная:")
	switch num := rnd.Intn(max-min) + min; num {
	case 1:
		fmt.Printf("%d\n\n", num)
	case 2, 3:
		fmt.Printf("%d or %d\n\n", 2, 3)
	case getFour():
		fmt.Printf("%d\n", num)
		fallthrough // переход к следующему case
	case 10:
		fmt.Printf("Strange things happen here \n\n")
	default:
		fmt.Printf("default case is showed \n\n")
	}

	// 3. switch без условия
	fmt.Println("(3) switch без условия:")
	switch {
	case value > 2:
		fmt.Printf("value %d greater than 2\n\n", value)
	case value < 2:
		fmt.Printf("value %d less than 2\n\n", value)
	default:
		fmt.Printf("value %d equals to 2\n\n", value)
	}

}
func getFour() int {
	fmt.Println("getFour is called")
	return 4
}
