package main

import "fmt"

// 算術演算子

func main() {
	fmt.Println("Fizz" + "Buzz")

	// 算術演算子と代入
	n := 0
	n += 4
	fmt.Println(n)
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)

	s := "ABC"
	s += "DEF"
	fmt.Println(s)
}
