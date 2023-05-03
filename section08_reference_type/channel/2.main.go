package main

import (
	"fmt"
	"time"
)

//channel and goroutine
//チャネルは複数のgoroutine間でのデータ送受信用データ構造
//そもそも単一の関数の中だけで使う意味は無い

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//並行で2つの無限ループを走らせる
	go reciever(ch1)
	go reciever(ch2)

	//チャネルにデータを送信する
	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		//上記は2行とも瞬時に行われる。
		//実際は以下のように待ちを挟まないと高確率でランダムにズレる
		time.Sleep(50 * time.Millisecond)
		//ハイスペックな実行環境なら以下でもズレないが、環境依存
		//time.Sleep(1 * time.Millisecond)

		i++
	}
}
