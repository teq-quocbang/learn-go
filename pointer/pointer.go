package pointer

import "fmt"

// & get dia chi cua bien
// * get gia chi cua dia chi do tren con RAM

// thay doi gia tri cua bien o bat ky dau
// nhẹ
func Pointer() {
	var a *int
	b := 1
	a = &b
	add(a, 10)
	fmt.Println(*a)

	// pointer in pointer
	aaa := "akdsas"
	bbb := &aaa

	c := &bbb

	fmt.Println(*c)
}

func add(input *int, num int) {
	*input += num
}
