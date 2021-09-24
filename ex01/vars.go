package main

import "fmt"

type FortyTwo struct{}

func main() {
	vars := []interface{}{
		"42",
		uint(42),
		int(42),
		uint8(42),
		int16(42),
		uint32(42),
		int64(42),
		false,
		float32(42),
		float64(42),
		complex64(42 + 0i),
		complex128(42 + 21i),
		FortyTwo{},
		[1]int{42},
		map[string]int{"42": 42},
		(*int)(nil),
		[]int{},
		(chan int)(nil),
		nil,
	}
	for _, v := range vars {
		printvar(v)
	}
}

func printvar(v interface{}) {
	vartype := fmt.Sprintf("%T", v)
	printa_or_an := "a"

	if vartype == "*int" {
		fmt.Printf("%p is an %s.\n", v, vartype)
		return
	}

	if vartype[0] == 'i' {
		printa_or_an = "an"
	}
	fmt.Printf("%v is %s %s.\n", v, printa_or_an, vartype)
}
