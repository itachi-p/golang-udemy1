package main

import (
	"fmt"
	"time"
)

//標準パッケージ
//time

func main() {
	//時間の生成
	t := time.Now()
	fmt.Println(t)

	//指定した時間を生成
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

}
