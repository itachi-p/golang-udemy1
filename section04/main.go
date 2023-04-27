package main

// 型の宣言と初期化、キャスト（型変換）

import "fmt"

func main() {
	/*
		単にintとした場合も内部的には8,16,32,64bitのいずれかになる。
		上記は環境依存で自動的に決定されるが、明示的な指定や型変換も可能。
		同じintでもそれぞれ別の型と見做される為、計算時は型変換が必要。
	*/
	var x int          // 標準のint型。宣言のみ、初期化なし
	var i8 int8 = 8    // 8-bit integer (-128 to 127)
	var i16 int16 = 16 // 16-bit integer (-32768 to 32767)
	var i32 int32 = 32 // 32-bit integer (-2147483648 to 2147483647)
	var i64 int64 = 64 // 64bit型intの明示的指定
	var u8 uint = 255  // 符合無し正の整数 (0 to 255, 0 to 65535, ...)

	// 宣言だけした場合、intの初期値0で初期化される。
	// なお、Golangでは宣言した変数は全て使わないと静的エラーになる。
	fmt.Printf("宣言だけして初期化しなかったintの値: %d\n\n", x)

	// 同じintでも別の型同士ではいずれかへの型変換が必要
	// ※ 型変換の方法は(int)xではなくint(x)
	fmt.Println(x + int(u8)) // intとuintでも型変換必須
	fmt.Println(i8 + int8(i16))
	fmt.Println(int64(i32) + i64)

	// u8 = -10 // 符合無し(unsigned)intに負数は代入できない

	fmt.Printf("\n符合無し整数値:%d\n", u8)

	// 変数の型を調べる書式
	fmt.Printf("\n%T\n", x)
	fmt.Printf("%T\n\n", int32(x))

	// 浮動小数点型
	var fl64 float64 = 3.2

	// 暗黙的な定義の場合、int型と違って自動ではなく固定でfloat64になる。
	fl := 2.4
	fmt.Println(fl64 + fl) // 暗黙でもfloat64なので型変換無しで計算可能
	fmt.Printf("%T, %T\n", fl64, fl)

	// 明示的にfloat32を指定したり、型変換もできるが基本的にあまり使わない
	var fl32 float32 = 1.0
	fmt.Printf("\n%T\n", fl32)
	fmt.Printf("%T\t%f\n", float32(fl64), fl32+float32(fl64))

	// 複素数型もあるが、float32以上にあまり使われない。
	var c128 complex128 = -5 + 12i // complex型
	fmt.Printf("\n%T\t%f", c128, c128)
}
