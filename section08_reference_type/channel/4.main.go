package main

import "fmt"

// channel : select
//複数チャネルに対する送受信をルーチンを停止させることなく制御する

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)
	/*
		ch2 <- "A"
		ch1 <- 1
		ch2 <- "B"
		ch1 <- 2
	*/
	/*
		//v1にまだ値が入ってないので、v2に到達できずdeadlockになる
		v1 := <-ch1
		v2 := <-ch2
		fmt.Println(v1)
		fmt.Println(v2)
	*/

	select {
	//caseの式の評価は上から順ではなく毎回ランダム
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("None\n")
	}

	//selectの活用例
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	//reciever(無名関数で)
	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()
	//reciever2
	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	//チャネルにデータを渡す部分
	n := 0
Lavel:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			//まずch3にnの値が渡り、ch4,ch5と渡され最後にi3として表示
			fmt.Println("received", i3) //最初の値はn=0×2-1=-1
		default:
			if n > 100 {
				break Lavel
			}
		}
		//これをselect文のdefault式に入れることもできる
		//その場合単純にbreakだと無限ループになるのでラベルを付ける
		/*
			if n > 100 {
				break
			}
		*/
	}
}
