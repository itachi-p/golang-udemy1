package main

import (
	"fmt"
	"time"
)

// channel : close
// チャネルを明示的に閉じる
//閉じたチャネルに対してデータ送信はできなくなる(ランタイムエラー)
//ただしクローズしたチャネルからのデータ受信は可能

func reciever2(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + ": END")
}

func main() {
	ch1 := make(chan int, 2)

	/*
		// panic: send on closed channel発生
		ch1 <- 1

		//チャネルを明示的に閉じる
		close(ch1)

		//閉じたチャネルへの送信はpanic: send on closed channel発生
		//ch1 <- 1

		//閉じたチャネルからも受信は可能
		//fmt.Println(<-ch1)

		//チャネルからの受信は2つの引数を受け取ることができる
		//2つ目の引数の真偽値はチャネルの状態を示す
		//チャネルの状態がcloseかつバッファが0の時falseになる
		i, ok := <-ch1
		fmt.Println(i, ok)
		//上でバッファが0になったので次はfalseになる
		i, ok = <-ch1
		fmt.Println(i, ok)
	*/

	go reciever2("1.goroutine", ch1)
	go reciever2("2.goroutine", ch1)
	go reciever2("3.goroutine", ch1)

	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
	//チャネルを閉じた後3つのgoroutineの終了を簡易的に3秒待つ
	//本来は全goroutineの終了を正確に検出するが、詳細は後述
	time.Sleep(3 * time.Second)
}
