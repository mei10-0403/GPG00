package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	c, _ := strconv.Atoi(os.Args[1])
	var n int = countn(c)
	createstairs(n)
}

func countn(c int) int {
	n := 1
	for c >= n {
		c -= n
		n++
	}
	return n - 1
}

func createstairs(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
