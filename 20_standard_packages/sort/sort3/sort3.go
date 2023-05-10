package main

import (
	"fmt"
	"sort"
)

//sort
//CustomSort

type Entry struct {
	Name  string
	Value int
}
type List []Entry

// ---
// type Interface interface {} のカスタマイズ
// Len() int, Less(i, j int) bool, Swap(i, j int)の3つの関数を持つ
func (l List) Len() int {
	return len(l)
}
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// ここをカスタマイズする
func (l List) Less(i, j int) bool {
	//Valueが同じであればNameの昇順になる
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

//---

func main() {
	m := map[string]int{"S": 1, "J": 4, "A": 3, "N": 3}

	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	//Sort
	sort.Sort(lt)
	fmt.Println(lt)

	//Reverse
	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)
}
