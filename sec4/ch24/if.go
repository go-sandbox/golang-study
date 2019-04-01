package main

import "fmt"

func main() {
	num := 9
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("hoge")
	}

	x, y := 10, 10

	if x == 10 && y == 10 {
		fmt.Println("&&")
	}
}
