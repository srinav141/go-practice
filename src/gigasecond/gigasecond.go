/*
Package gigasecond implements AddGigasecond
*/
package gigasecond

import "time"

// AddGigasecond should adds giga seconds tp given time
func AddGigasecond(t time.Time) time.Time {
	secs := t.Unix()

	afterGiga := secs + 1000000000
	nt := time.Unix(afterGiga, 0)
	return nt
}
