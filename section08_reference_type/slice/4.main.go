package main

import "fmt"

// slice : for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	for _, v := range sl {
		fmt.Println(v)
	}

	// 古典的forでも可
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}
