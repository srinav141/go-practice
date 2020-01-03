package main

import (
	"fmt"
	"unicode"
)

func main() {

	var a [3]int
	fmt.Println(a)
	a[1] = 10
	fmt.Println(a)

	b := []int{5, 6, 7}
	fmt.Println(b)
	fmt.Println(len(b))

	c := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(c); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, c[i])
	}

	sum := float64(0)
	for i, v := range c {
		fmt.Printf("Adding %dth element\n", i)
		sum += v
	}
	fmt.Printf("Sum is %.2f\n", sum)

	d := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}

	for _, v1 := range d {
		fmt.Println("v1 is %s\n", v1)
		for _, v2 := range v1 {
			fmt.Printf("%s\n", v2)
			//fmt.Println(v2[0])
		}
		fmt.Printf("\n")
	}

	r := [6]string{"c", "j", "a", "l", "n", "k"}
	revArray(r)
	s := [6]string{"c", "$", "a", "l", "$", "k"}
	revLetter(s)
}

func revArray(p [6]string) {
	fmt.Println(p)
	l := 0
	r := len(p) - 1
	for l < r {
		p[l], p[r] = p[r], p[l]
		l++
		r--
	}
	fmt.Println(p)
}

func revLetter(k [6]string) {
	fmt.Println(k)
	l := 0
	r := len(k) - 1

	for l < r {
		if hasSymbol(k[l]) {
			l++
		} else if hasSymbol(k[r]) {
			r--
		} else {
			k[l], k[r] = k[r], k[l]
			l++
			r--
		}
	}
	fmt.Println(k)

}

func hasSymbol(str string) bool {
	for _, letter := range str {
		if unicode.IsSymbol(letter) {
			return true
		}
	}
	return false
}
