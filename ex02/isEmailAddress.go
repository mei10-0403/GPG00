package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid argument")
	}
	r := regexp.MustCompile(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`)

	for _, args := range os.Args[1:] {
		ms1 := r.MatchString(args)
		if len(args) > 256 || !ms1 {
			fmt.Printf("%s is NOT a valid email address.\n", args)
		} else {
			fmt.Printf("%s is a valid email address.\n", args)
		}
	}
}
