package main

import "fmt"

type Object[T any] struct {
	Result T
}

// abstract
type IMethod interface {
	GetName(string) string // empty
}

type name struct {
	name string
}

func (n name) GetName(s string) string {
	s = s + "sdf"
	return s
}

func NewName(n string) name {
	return name{
		name: n,
	}
}

func main() {
	o := Object[string]{}
	o.Result = "string"

	o2 := Object[int]{}
	o2.Result = 1

	// n :=  NewName()
	var n = name{name: "Bang"}
	o3 := Object[IMethod]{}

	// interface = struct(trong struct nay phai co method)
	// {interface(struct)}
	o3.Result = n
	// struct = struct(co method)
	data := o3.Result.GetName("uuuu")
	fmt.Println(data)
}
