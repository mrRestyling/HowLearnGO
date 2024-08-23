package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// AddMutex()
	// AddAtomic()
	// StoreLoadSwap()
	compareAndSwap()

	// atomicVal()
}

func AddMutex() {

	start := time.Now()

	fmt.Println(start.Format("15:04:05"))

	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 1000; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()

		}()

	}

	wg.Wait()

	fmt.Println(counter)

	fmt.Println(time.Now().Sub(start))
}

func AddAtomic() {

	start := time.Now()

	var counter int64
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 1; i <= 1000; i++ {

		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1)

		}()

	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start))
}

func StoreLoadSwap() {

	var count int32

	fmt.Println(atomic.LoadInt32(&count))

	atomic.StoreInt32(&count, 12)

	fmt.Println(atomic.SwapInt32(&count, 666))

	fmt.Println(atomic.LoadInt32(&count))
}

func compareAndSwap() {

	var (
		count int64 = 0
		wg    sync.WaitGroup
	)

	for i := 0; i < 10; i++ {

		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			fmt.Println(i)

			if !atomic.CompareAndSwapInt64(&count, 0, 9) {
				return
			}

			fmt.Println("Горутина поменяла свое значение:", i)
			fmt.Println(i)
		}(i)

	}
	wg.Wait()

	fmt.Println(count)
}
