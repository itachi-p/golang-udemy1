package main

import (
	"fmt"
	"math/rand"
	"time"
)

//rand : 乱数生成

func main() {
	//rand.Seed()はGo1.20より非推奨
	//rand.Seed(256)
	rand.NewSource(256)

	//0.0~1.0の間で乱数生成
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	//現在時刻(1970年1月1日からの累積ナノ秒値)をシード値に設定
	fmt.Println(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())

	//0~99(n-1)の間の乱数を生成
	fmt.Println(rand.Intn(100)) //nは生の整数に限る
	//int型の乱数(64bit環境ならrand.Int63()と同等の結果)
	fmt.Println(rand.Int())
	//int32型の乱数
	fmt.Println(rand.Int31())
	//int64型の乱数
	fmt.Println(rand.Int63())
	//uint32型の乱数(負の乱数が生成されない以上rand.int31()と同様？)
	fmt.Println(rand.Uint32())
}
