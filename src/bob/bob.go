/*
Package bog implements Hey function
*/
package bob

import (
	"regexp"
	"strings"
)

// Hey will return reposnse bsed on remark
func Hey(remark string) string {
	res := ""

	var re = regexp.MustCompile("[\t|\n|\r|\" \"]")
	t := strings.TrimSpace(remark)
	s := re.ReplaceAllString(t, "")

	var rex = regexp.MustCompile("[0-9]|,|[\t\n\f\r ]")
	st := rex.ReplaceAllString(remark, "")

	if s == "" {
		res = "Fine. Be that way!"
	} else if strings.ToUpper(remark) == remark && string(remark[len(remark)-1]) == "?" && strings.ToLower(remark) != remark {
		res = "Calm down, I know what I'm doing!"

	} else if strings.HasSuffix(t, "?") {
		res = "Sure."
	} else if strings.ToUpper(remark) == remark && len(st) > 0 {
		res = "Whoa, chill out!"
	} else {
		res = "Whatever."
	}
	return res
}
