package main

import "fmt"

// channel
// 複数のgoroutine間でのデータの受け渡しの為に設計されたデータ構造

func main() {
	// チャネルの宣言（特に指定しなければ送受信の双方向）
	var ch1 chan int

	/*
		//受信専用
		var ch2 <-chan int

		//送信専用
		var ch3 chan<- int
	*/

	//初期化
	ch1 = make(chan int)

	//宣言と初期化を同時に
	ch2 := make(chan int)

	// 容量（バッファサイズ）確認
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5) //バッファサイズを指定
	fmt.Println(cap(ch3))

	//チャネルにデータを送信
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	fmt.Println("len:", (len(ch3)))

	//チャネルからデータを受信
	i := <-ch3
	i2 := <-ch3
	fmt.Println(i, i2)
	fmt.Println("len:", (len(ch3)))

	//変数に入れずに取り出し
	fmt.Println(<-ch3)
	fmt.Println("len:", (len(ch3)))

	/*上記の実行結果から、チャネルはqueの性質
	（先入れ・先出しのデータ構造）を持つことがわかる*/

	// バッファサイズを超えたデータを送信した場合
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	//ch3 <- 6 //バッファを超えた時点でdeadlockエラー発生

	//deadlock回避
	fmt.Println(<-ch3) //これでサイズが1つ減る
	fmt.Println("len:", (len(ch3)))
	ch3 <- 6 //上で1つ減ったので容量を超えずに実行される
}
