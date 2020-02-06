package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	counts := make(map[string]int)
	filesData := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Taking standard input")
		countLines(os.Stdin, counts, filesData)
	} else {

		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filesData)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filesData[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filesData map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		counts[t]++
		if !strings.Contains(filesData[t], f.Name()) {
			filesData[t] += f.Name() + " "
		}

	}
}
