package datastructure

import (
	"fmt"
	"math/rand"
)

func Array() {
	a := [5]int{1, 0, 9, 8}
	fmt.Println(a)
}

func Slice() {
	// cach khai bao khong biet size
	s := []int{1, 1, 2, 3, 4, 5, 4, 6}
	fmt.Println(s)

	// cach khai bao biet truoc size
	// []int{0,0,0,0,0,0,0,0,0,0}
	// high performance
	ss := make([]int, 10)
	for i := 0; i < 10; i++ {
		ss[i] = rand.Int()
	}

	// low performance
	var sss []int
	sss = append(sss, 1)

	fmt.Println(ss)
	fmt.Println(sss)

	emptySlice := []int{}                // empty slice
	emptySliceWithMake := make([]int, 0) // empty slice

	fmt.Printf("%v \n", len(emptySlice))
	fmt.Printf("%v \n", len(emptySliceWithMake))

	var nilSlice []int
	fmt.Printf("nil slice %v \n", len(nilSlice))

	var s1 []int         // nil slice
	s2 := []int{}        // non-nil, empty slice
	s3 := make([]int, 0) // non-nil, empty slice

	fmt.Println("s1", len(s1), cap(s1), s1 == nil, s1[:], s1[:] == nil)
	fmt.Println("s2", len(s2), cap(s2), s2 == nil, s2[:], s2[:] == nil)
	fmt.Println("s3", len(s3), cap(s3), s3 == nil, s3[:], s3[:] == nil)

	for range s1 {
	}
	for range s2 {
	}
	for range s3 {
	}

	// Slice is pointer also
	a := []int{1, 2, 3}
	fmt.Println(a) // 1,2,3
	b := a
	b[1] = 9
	fmt.Printf("pointer: %p and value: %d", b, b) // 1,9,3
	// 1,9,3

	// khong muon copy dia chi
	aa := []int{1, 2, 3}
	bb := make([]int, 3)
	i := copy(bb, aa)
	fmt.Println(i)
	bb[1] = 9
	fmt.Println(aa)
	fmt.Println(bb)
}

func Map() {
	// co the truy xuat du lieu nhanh voi toc do O^1
	m := make(map[int]*string)

	q := "Quang"
	b := "Bang"
	t := "Thai"
	m[1] = &q
	m[2] = &b
	m[3] = &b
	m[1] = &t

	if value, ok := m[1]; ok {
		fmt.Println(*value)
	}

	// loai bpo gia tri trung lap
	people := []string{"Quang", "Thai", "Bang", "Bang", "Quang"}
	m2 := make(map[string]int)
	for _, person := range people {
		m2[person] = 10
	}

	fmt.Println(m2)
}
