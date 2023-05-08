package main

import "fmt"

// スライス
// copy
/* slice, map, channelはチャンネルは全て参照型の為、参照渡しになり
通常の変数と同じようにコピーしても同じアドレスを指してしまう
（C言語のポインタ参照と同じ) */

func main() {
	// 基本型なら以下でコピー元とは別の変数が作れる
	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i, i2)

	sl := []int{100, 200}
	// 単純に代入すると同じアドレスを参照する為、コピー元の値も変わる
	sl2 := sl

	sl2[0] = 1000
	fmt.Println(sl, sl2)

	// 参照型であるsliceのコピーを作る場合はcopy関数を使う
	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 10)
	fmt.Println(sl4)

	// copy関数はコピー先の配列の要素の前から順に埋めていく
	n := copy(sl4, sl3) // 戻り値 = コピーに成功した回数
	fmt.Println(n, sl4)

}
