package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// fmt, time, duration
func main() {
	return
	// fmt.Printf("print %s %s\n", "second", "third")
	// fmt.Println("print line")
	// str := "my name is: " + "bang"
	// name := fmt.Sprintf("my name is: %s, age: %d, f: %f", "bang", 24, 1.2)
	// a := []byte("abc")
	// result := fmt.Append(a, "d")
	// fmt.Println(string(result))

	// fmt.Println("end")

	// time
	n := time.Now() // according machine time
	fmt.Printf("before: %v \n", n)
	lo, err := time.LoadLocation("Europe/Budapest")
	if err != nil {
		log.Fatal(err)
	}

	nowAfter := n.In(lo)
	fmt.Printf("after: %v \n", nowAfter)

	lo2, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		log.Fatal(err)
	}

	hcm := nowAfter.In(lo2)
	fmt.Printf("reconvert: %v\n", hcm)

	// duration la khoang thoi gian
	// add 1 hour
	moreThanOneHour := hcm.Add(time.Hour * 1)
	fmt.Printf("time after add one hour: %v \n", moreThanOneHour)

	// in ra thoi gian tu a - b chenh lech bao nhieu giay
	// a ----------------- b ------------------c
	a := time.Now()
	b := a.Add(time.Second * 10)
	diff := b.Sub(a)
	fmt.Println(diff.Seconds())

	// unix := b.Unix()
	//
	d, err := time.ParseDuration("1m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
}

// io, os
func ioPkg() {
	bs := make([]byte, 100)
	for i := 0; i < 99; i++ {
		bs[i] = byte(i * 10)
	}

	// dung de copy data tu file A -> B
	a := &bytes.Buffer{}
	numberOfCopied, err := io.Copy(a, bytes.NewBuffer(bs))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numberOfCopied)

	fmt.Println(a)

	// b, err := io.ReadAll(a)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(a.Bytes())

	_, err = io.WriteString(a, "c")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a.Bytes())
}

func osPkg() {
	// lenh chmod la lenh de cap quyen
	// os.Chmod()
	// err := os.Chdir("common")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, err := os.ReadFile("data.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(data)

	// TAO VA GHI FILE
	{
		err := os.Chdir("../")
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create("tmp/myfile.text")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		str := "tui ten la bang"
		_, err = file.Write([]byte(str))
		if err != nil {
			log.Fatal(err)
		}
	}

	// TAO FILE DUONG DAN TAM
	{
		// f, err := os.CreateTemp("", "test.txt")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer f.Close()
		// str := "I'm groot"
		// f.Write([]byte(str))
	}

	{
		// mv
		// touch
		// mkdir
		//...
	}

	// add logic
}

// strconv
func strconvPkg() {
	// string to int
	{
		i, err := strconv.Atoi("123")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i)
	}

	// int to string
	{
		str := strconv.Itoa(123434)
		fmt.Println(str == "123434")
	}

	{
		data := `{
			"name": "bang",
			"age": 24
		}`

		str := strconv.QuoteToASCII(data)
		fmt.Println(str)
	}
}

func str() {
	// string builder
	{
		bd := strings.Builder{}
		a := "a"
		b := "b"
		bd.WriteString(a)
		bd.WriteString(b)
		fmt.Println(bd.String())
	}

	// replace
	{
		a := "abcdf"
		str := strings.Replace(a, "bc", "aa", -1)
		fmt.Println(str)
	}

	{
		a := "&toi la bang@"
		s := strings.Trim(a, "&@")
		fmt.Println(s)
	}

	{
		a := "  ss sss sss  "
		fmt.Println(strings.TrimSpace(a))
	}
	{
		a := "great"
		fmt.Println(strings.ToUpper(a))
	}
	{
		a := "gReAt"
		fmt.Println(strings.ToLower(a))
	}

	{
		a := "toi la bang"
		ss := strings.Split(a, " ")
		fmt.Println(ss)
	}
}
