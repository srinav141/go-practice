package charcount

import (
	"fmt"
	"unicode"
)

// func main() {

// 	//fmt.Println(os.Args[0])
// 	s := ""
// 	for i := 1; i < len(os.Args); i++ {
// 		s += os.Args[i]
// 	}
// 	count := getCharCount(s)
// 	fmt.Println()
// 	fmt.Println("Char count is ", count)
// }

//GetCharCount gets char gount of given string
func getCharCount(s string) (count int) {
	count = 0
	for _, r := range s {
		fmt.Printf("%v ", r)
		if unicode.IsSpace(r) {
			continue
		}
		count++
	}
	// for i := 1; i < len(os.Args); i++ {
	// 	for _, r := range s {
	// 		fmt.Printf("%v ", r)
	// 		count++
	// 	}
	// }
	return
}
