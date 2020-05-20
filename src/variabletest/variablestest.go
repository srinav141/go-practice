package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is varibles test main start")

	var test1 int
	fmt.Println(test1)
	test1 = 10
	fmt.Println(test1)
	var test2 string = "test2"
	fmt.Println(test2)
	test2 = "re-declare test2"
	fmt.Println(test2)
	test3 := "test3"
	fmt.Println(test3)

	var test4 = 6
	fmt.Println(test4)

	var Test6 = 9
	fmt.Println(Test6)
	test_func(Test6)
	fmt.Println(Test6)
	p := fmt.Println
	now := time.Now()
	p(now)
	year, month, day := now.Date()
	p(year)
	p(month)
	p(day)
	f := now.Format("01-02-2006")
	p(f)

}

func test_func(testo int) {
	fmt.Println(testo)
	testo += 10
	fmt.Println(testo)

}
