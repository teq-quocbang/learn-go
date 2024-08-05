package main

import "fmt"

// for loop with recurrent
func recurrent(i int) {
	if i == 0 {
		return
	}

	fmt.Printf("day la i: %d\n", i)

	recurrent(i - 1)
}

func main() {
	// for
	ss := []string{"a", "0", "1", "2"}
	for idx, s := range ss {
		if s == "1" {
			return
		}
		fmt.Printf("index: %d, ad value: %s \n", idx, s)
	}

	// do something
	fmt.Println("do something")

	// de quy
	// recurrent(10)
}
