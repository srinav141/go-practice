package main

import "fmt"

func main() {
	fmt.Println("HI")

	seen := make(map[string]string)
	seen["one"] = "1"
	seen["two"] = "2"
	seen["three"] = "3"

	for k, v := range seen {
		fmt.Println(k, "-->", v)
	}
}
