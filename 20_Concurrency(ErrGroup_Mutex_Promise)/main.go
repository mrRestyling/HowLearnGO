package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// chanAsPromise() // канал в качестве промисов (асинхронные запросы без блокировки для ожидания ответа)
	// chanAsMutex() // имитация мьютектов с  использованием каналов
	// withErrGroup()
	errGroup() // обертка над WG (не входит в стандартвую библиотеку)

}

func chanAsPromise() {
	// получаем два канала
	firstResponceChan := makeRequest(1)
	secondResponceChan := makeRequest(2)
	// (не блокируемся, не ждем)
	// далее выполняем какой-то код

	fmt.Println("нет блокировки")

	// Получаем ответы:
	fmt.Println(<-firstResponceChan, <-secondResponceChan)
}

func makeRequest(num int) <-chan string {

	responceChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		responceChan <- fmt.Sprintf("responce number %d", num)
	}()

	return responceChan
}

func chanAsMutex() {

	var counter int
	mutexChan := make(chan struct{}, 1) // исп. структуру потому что она ничего не весит(вес определяется по кол-ву полей)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {

		wg.Add(1) // в каждой итерации по одной задаче

		go func() {
			defer wg.Done()

			mutexChan <- struct{}{} // буфер 1 (занять сможет только одна горутина)

			counter++

			<-mutexChan
		}()

	}

	wg.Wait() // ждем пока завершатся
	fmt.Println(counter)
}

func withErrGroup() {

	var err error
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(3)

	// спит секунду(не успевает запуститься)
	// (контекст отменился - не нужно выполнять основную работу)
	go func() {

		time.Sleep(time.Second)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Первая запустилась")
			time.Sleep(time.Second)
		}
	}()

	// Имитация ошибки при выполнении запроса
	// * возвращает ошибку + отменяет контекст
	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:

			fmt.Println("Вторая запустилась")
			err = fmt.Errorf("Любая ошибка")
			cancel()
		}
	}()

	// Стандартный случай
	// (запускается и что-то делает)
	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Третья запустилась")
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println(err)
}

func errGroup() {
	// Некий синтаксический сахар
	// чтобы не писать withErrGroup
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("Первая запустилась")
			time.Sleep(time.Second)
			return nil
		}
	})

	g.Go(func() error {
		fmt.Println("Вторая запустилась")
		return fmt.Errorf("Ошибка во втором не найдена")
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			// можем тут не писать

		default:
			fmt.Println("Третья запустилась")
			time.Sleep(time.Second)
		}
		// а написать здесь
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
