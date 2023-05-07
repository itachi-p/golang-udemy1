package main

import "fmt"

// slice 1
// 宣言、操作

func main() {
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	// 暗黙的な宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	// make関数 指定した要素数そのデータ型の初期値で初期化される
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	// 値の代入
	sl2[0] = 300
	fmt.Println(sl2)

	// 値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)
	fmt.Println(sl5[0])          // 最初の要素のindexは1でなく0
	fmt.Println(sl5[2:4])        //2番目から4番目の手前まで取り出す
	fmt.Println(sl5[:2])         // 2番目の手前までを取り出す
	fmt.Println(sl5[2:])         //2番目以降最後まで取り出す
	fmt.Println(sl5[len(sl5)-1]) //配列の最後の要のindexはlen-1

}
