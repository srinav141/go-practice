package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/srinav141/go-practice/src/structs/issues"
)

func main() {

	result, err := issues.SearchIssues(os.Args[1:])
	if err != nil {

		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s %.65s\n", item.Number, item.User.Login, item.Title, item.Created)

	}
	pastDays := -365
	past := time.Now().AddDate(0, 0, pastDays)
	fmt.Println("Printing issues older than ", pastDays)
	for _, item := range result.Items {
		if item.Created.Before(past) {
			fmt.Printf("#%-5d %9.9s %.55s %.65s\n", item.Number, item.User.Login, item.Title, item.Created)
		}
	}

}
