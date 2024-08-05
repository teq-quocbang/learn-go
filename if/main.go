package main

import "fmt"

func main() {
	// && => và, | |=>  hoặc
	// ==, !=, !variable must is bool
	// a := 1
	// b := 1
	// c := false
	// if !c {
	// 	fmt.Println("is false")
	// }
	d := "b"
	switch d {
	case "bang", "a", "b":
		fmt.Println("day la bang")
	case "quang":
		fmt.Println("day la quang")
	default:
		fmt.Println("default")
	}
}
