package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int = 1
	var num2 int64 = 1
	var num3 int32 = 1
	var num4 int16 = 1
	var num5 int8 = 1

	fmt.Println("Size of num1 (int):", unsafe.Sizeof(num1))
	fmt.Println("Size of num2 (int64):", unsafe.Sizeof(num2))
	fmt.Println("Size of num3 (int32):", unsafe.Sizeof(num3))
	fmt.Println("Size of num4 (int16):", unsafe.Sizeof(num4))
	fmt.Println("Size of num5 (int8):", unsafe.Sizeof(num5))
}
