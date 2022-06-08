package main

import (
	"fmt"
)

func main() {
	a, b, c := 100, 1.0e3, "a"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var val int32 = 100
	var val2 int64 = int64(val)
	fmt.Println(val2)
}
