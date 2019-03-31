package main

import "fmt"

func main() {
	// Q1
	f := 1.11
	fmt.Println(int(f))

	// Q2 => 5,6

	// Q3
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)

}
