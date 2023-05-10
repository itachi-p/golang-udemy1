package main

import (
	"context"
	"fmt"
	"time"
)

//context
//APIのサーバーやクライアント使用時のキャンセルやタイムアウトの仕組み
//リクエストを送って一定時間待っても処理結果が返らない場合などに使われる

// 処理時間のかかる関数を用意
func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("開始")
	//2秒間待つ（これがこの関数のメイン）
	time.Sleep(2 * time.Second)
	fmt.Println("終了")
	ch <- "実行結果"
}

func main() {
	//channelを生成
	ch := make(chan string)
	//contextを生成
	ctx := context.Background()
	//contextに1秒間のタイムアウト設定を付ける
	//2秒未満の設定なら処理が終わらず、エラーでmainのgoroutineが終了される
	//3秒以上に設定すれば先に処理が終わってからmainも終了する
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	//リソースの解放処理。確実に実行されるようにdefer
	defer cancel()

	go longProcess(ctx, ch)
L:
	for {
		select {
		//context.WithTimeoutの時間内に処理が終わらなければキャンセルが実行される
		case <-ctx.Done():
			fmt.Println("########Error########")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}
	}
	fmt.Println("ループ抜けた")
}
