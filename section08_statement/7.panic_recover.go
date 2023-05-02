package main

import "fmt"

// パニック＆リカバー
/* 要は例外処理だが、Goのランタイムを強制的に停止させる為、
パニックよりもエラーハンドリングを使う方が推奨されている */

func main() {
	panic("Runtime error")
	fmt.Println("Start")
}
