//Package hamming implements a simple function
//to calculate haomming distence between two
//DNA strings
package hamming

import "errors"

// Distance retruns hamming distance of 2 strings
func Distance(a, b string) (int, error) {

	diffCount := 0
	ar := []rune(a)
	br := []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("string lengths do not match")
	}

	for i := range ar {

		if string(ar[i]) != string(br[i]) {
			diffCount++
		}
	}
	return diffCount, nil

}
