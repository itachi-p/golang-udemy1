package main

import "fmt"

// init
/* パッケージの初期化を目的とした特殊な関数
mainより先に実行される。
main関数の中で呼び出す必要はない。 */

// main関数実行前に確実に初期化を行う
func init() {
	fmt.Println("init")
}

// 複数作成も可能。順番に実行されるが、あまり分けるメリットはない。
func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("main")
}
