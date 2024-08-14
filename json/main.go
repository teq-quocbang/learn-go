package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type LoginRequest struct {
	Username string `json:"username"`
	Pass     string `json:"password"`
}

func main() {
	// unmarshal with string
	// str := `{"username": "bang","password": "123"}`
	// var lr LoginRequest // username: "", password: ""

	// unmarshal with []byte
	// data, err := readFile()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := json.Unmarshal(data, &lr); err != nil {
	// 	log.Fatal(err)
	// }

	// struct to byte
	d, err := convertStructToByte()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)
}

func convertStructToByte() ([]byte, error) {
	lr := LoginRequest{
		Username: "bằng",
		Pass:     "13239u3eihjs",
	}
	return json.Marshal(lr)
}

func readFile() ([]byte, error) {
	return os.ReadFile("data.json")
}
