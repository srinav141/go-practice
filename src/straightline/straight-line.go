package main

import (
	"fmt"
)

func checkStraightLine(coordinates [][]int) bool {

	p1 := coordinates[1]
	p2 := coordinates[0]
	slope := getSlope(p1, p2)
	//fmt.Println(slope)

	for i := 1; i < len(coordinates); i++ {
		point1 := coordinates[i]
		point2 := coordinates[i-1]
		s := getSlope(point1, point2)

		if s != slope {
			return false
		}
	}

	return true

}

func getSlope(p1, p2 []int) float64 {

	if p1[0]-p2[0] == 0 {

		return 0
	}

	y := float64(p2[1]) - float64(p1[1])
	x := float64(p2[0]) - float64(p1[0])

	return (y / x)

}

func main() {

	// s := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	// r := checkStraightLine(s)
	// fmt.Println(r)

	// s = [][]int{{-3, -2}, {-1, -2}, {2, -2}, {-2, -2}, {0, -2}}
	// r = checkStraightLine(s)
	// fmt.Println(r)

	s := [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
	r := checkStraightLine(s)
	fmt.Println(r)

	s = [][]int{{-4, -3}, {1, 0}, {3, -1}, {0, -1}, {-5, 2}}
	r = checkStraightLine(s)
	fmt.Println(r)
}
