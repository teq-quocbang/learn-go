package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/teq-quocbang/learn-go/auth"
	"github.com/teq-quocbang/learn-go/interfaces"
)

type server struct {
	repo     interfaces.Repository
	authRepo interfaces.AuthRepository
}

func (s server) Login() error {
	return nil
}

func main() {
	// f := func() string {
	// 	return "my name"
	// }
	// MyFunc(f)
	repoClient := NewRepositoryClient(&gorm.DB{})
	authRepo := auth.NewAuthRepo(&gorm.DB{})
	s := server{
		repo:     repoClient,
		authRepo: authRepo,
	}

	err := s.authRepo.Login()
	if err != nil {
		log.Fatal(err)
	}

	user, err := s.repo.GetUserInfo(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	// method()
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

func FuncForReturn() string {
	return ""
}

func MyFunc(f func() string) (string, int, func() string, error) {
	// if err := o.Validate(); err != nil {
	// 	return "", 0, err
	// }

	// 01010101010101010010101010101010010101010010100101010100101010010101
	// OOP
	// FP
	name := f()

	fmt.Println(name)

	return "", 0, FuncForReturn, nil
}

type option struct {
	isAdmin       bool
	withMigration bool
}

type Options func(*option)

func WithAdmin() Options {
	return func(o *option) {
		o.isAdmin = true
	}
}

func WithMigrate() Options {
	return func(o *option) {
		o.withMigration = true
	}
}

// optional pattern
func NewServer(options ...Options) string {
	opt := &option{}

	for _, o := range options {
		o(opt)
	}

	println(opt)

	return ""
}

type MyObject struct {
	name string
}

var MyVar string

func MyFunc2(mo MyObject, moj MyObject) {
	fmt.Printf("mo: %v\n", mo)

	fmt.Printf("MyObject: %v\n", MyObject{})
}

// custom struct to slice
func method() {
	mos := []MyObject{
		{
			name: "a",
		},
		{
			name: "b",
		},
		{
			name: "b",
		},
		{
			name: "c",
		},
	}

	mo := MyObjects(mos)
	uniqueName := mo.RemoveDuplicate()
	fmt.Println(uniqueName)
}

// type, not slice, array, not interface => only struct
func (mo MyObject) GetObjectName() string {
	return ""
}

type MyObjects []MyObject

// RemoveDuplicate is remove duplicate name in object
func (mos MyObjects) RemoveDuplicate() map[string]struct{} {
	// remove duplicate and print
	uniqueName := make(map[string]struct{})
	// key-value
	// a - {}
	// b - {}
	// c - {}
	for _, mo := range mos {
		uniqueName[mo.name] = struct{}{}
	}
	return uniqueName
}

// concrete
type RepositoryClient struct {
	client *gorm.DB
}

func NewRepositoryClient(client *gorm.DB) *RepositoryClient {
	return &RepositoryClient{
		client: client,
	}
}

func (rc *RepositoryClient) GetUserInfo(id int) (interfaces.User, error) {
	return interfaces.User{
		Email:    "bangthong@teqnological.asia",
		Name:     "bang thong",
		Provider: "google",
	}, nil
}
