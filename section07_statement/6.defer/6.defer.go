package main

import (
	"fmt"
	"os"
)

//defer
//関数の終了時に実行される処理を登録できる

func TestDefer() {
	/* 通常はプログラムは上から順に実行されるが、
	deferを設定すると関数の終了時に1回だけ実行される。 */
	defer fmt.Println("End")
	fmt.Println("Start")
}

// defer文を複数登録
func RunDefer() {
	//複数登録すると逆順に実行される点に注意
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()

	//deferで複数の処理を記述したい場合は無名関数を使う
	defer func() {
		fmt.Println("F")
		fmt.Println("i")
		fmt.Println("n")
	}()

	RunDefer()

	// defer文を使ったリソースの解放処理
	file, err := os.Create("test.txt") //osパッケージ詳細は後述
	if err != nil {
		fmt.Println(err)
	}
	//ファイルを開いたら最後に閉じる必要がある
	defer file.Close()
	//ファイルを作成し、文字列を書き込む処理
	file.Write([]byte("Hello"))
}
