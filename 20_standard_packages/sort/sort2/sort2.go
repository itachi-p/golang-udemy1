package main

import (
	"fmt"
	"sort"
)

//sort
//SliceとSliceStable

type Entry struct {
	Name  string
	Value int
}

func main() {

	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}
	fmt.Println(el)

	//sort.Slice
	//第1引数に対象のSlice型(の構造体)、第2引数にソート条件式
	//条件式には戻り値として真偽値を返す無名関数で設定(Nameの昇順)
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	fmt.Println("----------")
	fmt.Printf("Name昇順:\n%v\n", el)
	//続いてValueの昇順でもソートする。
	//こちらは前のNameの昇順並び替えは破棄される。
	/* (筈だが、どう変えてもStableと同じになる…仕様変更でもなさげ。謎。)*/
	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	fmt.Printf("Value昇順:\n%v\n", el)
	fmt.Println("----------")

	//SliceStable : 安定ソート
	//通常のソートとの違いは、ソート前のデータの並び順がソート後も保存されるもの
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	fmt.Printf("Name昇順(Stable):\n%v\n", el)
	//続いてValueの昇順でもソート。Stableでは前のName昇順も保持される。
	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	fmt.Printf("Value昇順(Stable):\n%v\n", el)
	fmt.Println("----------")
}
