package main

import "fmt"

func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := make([][]int, 0)

	i := 0
	j := 0

	for i < len(A) && j < len(B) {

		a := A[i]
		b := B[j]

		if a[0] < b[0] && a[1] < b[0] {
			i += 1
		} else if a[0] > b[1] && a[0] > b[0] {
			j += 1
		} else {
			st := b[0]
			if a[0] > b[0]{
				st=a[0]
			}
			en := 0
			if a[1] >= b[1] {

				en = b[1]
				j += 1
			} else {
				en = a[1]
				i += 1
			}

			r := []int{st, en}
			res = append(res, r)

		}

	}

	return res
}

func main() {

	A := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	B := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}

	r := intervalIntersection(A, B)
	fmt.Println(r)

	a:= "a"
	fmt.Println(string([]rune(a)))

}
