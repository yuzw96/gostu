package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
)

//这里我们引用的并不是自己编写的结构体，而是来自gopl.io -_-
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
