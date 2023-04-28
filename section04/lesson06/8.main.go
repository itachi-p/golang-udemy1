package main

import "fmt"

// 型変換

func main() {
	var i int = 1
	fl64 := float64(i) // intからfloatへの変換は容易にできる。
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

}
