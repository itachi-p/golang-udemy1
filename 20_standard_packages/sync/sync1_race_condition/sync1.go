package main

import (
	"fmt"
	"time"
)

//sync : 並行処理
//Mutex（相互排他ロック）による同期処理（で解決する前の、競合発生状態）

/* レースコンディション【Race Condition/競合状態】
並列動作する複数のプロセスやスレッドが同じリソースに
ほぼ同時にアクセスした際に、想定外の処理結果が生じること。*/

// パッケージ変数を設定
var st struct{ A, B, C int }

// プログラム的には3つの変数全てに同じ値が入る筈だが、実際はランダムにズレる
func UpdateAndPrint(n int) {
	st.A = n
	//確実にズレを生じさせる為に意図的にラグを発生させる
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)
}

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	// 意図的な無限ループ
	// for {
	// これがないと並行処理が終わらないうちにmain関数が先に終了する
	// （他にもっといいやり方がありそうな気もする。）
	// }

	time.Sleep(1 * time.Second)
}
