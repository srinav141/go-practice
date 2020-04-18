/*
package triangle implements a function to detemine
type of traingle based on given sides
*/
package triangle

import (
	"math"
	"sort"
)

//Kind is a sting type
type Kind string

const (
	// NaT not a triangle
	NaT Kind = "NaT"
	// Equ equilateral triangle
	Equ Kind = "Equ"
	// Iso isosceles triangle
	Iso Kind = "Iso"
	// Sca scalene triangle
	Sca Kind = "Sca"
	// Deg degenerate triangle
	Deg Kind = "Deg"
)

// KindFromSides determines kind of the triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	s := []float64{a, b, c}
	sort.Float64s(s)

	if !isTriangle(s) {
		k = NaT
		return k
	} else if a == b && a == c {
		k = Equ
	} else if a != b && b != c && a != c {
		k = Sca
	} else if s[0] == s[1] || s[1] == s[2] {
		k = Iso
	} else if s[0]+s[1] == s[2] {
		k = Deg
	}

	/*
		if (a == b) || (b == c) || (a == c) {
				// If two sides are equal, check if all three are equal
				if (a == b) && (b == c) {
					k = Equ
				} else {
					k = Iso
				}
			} else {
				k = Sca
			}

	*/

	return k
}

func isTriangle(s []float64) bool {
	sort.Float64s(s)
	return s[0] > 0 && s[1] != math.Inf(1) && s[0]+s[1] >= s[2]
}
