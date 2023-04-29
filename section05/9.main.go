package main

import "fmt"

// 定数
// 関数内にも書けるが、基本的に定数はグローバル変数として定義される。
// グローバル定数かつ頭文字を大文字にすることで、他のpackageからも呼び出せる。

const Pi = 3.14159

// 複数の定数を一度に定義
const (
	URL      = "https://itachi-p.co.jp"
	SiteName = "test"
)

// 値の省略
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// 変数では型int64の最大値を超えたらオーバーフローエラー
//var Big int = 9223372036854775807 + 1

// 定数に最大値の制限は無いので巨大な数値でも定義じたいは可能。
const Big = 9223372036854775807 + 1

/* 連番の自動生成 */
const (
	C0 = iota
	C1
	C2
)

func main() {
	fmt.Println(Pi)

	// 定数は上書きできない
	//Pi = 3

	fmt.Printf("URL %s/%s\n", URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	// 定義の段階ではエラーにならないが、表示の段階では最大値を超えられない
	fmt.Println(Big - 1)

	fmt.Println(C0, C1, C2)
}
