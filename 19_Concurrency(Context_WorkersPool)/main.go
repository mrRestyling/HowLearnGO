package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Context:
// 1) Хранить значения
// 2) Сообщать о завершении

func main() {
	// baseKnowledge()
	workerPool()
}

func baseKnowledge() {
	ctx := context.Background() // создает корневой контекст (context.Background() - родительский)
	fmt.Println(ctx)

	toDo := context.TODO() // больше для тестов
	fmt.Println(toDo)

	withValue := context.WithValue(ctx, "name", "Daria")
	fmt.Println(withValue.Value("name")) // лучше так не делать; считается антипаттерном
	fmt.Println()

	withCancel, cancel := context.WithCancel(ctx) // контекст создается на основе родительского
	fmt.Println(withCancel.Err())                 // метод контекста
	cancel()                                      // вызываем колбэк, который завершает наш контекст вручную
	// в момент когда наш контекст завершается - у него кладется ошибка о завершении контекста
	fmt.Println(withCancel.Err())
	// ! контекст нужно закрывать на том же уровне !
	fmt.Println()

	// создаем контекст, вызываем дедлайн
	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	fmt.Println(withDeadline.Deadline())
	fmt.Println(withDeadline.Err())
	fmt.Println(<-withDeadline.Done()) // ждем дедлайн
	fmt.Println()

	// Таймаут
	withTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()                    // контекст закрывается в двух случаях - таймаут или defer (при завершении функции)
	fmt.Println(<-withTimeout.Done()) // приходит пустая структура
}

// workerPool - один из паттернов использования конкурентности в GO
func workerPool() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(context.Background(), time.Millisecond*20) // создаем контекст на основе другого контекста
	defer cancel()

	wg := &sync.WaitGroup{}
	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)
	// numbersToProcess - тут будут лежать числа, которые нам нужно обработать
	// processedNumbers - числа, которые мы обработали(результаты)

	for i := 1; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			numbersToProcess <- i
			// if i == 500 {
			// 	cancel()
			// }
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()               // так мы помоймем, когда все воркеры закончили работу
		close(processedNumbers) // нужно закрывать канал, после того, как мы туда перестали писать
	}()

	var counter int // считает сколько мы получили значений
	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}
	fmt.Println(counter)

}

// способ сообщить о завершении
func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	// toProcess - отсюда будем брать цифры для обработки
	// processed - сюда будем складывать обработанные цифры

	for { // из-за select не можем использовать for range
		select {
		case <-ctx.Done(): // блокирующая
			return
		case value, ok := <-toProcess: // нам нужно проверять признак закрытости канала вручную
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}

}
