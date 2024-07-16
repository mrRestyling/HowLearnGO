package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// variadicFunctions()
	// convertToArrayPointer()
	// passToFunction()
	// sliceWithNew()
	// getSlice()
	// copySlice()
	deleteElement()
}

func variadicFunctions() {
	showAllElements(1, 2)
	showAllElements(1, 2, 3, 4, 5, 6, 7)

	firstSlice := []int{5, 6, 7, 8}
	secondSlice := []int{9, 3, 2, 1}

	showAllElements(firstSlice...) //"вариативным параметром" передача отдельных аргументов
	// и сразу отправляет в функцию (5,6,7,8)

	newSlice := append(firstSlice, secondSlice...)
	fmt.Printf("%T, %#v\n", newSlice, newSlice)

}

func showAllElements(values ...int) { //данная функция может принять сколько угодно параметров "..." указанного типа
	// все эти элементы будут сложены в слайс values

	for _, val := range values {
		fmt.Println("Value:", val)
	}
	fmt.Println()
}

// слайс вообще является структурой
type _slice struct {
	elements unsafe.Pointer // указательна массив с определенным типом данных
	len      int            // количество элементов
	cap      int            // текущая вместимость
}

func convertToArrayPointer() {
	initialSlice := []int{1, 2}
	fmt.Printf("%T, %#v\n", initialSlice, initialSlice)
	fmt.Printf("len: %d; cap: %d\n", len(initialSlice), cap(initialSlice))
	fmt.Println()

	// конвертация слайса в указатель на массив
	intArray := (*[2]int)(initialSlice) // нужно чтобы длина совпадала
	fmt.Printf("%T, %#v\n", intArray, intArray)
	fmt.Printf("len: %d; cap: %d\n", len(intArray), cap(intArray))
}

// передача слайса в функцию
func passToFunction() {
	initialSlice := []int{1, 2}
	fmt.Printf("%T, %#v\n", initialSlice, initialSlice)
	fmt.Printf("len: %d; cap: %d\n", len(initialSlice), cap(initialSlice))
	// тк слайс внутри содержит ссылку на массив
	// то когда мы передаем слайс в функцию, то его значения копируются
	// копируется строчка с указателем на массив

	changeValue(initialSlice)
	fmt.Printf("%T, %#v\n", initialSlice, initialSlice)
	fmt.Printf("len: %d; cap: %d\n", len(initialSlice), cap(initialSlice))

	newSlice := append(initialSlice, 3)
	fmt.Printf("%T, %#v\n", newSlice, newSlice)
	fmt.Printf("len: %d; cap: %d\n", len(newSlice), cap(newSlice))

	newSlice2 := appendValue(newSlice)
	// fmt.Printf("%T, %#v\n", newSlice2, newSlice2)
	// fmt.Printf("len: %d; cap: %d\n", len(newSlice2), cap(newSlice2))
	_ = newSlice2

}

func changeValue(slice []int) {
	slice[1] = 15
}

func appendValue(slice []int) []int {
	slice = append(slice, 4, 5, 6)
	fmt.Printf("%T, %#v\n", slice, slice)
	fmt.Printf("len: %d; cap: %d\n", len(slice), cap(slice))

	return slice
}

// создание слайса через new
func sliceWithNew() {

	slicePointer := new([]int) // ф-я new возвращает указатель на какой-то тип (тут слайс интов)

	fmt.Printf("%T, %#v\n", slicePointer, slicePointer)
	fmt.Printf("len: %d; cap: %d\n", len(*slicePointer), cap(*slicePointer))

	newSlice2 := append(*slicePointer, 1)
	fmt.Printf("%T, %#v\n", newSlice2, newSlice2)
	fmt.Printf("len: %d; cap: %d\n", len(newSlice2), cap(newSlice2))
}

// Получение слайса от слайса/массива
func getSlice() {
	intArr := [...]int{1, 2, 3, 4, 5}

	fmt.Println()
	fmt.Printf("%T, %#v\n", intArr, intArr)
	fmt.Printf("len: %d; cap: %d\n", len(intArr), cap(intArr))

	intSlice := intArr[1:3] // Арислайсинг - получение слайса из массива

	fmt.Println()
	fmt.Printf("%T, %#v\n", intSlice, intSlice)
	fmt.Printf("len: %d; cap: %d\n", len(intSlice), cap(intSlice))

	fullSlice := intArr[:] // intArr[0:5] - все значения

	fmt.Println()
	fmt.Printf("%T, %#v\n", fullSlice, fullSlice)
	fmt.Printf("len: %d; cap: %d\n", len(fullSlice), cap(fullSlice))

	sliceFromSlice := fullSlice[:3]

	fmt.Println()
	fmt.Printf("%T, %#v\n", sliceFromSlice, sliceFromSlice)
	fmt.Printf("len: %d; cap: %d\n", len(sliceFromSlice), cap(sliceFromSlice))

	intArr[1] = 500

	fmt.Println()
	fmt.Printf("%T, %#v\n", intArr, intArr)
	fmt.Printf("%T, %#v\n", intSlice, intSlice)
	fmt.Printf("%T, %#v\n", fullSlice, fullSlice)
	fmt.Printf("%T, %#v\n", sliceFromSlice, sliceFromSlice)
}

// копирование встроенной функцией copy
func copySlice() {

	// // черновик
	// destination := []int{1, 2, 3}
	// source := []int{6, 7}

	// fmt.Println()
	// fmt.Println("Copied", copy(destination, source))
	// fmt.Printf("%T, %#v\n", destination, destination)
	// fmt.Printf("len: %d; cap: %d\n", len(destination), cap(destination))

	destination := make([]string, 0, 2)
	source := []string{"Egor", "Daria", "Vova"}

	fmt.Println()
	fmt.Println("Copied", copy(destination, source))
	fmt.Printf("%T, %#v\n", destination, destination)
	fmt.Printf("len: %d; cap: %d\n", len(destination), cap(destination))

	destination = make([]string, 2, 3)

	fmt.Println()
	fmt.Println("Copied", copy(destination, source))
	fmt.Printf("%T, %#v\n", destination, destination)
	fmt.Printf("len: %d; cap: %d\n", len(destination), cap(destination))

	destination = make([]string, len(source)) // правильный метод копирования для всех элементов

	fmt.Println()
	fmt.Println("Copied", copy(destination, source))
	fmt.Printf("%T, %#v\n", destination, destination)
	fmt.Printf("len: %d; cap: %d\n", len(destination), cap(destination))

	var defaultSlice []string

	fmt.Println()
	fmt.Printf("%T, %#v\n", defaultSlice, defaultSlice)
	fmt.Printf("len: %d; cap: %d\n", len(defaultSlice), cap(defaultSlice))

	fmt.Println()
	fmt.Println("Copied", copy(defaultSlice, source))
	fmt.Printf("%T, %#v\n", defaultSlice, defaultSlice)
	fmt.Printf("len: %d; cap: %d\n", len(defaultSlice), cap(defaultSlice))

	rightCopy := append(make([]string, 0, len(source)), source...)

	fmt.Println()
	fmt.Printf("%T, %#v\n", rightCopy, rightCopy)
	fmt.Printf("len: %d; cap: %d\n", len(rightCopy), cap(rightCopy))
}

// удаление элемента из слайса
func deleteElement() {

	slice := []int{1, 2, 3, 4, 5}
	i := 2 // индекс элемента, который мы хотим удалить

	fmt.Println()
	fmt.Printf("%T, %#v\n", slice, slice)
	fmt.Printf("len: %d; cap: %d\n", len(slice), cap(slice))

	// // 1. С помощью append

	// withAppend := append(slice[:i], slice[i+1:]...) // ломает исходный слайс

	// fmt.Println()
	// fmt.Printf("%T, %#v\n", withAppend, withAppend)
	// fmt.Printf("len: %d; cap: %d\n", len(withAppend), cap(withAppend))

	// fmt.Println()
	// fmt.Println(slice)
	// fmt.Printf("%T, %#v\n", slice, slice)
	// fmt.Printf("len: %d; cap: %d\n", len(slice), cap(slice))

	// 2. Лучше менять изначальный слайс
	slice = append(slice[:i], slice[i+1:]...)

	fmt.Println()
	fmt.Println(slice)
	fmt.Printf("%T, %#v\n", slice, slice)
	fmt.Printf("len: %d; cap: %d\n", len(slice), cap(slice))

	slice = []int{1, 2, 3, 4, 5}

	// withCopy := slice[ ]

	// fmt.Println()
	// fmt.Println(withCopy)

	withCopy := slice[:i+copy(slice[i:], slice[i+1:])]
	// copy возвращает количество скопированных элементов

	fmt.Println()
	fmt.Println(withCopy)
	// {1,2,4,5,5}
	// slice[:4]
	// {1,2,4,5}
	// {1,2}
	// [3,4,5] + [4,5]

}
