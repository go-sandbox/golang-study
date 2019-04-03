package main

import "fmt"

type Vertex struct {
	X int
	Y int // 小文字で宣言すると、プライベートになり、外部からアクセスできなくなる
	S string
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100

	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "ge"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)

	var v5 Vertex
	fmt.Println(v5)

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)
}
