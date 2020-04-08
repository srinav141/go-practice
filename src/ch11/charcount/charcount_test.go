package charcount

import "testing"

func TestCharCount(t *testing.T) {
	var tests = []struct {
		intput string
		want   int
	}{{"test", 4}, {"abc", 3}, {"tester123 456", 12}, {"⌘c", 2}}

	for _, test := range tests {
		if got := getCharCount(test.intput); got != test.want {
			t.Errorf("getCharCount(%q) = %v", test.intput, got)
		}
	}
}
