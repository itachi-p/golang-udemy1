package main

import "fmt"

// if構文:条件分岐

func main() {
	a, b := 1, 2
	if a == 1 && b == 1 {
		fmt.Println("one")
	} else if a == 2 || b == 2 {
		fmt.Println("two")
	} else if a == 1 && b == 2 { // 上から順に評価される為、else構文ではここには入らない
		fmt.Println("one & two")
	} else {
		fmt.Println("I don't know.")
	}

	// 簡易文付きif構文
	/* 簡易文とは変数の定義と代入などの「構造を持たない」文。
	関数の戻り値を変数で受け取り、続いてその変数を評価して条件分岐もできる。
	エラーハンドリングが必要な処理で活躍する。 */
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	/* 注意点として、簡易文の変数のスコープはif文の中だけ。
	簡易文で定義した変数にはif文の外からはアクセスできず、
	簡易文で同じ変数名で別の値を設定しても、if文の外の変数には影響されない。 */
	x := 0
	if x := 2; x == 2 {
		fmt.Printf("簡易文付きif文内のx: %d\n", x)
	}
	fmt.Printf("if文の外のx: %d\n", x) // 名前は同じでも別物
	// ややこしくなるので、上記のような同名変数の使用は避けた方がよさげ。

	// ただし、既に初期化済みの変数の値の変更は可能。（":"だけの違い）
	y := 0
	if y = 2; y == 2 {
		fmt.Printf("簡易文付きif文内のy: %d\n", y)
	}
	fmt.Printf("if文の外のy: %d\n", y) // yのスコープはif文内も含む全体

}
