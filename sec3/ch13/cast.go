package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)

	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"

	// GoではPython等とはことなり、型変換できない
	// z = int(s)

	// Stringからint変換する場合は、この方法を用いる ASCI To Integer
	i, _ := strconv.Atoi(s)

	fmt.Println(i)

	h := "Hello World"
	fmt.Println(string(h[0]))
}
