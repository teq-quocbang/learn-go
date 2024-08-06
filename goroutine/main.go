package main

import (
	"fmt"
	"sync"
	"time"
)

// concurrency => song song
// parallel => đa luồng

// giống nhau: có thể chạy nhiều task mà k cần phải đợi
// khác nhau: concurrency thì với 1 thread của CPU cho phép chạy nhiều task, còn đối với thằng parallel
// thì 1 thread chỉ chạy được 1 task nhưng có thể chạy nhiều thread cùng 1 lúc

var count int

func counter() {
	time.Sleep(time.Second * 10)
	count += 1
	fmt.Println(count)
}

func main() {
	// start task
	wg := sync.WaitGroup{}
	times := 10

	wg.Add(times)
	for i := 0; i < times; i++ {
		// execute task
		go func() {
			// step 1

			// step 2
			counter()

			// step 3
			wg.Done()
		}()
	}

	go func() {
		// 8080

	}()

	go func() {
		// 8081

	}()

	// wait to tasks done
	wg.Wait()

	fmt.Println("done")
}
