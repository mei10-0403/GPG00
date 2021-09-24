package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG = "Arguments is invalid."

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}

func calculationStr(args []string) (string, bool) {
	n := len(args)
	if n != 3 {
		return "error", false
	}
	a, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return "error", false
	}
	b, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil || b == 0 {
		return "error", false
	}
	str := fmt.Sprintf("sum: %d\ndifference: %d\nproduct: %d\nquotient: %d\n", a+b, a-b, a*b, a/b)
	return str, true
}
