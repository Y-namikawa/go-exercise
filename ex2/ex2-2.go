package main

import (
	"fmt"
	"strconv"
)

func main() {
	val, err := strconv.Atoi("1")
	fmt.Println(val)
	fmt.Println(err)

	val2, err2 := strconv.Atoi("aaaaa")
	fmt.Println(val2)
	fmt.Println(err2)
}
