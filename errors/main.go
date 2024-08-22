package main

import (
	"fmt"

	learn_go "github.com/teq-quocbang/learn-go/code"
)

type Error struct {
	Code   learn_go.Code
	Detail string
}

func (e Error) Error() string {
	msg := fmt.Sprintf("detail: %s", e.Detail)
	if e.Code != 0 {
		msg += fmt.Sprintf(" code: %d", e.Code)
	}
	return msg
}

func main() {
	err := getErr()
	str := err.Error()
	fmt.Println(str)
}

func getErr() error {
	return Error{
		Code:   learn_go.Code_INVALID_PARAMS,
		Detail: "invalid param",
	}
}
