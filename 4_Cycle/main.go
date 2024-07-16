package main

import "fmt"

func main() {
	for numbers := 0; numbers < 10; numbers++ {
		if numbers%2 == 1 {
			continue
		}
		fmt.Println(numbers)
	}

	// Label1: // !!! метка для цикла !!!

	// 	for i := 1; i <= 20; i++ {
	// 		for j := 1; j <= 3; j++ {
	// 			if i > 3 {
	// 				continue Label1
	// 			}
	// 			fmt.Println("I:", i, "J:", j)

	// 		}
	// 	}

	// цикл брейка
	// for i := 1; i < 20; i++ {
	// 	if i > 10 {
	// 		break
	// 	}
	// 	fmt.Println("I:", i)
	// }

	// Label2: // !!! метка для цикла !!!

	// 	for i := 1; i <= 20; i++ {
	// 		for j := 1; j <= 10; j++ {
	// 			fmt.Println("I:", i, "J:", j)
	// 			if i >= 10 {
	// 				break Label2
	// 			}
	// 		}
	// 	}
	// после значения брейка метки Label2 код идет дальше вниз
}

// package kata

// import (
//   "strings"
//   "strconv"
//   )

// func FreqSeq(str string, sep string) string {

//   s := []rune(str)

//   slise := []string{}

//   for _, w := range s {

//       amt := 0

//     for _, j := range s {

//       if w == j {
//         amt++
//       }

//     }
//     slise = append(slise, strconv.Itoa(amt))

//   }
//   return strings.Join(slise, sep)
// }
