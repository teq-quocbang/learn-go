package datatype

const (
	name string = "Bang" // cach khai bao co ban
)

// const khong the thay doi gia tri
const (
	RedisKey int = iota + 1 // de danh index cho bien
	AwsKey
	FBKey
)

type MyString string

const (
	QuangKey MyString = ""
	BangKey  MyString = "Bang"
)

func Data() {

}
