/*
Package collatzconjecture implements
*/
package collatzconjecture

import "errors"

//CollatzConjecture returns int
func CollatzConjecture(i int) (int, error) {

	if i <= 0 {
		return -1, errors.New("Not a valid input")
	}

	inc := 0
	for i != 1 {
		i = getValue(i)
		inc++

	}
	return inc, nil
}

func getValue(i int) int {
	if i%2 == 0 {
		return i / 2
	}
	return (3 * i) + 1
}
