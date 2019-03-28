package main

import "fmt"

func main() {
	var t, f bool = true, false
	fmt.Printf("%T %v\n", t, t)
	fmt.Printf("%T %v\n", f, f)

	fmt.Println(true && true)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(false || false)
	fmt.Println(false || true)

	fmt.Println(!true)
	fmt.Println(!false)
}
