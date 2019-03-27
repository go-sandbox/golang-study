package main

import "fmt"

// var 宣言は、関数外でも宣言できる
var (
	i   int     = 1
	f64 float64 = 1.2
	s   string  = "test"
	t   bool    = true
	f   bool    = false
)

func main() {

	fmt.Println(i, f64, s, t, f)

	// 型推論かつvarが不要
	// ただし、関数外では宣言できない
	// floatとか型をしっかりと指定したい場合はvarを使って型を明示的に宣言する必要あり
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xt)
}
