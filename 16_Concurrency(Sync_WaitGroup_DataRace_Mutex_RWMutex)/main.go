package main

import (
	"fmt"
	"sync"
	"time"
)

// пакет sync

// WaitGroup - механизм ожидания завершения группы задач

// DataRace - обращение к одним и тем же данным из разных программ

// Mutex и RWMutex - механизм получения исключительной блокировки
// m.Lock и m.Unlock

func main() {

	// _WaitGroup_
	// withoutWait()
	// withWait()
	// wrongAdd()

	// _Mutex_
	// writeWithoutConcurrent()
	// writeWithoutMutex()
	// writeWithMutex()
	// readWithMutex()
	readWithRWMutex()

}

// Тут нужно использовать инструмент синхронизации waitGroup
func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("exit")
}

func withWait() {

	var wg sync.WaitGroup // создаем WG - переменная, у которой тип WG из пакета sync

	wg.Add(10) // 2 вариант, если заранее известно количество задач

	for i := 0; i < 10; i++ {

		// wg.Add(1) // 1. варик здесь! добавляем опр-е количество задач, которые она должна подождать

		go func(i int) { // на каждую итерацию мы запускаем горутину
			//
			fmt.Println(i + 1)
			wg.Done() // обязательно, тут мы показываем, что добавленная таска завершается
		}(i)
	}
	wg.Wait() // горутина основная блокируется и ждет пока в WG не останется невыполненных задач
	fmt.Println("exit")
}

// пример !НЕПРАВИЛЬНОГО! добавления задач в WG (в горутину добавлять счетчик)
func wrongAdd() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		go func(i int) {
			wg.Add(1)

			defer wg.Done()

			fmt.Println(i + 1)
		}(i)
	}

	wg.Wait()
	fmt.Println("exit")
}

func writeWithoutConcurrent() {
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		counter++
	}

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

// DataRace (пример)
func writeWithoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup

	wg.Add(1000)
	for i := 0; i < 1000; i++ {

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)

	for i := 0; i < 1000; i++ {

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock() // работает только одна горутина
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mu.Unlock()

		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithRWMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.RWMutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.RLock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mu.RUnlock()

		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
