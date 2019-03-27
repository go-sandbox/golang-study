package main

import "fmt"

// 通常は、グローバルに宣言
// 普遍的な定数を宣言
const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

// var big int = 9223372036854775807 + 1

// const ではオーバーフローしない
// なぜならconst では型は呼ばれた時に指定されるから（型推論）
//　数値でいけそうなら数値型とか
const big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
}
