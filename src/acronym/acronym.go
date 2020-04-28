/*
Package acronym implements Abbrevaite function
*/
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate return abbreviation of given string.
func Abbreviate(s string) string {
	result := ""
	wArray := strings.Split(s, " ")

	for _, w := range wArray {
		result += getAcronymChar(w)
	}

	return result
}

func getAcronymChar(st string) string {

	wSplit := strings.Split(st, "-")
	res := ""

	if len(wSplit) > 1 {
		for _, word := range wSplit {
			res += strings.ToUpper(getAcronymChar(word))
		}

	} else {
		var re = regexp.MustCompile("[-|_]")

		replaced := re.ReplaceAllString(st, "")
		if len(replaced) > 0 {

			res = strings.ToUpper(string(strings.TrimSpace(replaced)[0]))
		}
	}

	return res
}

// Abbreviate2 implements
func Abbreviate2(s string) string {
	var result string
	illegalCharactersMatch, _ := regexp.Compile(`[^a-zA-Z']`)
	words := strings.Fields(illegalCharactersMatch.ReplaceAllString(s, " "))

	for i := range words {
		result += string(words[i][0])
	}
	return strings.ToUpper(result)
}
