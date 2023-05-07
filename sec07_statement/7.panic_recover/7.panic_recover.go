package main

import "fmt"

//パニック＆リカバー
/* 要は例外処理だが、Goのランタイムを強制的に停止させる為、
パニックよりもエラーハンドリングを使う方が推奨されている */

func main() {

	//無名関数でリカバー処理
	defer func() {
		// パニック状態であれば変数xに値が返ってくる
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	//ランタイムを強制的に終了させる
	/* 原則deferとセットで使うが、そもそもあまり出番がない。
	同じ例外処理でもエラーハンドリングの方が推奨される。 */
	panic("Runtime error")
	//sfmt.Println("Start") //到達しない
}
