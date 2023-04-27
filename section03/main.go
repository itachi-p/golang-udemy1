package main

// 変数

import (
	"fmt"
	"time"
)

func main() {
	var nowMonth = time.Now().Month()
	var today int = time.Now().Day()
	fmt.Printf("今日は%d月%d日です。", nowMonth, today)
}
