package main

import "fmt"

func main() {
	f := func() string {
		return "my name"
	}
	MyFunc(f)
}

type Object struct {
	Name string
	Age  int
}

func (o Object) Validate() error {
	if o.Name == "" {
		return fmt.Errorf("missing name")
	}
	return nil
}

func MyFunc(f func() string) (string, int, error) {
	// if err := o.Validate(); err != nil {
	// 	return "", 0, err
	// }

	// 01010101010101010010101010101010010101010010100101010100101010010101
	// OOP
	// FP
	name := f()

	fmt.Println(name)

	return "", 0, nil
}
