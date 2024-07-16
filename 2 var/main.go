package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int = 1
	var num2 int8 = 12
	var num3 int16 = 1
	var num4 int32 = 1
	var num5 int64 = 1

	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))
	fmt.Println(unsafe.Sizeof(num3))
	fmt.Println(unsafe.Sizeof(num4))
	fmt.Println(unsafe.Sizeof(num5))
}
