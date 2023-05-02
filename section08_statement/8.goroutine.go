package main

import (
	"fmt"
	"time"
)

// 並行処理(goroutine)
// 本来難易度が高い処理が他の言語より簡単に実装できる

func sub() {
	// 通常は前の無限ループが終わるまで次の処理には進めない
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub() // 呼び出し関数の前に'go'と付けるだけ
	go sub() // 3つの処理を同時に並行処理

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
