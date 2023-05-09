package main

import (
	"fmt"
	"sync"
	"time"
)

//sync : 並行処理
//Mutex（相互排他ロック）による同期処理（競合状態の回避）
//（他にも解決策は複数あり、既に学習したchannelを使っても回避可能。）

var st struct{ A, B, C int }

// Mutexを保持するパッケージ変数を設定
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	//排他ロック
	//1つのgoroutineのみロックを取得できる
	//解放されるまで他のgoroutineはロックを取得できず待つ
	//つまり常に1つのgoroutineが一貫性を持って実行できる。
	mutex.Lock()

	//排他制御により、ズレは回避できる(A,B,C全て同じ値が入る)
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)

	//アンロック（必ずセットで必要）
	mutex.Unlock()
}

func main() {
	//定義したMutexをmain関数の先頭で生成
	mutex = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	// 意図的な遅延（並行処理の終了待ち）
	time.Sleep(2 * time.Second)
}
