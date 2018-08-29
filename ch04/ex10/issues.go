// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch04/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %22s %.55s\n", item.Number, item.User.Login,
			ageCategory(item), item.Title)
	}
}

func ageCategory(issue *github.Issue) string {
	const monthInHours = 24 * 30
	age := time.Since(issue.CreatedAt)
	months := age.Hours() / monthInHours
	switch {
	case months < 1:
		return "less than a month old"
	case months > 12:
		return "more than a year old"
	default:
		return "more than a month old"
	}
}
