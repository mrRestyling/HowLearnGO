package main

import "fmt"

func main() {
	elements()
}

func elements() {

	slice := []int{1, 2, 3, 4, 5} //
	i := 2

	// fmt.Println(slice)

	// withSlice := slice[:i+copy(slice[i:], slice[i+1:])]

	fmt.Println(slice)
	// fmt.Println(withSlice)

	// 1) 3 4 5
	// 2) 4 5
	// slice 1 2 4 5 5
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Println(slice)
}

// slice[:(2+2)]
// copy(slice[3,4,5],slice[4,5])
// slice = [1,2,4,5,5]
// withSlice = [1,2,4,5]
