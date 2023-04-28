package main

import "fmt"

// 文字列型

func main() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "256"
	fmt.Println(si + "512")

	// 複数行文字列
	fmt.Println(`test
	test
		test`)

	// エスケープ文字
	fmt.Println("\"")
	fmt.Println('"')

	// string型の中から特定の位置の文字を(ASCIIコード番号のbyte型で)表示
	fmt.Println(s[0])
	// byte型のASCIIコード番号に対応する文字として表示
	fmt.Println("s[0]")
}
