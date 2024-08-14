package main

import (
	"fmt"
	"log"
	"time"
)

// fmt, time, duration
func main() {
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
