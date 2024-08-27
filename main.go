package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100000; i++ {
		fmt.Println("hello world")
		time.Sleep(time.Second * 1)
	}
}

// => sort package => const => var => init => go main
