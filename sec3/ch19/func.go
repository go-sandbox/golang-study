package main

import "fmt"

func add(x, y int) (int, int) {
	fmt.Println("add func")
	fmt.Println(x + y)
	return x + y, x - y
}

func cal(price, item int) (r_result int) {
	r_result = price * item
	return
}

func main() {
	result, result2 := add(10, 20)
	fmt.Println(result, result2)

	r3 := cal(100, 2)
	fmt.Println(r3)
}
