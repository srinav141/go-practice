package hamming

import "errors"

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
