package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	start := time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("echo1 is ", s)
	fmt.Println(elapsed)

}

func echo2() {
	fmt.Println("echo2 results")
	start := time.Now()
	for i, v := range os.Args[1:] {
		fmt.Printf("index of %v is %v\n", v, i)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)

}

func echo3() {
	fmt.Println("echo3 results")
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}

func main() {
	// argsWithProg := os.Args
	// fmt.Println(argsWithProg[1])
	echo1()
	echo2()
	echo3()
}
