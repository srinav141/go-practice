/*
Package darts implements score
*/
package darts

import "math"

//Score retruns int
func Score(i, j float64) int {

	dist := math.Sqrt(math.Pow((i-0), 2) + math.Pow((j-0), 2))

	if dist > 10.0 {
		return 0
	} else if dist <= 10.0 && dist > 5.0 {
		return 1
	} else if dist <= 5.0 && dist > 1.0 {
		return 5
	}
	return 10
}
