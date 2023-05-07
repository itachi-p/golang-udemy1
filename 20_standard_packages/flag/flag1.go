package main

import (
	"flag"
	"fmt"
)

//flag

func main() {
	//コマンドライン引数・オプション処理
	//例) $>go run main.go -n 20 -m message -x

	//オプションの値を格納する変数の定義
	var (
		max int
		msg string
		x   bool
	)

	//引数の3番目はデフォルト値
	//4番目は未定義オプションが渡されたエラー発生時の説明文
	//IntVar 整数のオプション
	flag.IntVar(&max, "n", 32, "処理数の最大値")
	//StringVar 文字列のオプション
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	//BoolVar bool型のオプション コマンドラインに与えられればtrue
	flag.BoolVar(&x, "x", false, "拡張オプション")

	//コマンドラインをパース
	flag.Parse()

	fmt.Println("最大処理数:", max)
	fmt.Println("処理メッセージ:", msg)
	fmt.Println("拡張オプション:", x)

}
