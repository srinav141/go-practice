/*Package scale implements Scale function*/
package scale

import (
	"strings"
	"unicode"
)

//Scale1 retruns the slice if tones
func Scale1(tonic string, interval string) []string {
	sharps := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flats := []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

	useSharps := []string{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "d#", "g#"}
	useFlats := []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}

	indexOfTonic := 0

	s, ok := find(useSharps, tonic)
	if ok {

		for i, v := range sharps {
			if v == tonic {
				indexOfTonic = i
			}
		}

		if interval == "" {
			res := sharps[indexOfTonic:]
			exp := len(sharps) - len(res)
			res = append(res, sharps[0:exp]...)
			return res
		} else {
			res := make([]string, 0)

			res = append(res, sharps[s])
			for j, d := range interval {
				if j == len(interval)-1 {
					break
				}
				if unicode.IsUpper(d) {
					if s == len(sharps)-1 {
						s = 1

					} else if s == len(sharps)-2 {
						s = 0
					} else {
						s += 2

					}

				} else {
					if s == len(sharps)-1 {
						s = 0

					} else {
						s++
					}

				}
				res = append(res, sharps[s])

			}
			return res
		}

	}

	f, ok := find(useFlats, tonic)
	if ok {

		for i, v := range flats {
			if v == tonic {
				indexOfTonic = i
			}
		}

		if interval == "" {
			res := flats[indexOfTonic:]
			exp := len(flats) - len(res)
			res = append(res, flats[0:exp]...)
			return res
		} else {

			res := make([]string, 0)

			res = append(res, flats[f])
			for j, d := range interval {
				if j == len(interval)-1 {
					break
				}
				if unicode.IsUpper(d) {
					if f == len(flats)-1 {
						f = 1

					} else if f == len(sharps)-2 {
						f = 0
					} else {
						f += 2

					}

				} else {
					if f == len(sharps)-1 {
						f = 0

					} else {
						f++
					}

				}
				res = append(res, flats[f])

			}

			return res
		}

	}
	return nil
}

func find(source []string, value string) (int, bool) {

	for i, item := range source {
		if strings.ToLower(item) == strings.ToLower(value) {
			return i, true
		}
	}
	return -1, false
}

var (
	chromatics = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flats      = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
)

func gen(t string, interval string, collection []string) []string {
	res := []string{}
	tonic := []rune(t)
	if tonic[0] >= 'a' && tonic[0] <= 'g' {
		tonic[0] = tonic[0] - 32
	}
	pos := 0
	for i := 0; i < 12; i++ {
		if strings.Compare(string(tonic), collection[i]) == 0 {
			pos = i
			break
		}
	}
	res = append(res, string(tonic))
	if strings.Compare(interval, "") == 0 {
		interval = "mmmmmmmmmmmm"
	}
	for _, v := range interval {
		switch v {
		case 'M':
			pos += 2
		case 'm':
			pos++
		case 'A':
			pos += 3
		}
		res = append(res, collection[pos%12])
	}
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	return res
}

// Scale returns
func Scale(tonic string, interval string) []string {
	res := []string{}
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		res = gen(tonic, interval, chromatics)
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		res = gen(tonic, interval, flats)
	}
	return res
}
