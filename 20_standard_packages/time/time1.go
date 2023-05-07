package main

import (
	"fmt"
	"time"
)

//標準パッケージ
//time

func main() {
	//現在時刻
	t := time.Now()
	fmt.Println(t)

	//指定した時間の生成
	t2 := time.Date(2024, 12, 31, 12, 34, 56, 99, time.Local)
	fmt.Println(t2.Year())
	fmt.Println("(YearDay:)", t2.YearDay()) //閏年の3/1以降は+1日
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())
	fmt.Println(t2.Date()) //year,dayのint値とmonthの合成値
	fmt.Println("Hour:", t2.Hour())
	fmt.Println("Minute:", t2.Minute())
	fmt.Println("Second:", t2.Second())
	fmt.Println("Nanosecond:", t2.Nanosecond())
	fmt.Println(t2.Zone()) //stringとintの合成値

	//time.Duration型を返す
	fmt.Printf("\n%T\n", time.Hour)
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	//time.Duration型を文字列から生成
	d, _ := time.ParseDuration("2h30m")
	fmt.Println("\n", d)

	//現在時刻の2分30秒後を表すtime.Time型を取得
	t3 := time.Now()
	fmt.Println(t3)
	t3 = t3.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t3)
	fmt.Printf("%T\n\n", t3)

	//時刻の差分を取得
	t4 := time.Date(2023, 5, 24, 0, 0, 0, 0, time.Local)
	t5 := time.Now()
	fmt.Println(t5)

	//t4 - t5
	d2 := t4.Sub(t5)
	fmt.Println("時刻の差分:", d2)

	//時刻を前後関係で比較判定する
	fmt.Println(t5.Before(t4))
	fmt.Println(t5.After(t4))
	fmt.Println(t4.Before(t5))
	fmt.Println(t4.After(t5))

	//指定時間のスリープ
	//実行しているgoroutineを5秒間停止
	time.Sleep(5 * time.Second)
	fmt.Println("\n5秒停止後表示")

}
