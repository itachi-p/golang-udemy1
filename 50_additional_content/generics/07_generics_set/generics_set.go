package main

import "fmt"

// 型パラメータを持つ型set:KeyとValueを持つMap型のような構造
// type comparable:比較が可能な要素を持つ、という制約を持たせる
// mapのValueではなく、Keyが比較可能な値である必要がある
// comparableインタフェースは型パラメータの制約としてのみ使用でき、変数の型としては使用できない。
type Set[T comparable] map[T]struct{}

// コンストラクタ関数
func NewSet[T comparable](xs ...T) Set[T] {
	//新規で中身が空の入れ物を生成
	s := make(Set[T])
	//要素を追加していく
	for _, v := range xs {
		//s[v] = struct{}{} //それぞれのKeyに対し空のValueが入る
		s.Add(v)
	}
	return s
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Remove(x T) {
	//mapで使えるdeleteを使って要素を削除できる
	delete(s, x)
}

func main() {
	s := NewSet[int](1, 2, 3) //型推論できるので[int]は省略可
	fmt.Println(s)

	s.Remove(1)
	fmt.Println(s)
}
