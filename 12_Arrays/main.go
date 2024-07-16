package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// arrays()
	slices()
}

func arrays() {

	var intArr [5]int
	intArr[0] = 5
	intArr[2] = 4
	intArr[4] = 3
	fmt.Printf("%T; %#v\n", intArr, intArr)

	people := [2]Person{
		{Age: 28, Name: "Egor"},
		{Age: 30, Name: "Daria"},
	}

	fmt.Printf("%T, %#v\n", people, people)

	stringsArr := [...]string{"1", "2", "3", "4"} // создаем массив
	// stringsArr := []string{"1", "2", "3", "4"} // создакем слайс
	fmt.Printf("%T, %#v\n", stringsArr, stringsArr)

	fmt.Printf("Len: %d, Cap: %d\n", len(stringsArr), cap(stringsArr))

	for i := 0; i < len(stringsArr); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, stringsArr[i])
	}

	for inx, value := range stringsArr {
		fmt.Printf("inx: %d, value: %s\n", inx, value)
	}

	for _, value := range intArr {
		fmt.Printf("value: %d\n", value)
	}

	newIntArr := changeArrays(&intArr) // массив передается в функцию по значению
	fmt.Println(intArr)
	fmt.Println(newIntArr)
}

func changeArrays(arr *[5]int) [5]int {
	arr[1] = 6
	return *arr
}

func slices() {
	var defaultSlice []int
	fmt.Printf("%T; %#v\n", defaultSlice, defaultSlice) // nil

	stringsSliceLiteral := []string{"1", "2"}
	fmt.Printf("%T; %#v\n", stringsSliceLiteral, stringsSliceLiteral)
	fmt.Printf("%d; %d\n", len(defaultSlice), cap(stringsSliceLiteral))

	sliceByMake := make([]int, 0, 5) // cap ? 5
	fmt.Printf("%T; %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("%d; %d\n", len(sliceByMake), cap(sliceByMake))

	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5, 6, 7) // если мы даем больше значений, чем вместимость
	// go увеличивает изначальный массим в два раза
	fmt.Printf("%T; %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("%d; %d\n", len(sliceByMake), cap(sliceByMake))

	for inx, value := range sliceByMake {
		fmt.Printf("inx: %d, value: %d\n", inx, value)
	}

}
