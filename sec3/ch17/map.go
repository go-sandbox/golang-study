package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)
	fmt.Println(m["nothinf"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["hoge"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// nilのためエラーとなる
	// // var m3 map[string]int
	// // m3["pc"] = 5000
	// fmt.Println(m3)
}
