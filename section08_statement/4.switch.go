package main

import "fmt"

// switch文
// 式switch

func main() {
	n := 3
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default: // どのケースにも当て嵌まらない場合に実行される
		fmt.Println("I don't know.")
	}

	// 下記のように簡易文付きでも書ける。
	/* こちらの方が変数nがswitch文の内部でのみ参照可能
	変数の局所性を高められるので変数の評価上でより望ましい */
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know.")
	}

	n2 := 6
	switch {
	case n2 > 0 && n2 < 4:
		fmt.Println("0 < n < 4")
	case n2 > 3 && n2 < 7:
		fmt.Println("3 < n < 7")
	}
}
