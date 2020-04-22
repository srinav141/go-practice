/*
Package proverb implemets a function to list possible rhymes
*/
package proverb

import "fmt"

// Proverb returns list of rhymes
func Proverb(rhyme []string) []string {

	if len(rhyme) == 0 {
		return []string{}
	}
	a := make([]string, 0)
	s := ""
	for i := 0; i < len(rhyme); i++ {

		if i == len(rhyme)-1 {
			s = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
			a = append(a, s)
		} else {
			s = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
			a = append(a, s)
		}

	}

	return a

}
