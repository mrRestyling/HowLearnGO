package main

import (
	"fmt"
	"sync"
	"time"
)

// Каналы - инструмент коммуникации, который позволяет обмениваться данными между горутинами

// type chan struct{
// 	mx sync.Mutex //--- что работает быстрее мьютекс или канал
// 	buffer []T //--- может быть, а может нет;
// 	readers []Goroutines //--- очередь читающей горутины
// 	writers []Goroutines //--- очередь пишущей горутины
// }

func main() {
	nilChanel()
	unbufferedChannel()
	bufferedChannel()
	forRange()
}

func nilChanel() {
	var nilChanel chan int // можно отправлять данные указанного типа
	fmt.Printf("Len: %d; Cap: %d\n", len(nilChanel), cap(nilChanel))

	// длина(len) - количество элементов, которые в данный момент находятся в буфере
	// cap - размер буфера канала

	// запись в nil канал (запись/получение будут заблокированы навечно)
	// nilChanel <- 1

	// запись в nil канал (чтение будет заблокировано навечно)
	// <-nilChanel

	// // закрытие нил канала (паника)
	// close(nilChanel)
}

func unbufferedChannel() {
	// // создание двунаправленного канала
	unb := make(chan int)

	// // создание канала только для записи
	// unb2 := make(chan<- int)

	// // создание канала только для чтения
	// unb3 := make(<-chan int)

	fmt.Printf("Len: %d; Cap: %d\n", len(unb), cap(unb))

	var wg sync.WaitGroup

	wg.Add(3)

	// 	// blocks until smb reads
	// 	// очередь на запись - дедлок
	// 	// обязательное условие - и читатель и писатель
	// unb <- 1

	// // blocks until smb write
	// <-unb

	// blocks on reading then write
	go func(chanForWrite chan<- int) {
		time.Sleep(time.Second)
		// <-chanForWrite
		unb <- 1 // или chanForWrite <- 1
		wg.Done()
	}(unb)

	v := <-unb
	fmt.Println(v)

	go func(chanForRead <-chan int) {
		time.Sleep(time.Second)
		ex := <-chanForRead
		fmt.Println(ex)
		wg.Done()
	}(unb)

	unb <- 6

	// закрытие канала
	// запись в закрытый канал - будет паниковать!
	// go func() {
	// 	time.Sleep(time.Millisecond * 500)
	// 	close(unb)
	// }()

	go func() {
		// time.Sleep(time.Second)
		v2 := <-unb
		fmt.Println(v2)
		wg.Done()
	}()

	unb <- 235

	wg.Wait()

	// // закрытие закрытого канала - ПАНИКА
	// close(unb)
	// close(unb)

}

func bufferedChannel() {

	// представление:
	// bufferedChannel {
	// len(buffer) > 0; [1,2] <- сюда
	// readers []Go
	// writers []Go [] <- не сюда
	// }

	bfc := make(chan int, 2)
	fmt.Printf("Len: %d; Cap: %d\n", len(bfc), cap(bfc))

	bfc <- 1
	bfc <- 2

	// bfc <- 3 // будет deadlock!

	fmt.Printf("Len: %d; Cap: %d\n", len(bfc), cap(bfc))

	fmt.Println(<-bfc)
	fmt.Println(<-bfc)

	fmt.Printf("Len: %d; Cap: %d\n", len(bfc), cap(bfc))

}

// использование циклов для получения значений из канала
func forRange() {
	bfc := make(chan int, 3)

	numbers := []int{5, 6, 7, 8}

	go func() {
		for _, num := range numbers {
			bfc <- num
		}
		close(bfc) // КАНАЛ нужно закрывать там, где мы в него записываем
	}()
	var arr []int
	for {
		// v := <-bfc
		// fmt.Println(v)

		v, ok := <-bfc
		fmt.Println(v, ok)
		if !ok {
			break
		}

		arr = append(arr, v)

	}
	fmt.Println(arr)

	bfc = make(chan int, 3) // мы создаем новый канал,тк старый закрыт

	go func() {
		for _, num := range numbers {
			bfc <- num
		}
		close(bfc)
	}()

	// когда мы используем for-range при чтении значения из канала
	// нам не нужно проверять закрыт канал или нет
	for v := range bfc {
		fmt.Println("buffered", v)
	}

	// условие обмена остается read = write
	unbfr := make(chan int)
	go func() {
		for _, num := range numbers {
			unbfr <- num
		}
		close(unbfr)
	}()

	for value := range unbfr {
		fmt.Println("unbfr", value)
	}
}
