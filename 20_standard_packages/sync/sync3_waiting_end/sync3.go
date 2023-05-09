package main

import (
	"fmt"
	"sync"
)

//sync3
//goroutineの終了を待ち受ける

func main() {
	//sync.WaitGroupを生成
	wg := new(sync.WaitGroup)
	//待ち受けるGoroutineの数は3個
	wg.Add(3)

	/*意図的に無限ループや遅延を付けないと、これらの3つの非同期処理が
	何も実行されないうちにmain関数が終了してしまう*/
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done() //完了
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2st Goroutine")
		}
		wg.Done() //完了
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3st Goroutine")
		}
		wg.Done() //完了
	}()

	//Goroutineの完了をここで待ち受ける
	//Doneが3つとも完了すると実行される
	wg.Wait()

	//time.Sleep(1 * time.Second)
}
