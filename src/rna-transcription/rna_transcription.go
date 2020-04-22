/*
Package strand implements method to converts DNA to RNA strand
*/
package strand

//ToRNA converts DNA strand to RNA
func ToRNA(dna string) string {

	m := make(map[string]string)

	m["C"] = "G"
	m["G"] = "C"
	m["T"] = "A"
	m["A"] = "U"

	if len(dna) == 0 {
		return ""
	}
	res := ""

	for _, r := range dna {
		res += m[string(r)]
	}

	return res
}

// import (
// 	"strings"
// )

// var nucleotides = map[rune]rune{
// 	'G': 'C',
// 	'C': 'G',
// 	'T': 'A',
// 	'A': 'U',
// }

// /*
// ToRNA converts nucleotides of a given DNA string to RNA nucleotides
// */
// func ToRNA(dna string) string {
// 	mapper := func(r rune) rune {
// 		return nucleotides[r]
// 	}
// 	return strings.Map(mapper, dna)
// }
