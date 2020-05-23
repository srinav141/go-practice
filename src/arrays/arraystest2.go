package main

import (
	"fmt"
)

func main() {

	a := [3]int{1, 2, 3}
	b := a
	fmt.Printf("Type is %T\n", a)
	fmt.Println(a)
	fmt.Println(b)
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

}
