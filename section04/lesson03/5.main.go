package main

import "fmt"

// byte(uint8)型
// C言語のchar型と同じくASCIIコード表に対応した1文字を正の整数の番号で表す

func main() {
	// 配列とスライス（後のlesson「配列とスライス」で詳しく）
	// byte(uint8)型の配列として文字列(=string)を定義できる
	byteA := []byte{72, 73 + 32, 33} // アルファベット小文字は大文字+32
	// 文字列として表示する時は型変換が必要
	fmt.Println(byteA)
	fmt.Println(string(byteA))

	// 文字列をbyte型のスライスに変換
	c := []byte("Hi!")
	fmt.Println(c)
	// byte型のスライスを文字列に変換
	fmt.Println(string(c))
}
