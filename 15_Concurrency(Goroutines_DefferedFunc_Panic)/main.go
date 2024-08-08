package main

import (
	"fmt"
)

//  Размер стека горутины 2Кб ; Размер треда > 1мб
//  горутина - легковестный поток (еще называют гринтредами)
//  параллельно и конкурентно
// планировщик GO управляет горутинами
//  ОС управляет потоками

// Программист не создает треды сам, работа с тредами происходит под капотом

// main как минимум одна горутина
func main() {

	// defer складываюются в стек ( в обратном порядке)
	// defer fmt.Println("defer3")
	// defer fmt.Println("defer2")
	// defer fmt.Println("defer")

	//  defer из функции SUM
	// fmt.Println(sum(2, 3))

	// вычисление значений defer функций (2 примера)

	// deferValues()

	// // посмотреть сколько логических ядер
	// fmt.Println(runtime.NumCPU())

	// // поменять количество горутин, которые выполняются одновременно
	// runtime.GOMAXPROCS(1)

	// go showNumber(100)

	// вручную переключиться на другую горутину
	// runtime.Gosched()

	// переключениями между горутинами занимается планировщик
	// time.Sleep(time.Second)

	// обработка паники
	// makePanic()

	// fmt.Println("exit1")

	var x = 1
	fmt.Println(x)
	defer func() {
		x++
		fmt.Println(x)
	}()

	defer func() {
		x++
		fmt.Println(x)
	}()

	defer fmt.Println(x)
	fmt.Println(x + 1)

}

func showNumber(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()
	sum = x + y
	return
}

func deferValues() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("first", i)
	}
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("second", i)
		}()
	}

	for i := 0; i < 10; i++ {
		k := i
		defer func() {
			fmt.Println("third", k)
		}()
	}
	for i := 0; i < 10; i++ {
		defer func(k int) {
			fmt.Println("fourth", k)
		}(i)
	}
}

func makePanic() {
	// отловить панику и завершить программу корректно
	defer func() {
		panicValue := recover() 
		fmt.Println(panicValue)
	}()

	panic("some panic")
	fmt.Println("codein")
}
