package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	baseSelect()
	gracefulShutdown()
}

func baseSelect() {
	fmt.Println("")

	brd := make(chan string, 3)
	brd <- "first"

	// select различает три вида операций:
	// 1. блокирующая операция
	// 2. неблокирующая операция
	// 3. ветка default

	// Select - анализирует и ищет неблокирующие операции (запись/чтение)
	// при *всевозможных* вызывается любой кейс (независимо от порядка)
	select {

	case str := <-brd: // достать запись из канала (2)
		fmt.Println("- Read:", str)
		fmt.Println("")
	case brd <- "second": // записать в канал новое значение (1)
		fmt.Println("- Write:", <-brd, <-brd)
		fmt.Println("")
	}

	unbrd := make(chan int)

	go func() { // записываем значение через секунку
		time.Sleep(time.Second)
		unbrd <- 1
	}()

	select {
	// case brd <- "third":
	// 	fmt.Println("- Unblocking writing")
	case val := <-unbrd:
		fmt.Println("- Blocking writing:", val)
	case <-time.After(time.Millisecond * 1500):
		fmt.Println("- Time's up")
		// default:
		// 	fmt.Println("- Default case")
	}

	fmt.Println("")

	resultChan := make(chan int)
	timer := time.After(time.Second)

	go func() {
		defer close(resultChan)

		for i := 1; i <= 1000; i++ {

			select {
			case <-timer:
				fmt.Println("Time's up!")
				return

			default:
				time.Sleep(time.Nanosecond)
				resultChan <- i

			}

		}

	}()

	for v := range resultChan {

		fmt.Println(v)
	}
}

func gracefulShutdown() {

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(time.Second * 10)

	for {
		select {
		case <-timer:
			fmt.Println("Time's up!")
			return
		case sig := <-sigChan:
			fmt.Println("Stopped by signal:", sig)
		}
	}
}
