package main

import "fmt"

//構造体でもGenericsを使ってみる

type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

// 構造体に型パラメータを使った場合、メソッドに新たな型パラメータを持つことはできない
// func (t T[A, B, C])m[U any](x U){} //エラーになる
func (t T[A, B, C]) m(x int) {
	//これなら通る（中身の実装は割愛）
}

func main() {
	var t T[int, []*int, *int]
	fmt.Printf("A: %T, B: %T, C: %T\n", t.a, t.b, t.c)

}
