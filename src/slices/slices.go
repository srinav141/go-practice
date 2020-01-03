package main

import (
	"fmt"
)

func main() {

	cars := []string{"Ferrari", "Ford", "Honda"}
	fmt.Printf("cars: %v,has length %d and capacity %d\n", cars, len(cars), cap(cars))
	cars = append(cars, "TATA")
	fmt.Printf("cars: %v,has length %d and capacity %d", cars, len(cars), cap(cars))
 	fmt.Println()
	st:=	"abc"
	fmt.Println(st[:2])

}
