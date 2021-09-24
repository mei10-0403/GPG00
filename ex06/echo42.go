package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	n := flag.Bool("n", false, "omit trailing newline")
	s := flag.String("s", " ", "separator")

	flag.Parse()
	if *n == true {
		str := strings.Join(flag.Args(), *s)
		fmt.Print(str)
	} else {
		str := strings.Join(flag.Args(), *s)
		fmt.Print(str)
		fmt.Print("\n")
	}
}
