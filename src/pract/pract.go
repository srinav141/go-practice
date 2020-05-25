package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {


	diffCount := 0
	ar := []rune(a)
	br := []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("string lengths do not match")
	}

	for i := range ar{

		if string(ar[i]) != string(br[i]) {
			diffCount++
		}
	}
	return diffCount, nil

}

func main() {

	s1 := "aüa"
	s2 := "aaa"

	res,_:= Distance(s1,s2)
	fmt.Println(res)

}
