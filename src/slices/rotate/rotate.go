package main

import "fmt"

func rotate(sl []int, i int) {
	for i > 0 {
		a := sl[0]
		copy(sl[:], sl[1:])
		//fmt.Println(sl)
		sl = append(sl[:len(sl)-1], a)

		i--
	}
	fmt.Println(sl)
}

func main() {
	sl := []int{0, 1, 2, 3, 4, 5}
	rotate(sl, 2)
}
