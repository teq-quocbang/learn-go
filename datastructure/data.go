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

type Object struct {
}

type Object2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MyError struct {
	Code string
}

func (MyError) Error() string {
	return "err from my err"
}

func MyErr(i int) error {
	if i == 1 {
		return MyError{
			Code: "ajhkahskklsjl",
		}
	} else {
		return fmt.Errorf("err from fmt")
	}
}

// interface with multiple type
func printData(req interface{}) {
	i := 2
	err := MyErr(i)

	if e, ok := err.(MyError); ok {
		fmt.Println(e)
	} else {
		fmt.Println(err)
	}

	// parse to type
	if data, ok := req.(Object); ok {
		fmt.Println(data)
		// xu ly
	}

	if data, ok := req.(Object2); ok {
		fmt.Println(data)
		// xu ly
	}

	// use switch case
	switch t := req.(type) {
	case string:
		fmt.Printf("print with type string: %s \n", t)
	case int:
		fmt.Printf("print with type int: %d \n", t)
	default:
		fmt.Println("unexpected type")
	}
}

// Abstract
type iMethod interface {
	GetName() string
}

type BangServer struct {
	name string
}

// concrete
func (bs BangServer) GetName() string {
	return bs.name
}

type QuangServer struct {
	name string
}

func (qs QuangServer) GetName() string {
	// add more logic
	return qs.name
}

type MethodIsFunction struct {
}

func (mif MethodIsFunction) GetName() string {
	return "method is function"
}

type User struct {
	id       int
	email    string
	provider string
}

func NewUser(id int, email string) User {
	return User{
		id:    id,
		email: email,
	}
}

func (u User) IsGoogleProvider() bool {
	// logc phuc tap
	return u.provider == "google"
}

func (u User) IsGithubProvider() bool {
	return u.provider == "github"
}

func interfaceWithMethod() {
	var i iMethod
	// i = BangServer{
	// 	name: "Quoc Bang",
	// }
	i = &MethodIsFunction{}

	fmt.Println(i.GetName())
}

func Interface() {
	// khai bao
	// var i interface{}
	// i = "Bang"

	user := NewUser(1, "email@gmail.com")

	fmt.Println(user.IsGoogleProvider())

	// interfaceWithMethod()
}
