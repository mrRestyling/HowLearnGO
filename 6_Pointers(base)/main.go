package main

import "fmt"

// Указатели - это тип данных, которые в качестве значения
// хранят адрес ячейки памяти значения, либо другого указателя
// (может быть nil).
func main() {
	// Разыменование указателя:
	// *<имя_указателя>
	// <имя_указателя> = &<переменная>
	//   *&a = a

	// 1.
	fmt.Printf("\n1. Дефолтное значение указателя:\n")
	var intPointer *int
	fmt.Printf("%T %#v \n\n", intPointer, intPointer)

	//
	//

	// 2.
	fmt.Printf("2. Получение not-nil указателей:\n")
	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)

	fmt.Printf("получаем значение указателя: \n")
	var pointerA *int64 = &a // амперсант перед переменной
	fmt.Printf("%T %#v %#v \n\n", pointerA, pointerA, *pointerA)

	// 3.
	fmt.Printf("3. Получить указатель с помощью нового ключевого слова:\n")
	var newPointer = new(float32)
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
	*newPointer = 3
	fmt.Printf("%T %#v %#v \n\n", newPointer, newPointer, *newPointer)
}
