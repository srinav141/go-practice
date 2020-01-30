package main

import "fmt"

func fibonacci() func() int {
	num1 := 0
	num2 := 1
	sum := 0
	return func() int {
		sum = num1 + num2
		num1 = num2
		num2 = sum
		return sum
	}

}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
