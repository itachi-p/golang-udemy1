package main

import (
	"fmt"
	"os"
)

//fmt

func main() {
	//fmt標準
	fmt.Print("vim-go!")
	//改行
	fmt.Println("A", "B", "C")
	//フォーマット指定
	//書式指定子 %#v: 値の文法をそのまま出力する
	fmt.Printf("%#v\n", "Hoge")

	//Sprint系
	//出力ではなく、フォーマットした結果を文字列で返す。
	s1 := fmt.Sprintf("%v", "Hello")
	s2 := fmt.Sprintln("Hey!")

	fmt.Println(s1)
	fmt.Println(s2) //Sprintln+Printlnで改行が二重に入る点に注意

	//Fprint系 書き込み先指定
	fmt.Fprintln(os.Stdout, "Hiho!")

	f, _ := os.Create("Fprint.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint test.")
}
