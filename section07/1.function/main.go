package main

import "fmt"

// 関数（超基礎）

// 引数の型が同じ場合は省略して1つにまとめることができる
func Plus(x, y int) int {
	return x + y
}

// 戻り値が複数ある場合
func Div(x, y int) (int, int) {
	q := x / y // 商：quotient
	r := x % y // 剰余：remainder
	return q, r
}

// 戻り値を表す変数
func Double(price int) (result int) {
	result = price * 2
	return // これだけでreturnの値が戻り値として使われる
}

func main() {
	i := Plus(3, 5)
	fmt.Println(i)

	// 関数の返り値が複数の場合、一度に複数の変数に代入もできる
	i2, i3 := Div(11, 4)
	fmt.Printf("商: %d 剰余: %d\n", i2, i3)

	// 複数の戻り値の一部を破棄したい場合
	i2, _ = Div(9, 4) // 変数で受け取った場合は必ず使う必要がある
	fmt.Printf("商: %d\n", i2)

	i4 := Double(1980)
	fmt.Printf("%d円\n", i4)

}
