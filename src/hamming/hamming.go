//Package hamming implements a simple function
//to calculate haomming distence between two
//DNA strings
package hamming

import "errors"

// Distance retruns hamming distance of 2 strings
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("string lengths do not match")
	}

	diffCount := 0
	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			diffCount++

		}
	}

	return diffCount, nil

}
