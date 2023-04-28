package main

import "fmt"

// 文字列型

func main() {
	var s string = "Hello Golang!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "256"
	// 文字列同士の連結
	fmt.Println(si + "abc")

	// 複数行文字列
	fmt.Println(`test
	test
		test`)

	// エスケープ文字
	fmt.Println("\"")
	fmt.Println(`"`) // シングルクォートではなくバッククォート
	// 文字(または特殊文字)のASCIIコード番号を表示(byte型として表示)
	fmt.Println('\n')

	// string型の中から特定の位置の文字を(byte型のASCIIコード番号で)表示
	fmt.Println(s[0])
	// 更にそれをbyte型のASCIIコード番号に対応する文字として表示
	fmt.Println(string(s[0])) // 大文字の'A'のASCIIコード65なので'H'は72
	/* つまり上記から「string型 = byte型の配列」であるとわかる。 */
}
