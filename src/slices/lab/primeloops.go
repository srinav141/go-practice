package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	for i := 2; i <= 25; i++ {
		res := getFactorsProdcut(i)
		if res == "" {
			fmt.Printf("%d is prime\n", i)
		} else {
			fmt.Println(i, "= ", res)
		}
	}
}

func isPrime(a int) bool {

	res := false
	for i := a / 2; i > 1; i-- {
		if a%i == 0 {
			res = true
		}
	}
	return res

}

func getFactorsProdcut(a int) string {
	res := ""
	for i := a / 2; i > 1; i-- {
		if a%i == 0 {
			b := a / i
			res = strconv.Itoa(i) + "*" + strconv.Itoa(b)
		}
	}
	return res

}

//Using slice
func getFactors(a int) []int {

	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return []int{i, a / i}
		}
	}
	return nil

}
