package main

import "fmt"

// スライス
// append make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 300) // 要素数の追加
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))

	// makeに第3引数を渡すと容量が指定できる
	sl3 := make([]int, 5, 10)
	fmt.Println(sl3)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	// 容量を超える要素が追加されると容量が一気に倍確保される
	// これによりメモリ消費も倍になるが、最初は気にせずとも可
	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sl3)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

}
