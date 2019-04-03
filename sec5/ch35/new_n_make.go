package main

import "fmt"

func main() {
	// メモリ領域をnewで確保
	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	// メモリ領域確保していないのでnil
	var p2 *int
	fmt.Println(p2)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

}
