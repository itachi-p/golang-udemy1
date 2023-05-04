package main

import "fmt"

//ポインタ
/* 値型に分類されるデータ構造(基本型、参照型、構造体）のメモリ上のアドレスと型の情報
C言語と同じく、データ構造を間接的に参照、操作ができる様になっている。
主に値型の操作に使われる。参照型は元からその機能を含んでいる為、基本的には不要 */

func Double(i int) {
	i = i * 2
}

// 上記をポインタ型の参照渡しに書き換える
func Double_v2(i *int) {
	*i = *i * 2
}

// スライスやマップ、チャネルなどは元々参照渡しの機能を持っている
func Double_v3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)

	//変数のメモリ上のアドレス表示
	fmt.Println(&n)

	//値n（のコピー）を渡しても関数内部でのみ倍になり、元のnは変わらない
	Double(n)
	fmt.Println(n) // 関数に渡す前のnがそのまま表示される

	//上記を解決するのがポインタ型（参照渡し）
	var p *int = &n //ポインタ型変数にnのアドレスを渡している

	fmt.Println(p)  //アドレス表示
	fmt.Println(*p) //実体

	//ポインタの実体の値を操作すると元の変数の値も変わる
	/*
		*p = 300
		fmt.Println(n)
		//逆も同じ
		n = 200
		fmt.Println(*p)
	*/

	Double_v2(&n)
	fmt.Println(n)

	Double_v2(p)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3}
	Double_v3(sl)
	fmt.Println(sl)

}
