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
	time.Sleep(time.Second * 1)
	count += 1
	fmt.Println(count)
}

func counter2(mt *sync.Mutex) {
	time.Sleep(time.Second * 5)
	mt.Lock()
	defer mt.Unlock()
	for i := 0; i < 10; i++ {
		count++
		fmt.Printf("counter: %d\n", count)
	}
}

// cho 1 map[string][]string nhiều user []string sẽ chứ các sở thích của từng người, yêu cầu truyền vào 1 sở thích bâts kỳ và trả
// lại những người có cùng sở thích
// user_id ["music", "ass"]
// user_2  ["m", "a", "e"]

// sql:
// select * có sử dụng join
// cho 1 chuỗi "Bang: " Quoc", yêu cầu select loại bỏ dấu ": và '"' "
// select replace(data, <chuoi can thay the>, <data thay the>)

func main() {
	// start task
	// wg := sync.WaitGroup{}
	// times := 10
	// mt := &sync.Mutex{}

	// wg.Add(times)
	// for i := 0; i < times; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		counter2(mt)
	// 	}()
	// }

	ch1 := make(chan string, 3)
	ch2 := make(chan float32, 3)
	// a := new(string)

	m := map[string]struct{}{}      // empty
	m1 := make(map[string]struct{}) // empty
	var m2 map[string]struct{}      // nil

	m["bang"] = struct{}{}
	m1["bang"] = struct{}{}
	m2["bang"] = struct{}{}

	return

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- 1
			ch1 <- "channel 1"
		}
		close(ch1)
		close(ch2)
	}()

	// for {
	func() {
		select {
		case <-ch1:
			for ch := range ch1 {
				fmt.Printf("channel 1, %s \n", ch)
			}
		case <-ch2:
			for ch := range ch2 {
				fmt.Printf("channel 2, %f \n", ch)
			}
		}
	}()
	// }
	// for i := 0; i < times; i++ {
	// 	// execute task
	// 	go func() {
	// 		// step 1
	// 		// i := 1 // 1
	// 		defer wg.Done()
	// 		// defer func(input int) {
	// 		// 	fmt.Printf("i value %d \n", input)
	// 		// }(i)

	// 		i++ // 2
	// 		// step 2
	// 		counter()
	// 		// return

	// 		// step 3
	// 	}()
	// }

	// wait to tasks done
	// wg.Wait()

	// TODO: worker with channel
	// numJobs := 10
	// jobs := make(chan int, numJobs)
	// results := make(chan int, numJobs)

	// for i := 0; i < 10; i++ {
	// 	go worker(i, jobs, results)
	// }

	// for i := 0; i < 10; i++ {
	// 	// time.Sleep(time.Second * 20)
	// 	jobs <- i
	// }
	// // close(jobs)

	// time.Sleep(time.Second * 4)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

// dev - stagging - master

// http://my-domain.dev.com // dev test
// http://my-domain.stagging.com // QC test
// https://my-domain.com

// dev -> feat/TEL-424
// git checkout -b feat/TEL-424 -> push code to git -> merge to dev -> CI/CD -> CI: test code, build image
