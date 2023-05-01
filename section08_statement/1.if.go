package main

import "fmt"

// if構文:条件分岐

func main() {
	a, b := 1, 2
	if a == 1 && b == 1 {
		fmt.Println("one")
	} else if a == 2 || b == 2 {
		fmt.Println("two")
	} else if a == 1 && b == 2 { // 上から順に評価される為、else構文ではここには入らない
		fmt.Println("one & two")
	} else {
		fmt.Println("I don't know.")
	}
}
