package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/teq-quocbang/learn-go/datastructure"
)

func main() {
	datastructure.Interface()

	str := "My_Name_Is Bang"
	fmt.Println(strcase.ToCamel(str))
}

// => sort package => const => var => init => main
