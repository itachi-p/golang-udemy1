package main

import "fmt"

//演算子

func main() {
	fmt.Println("Fizz" + "Buzz")

	//算術演算子と代入
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

	//比較演算子
	i1 := 3
	i2 := 5

	fmt.Println(i1 == i2)
	fmt.Println(i1 > i2)
	fmt.Println(i1 <= i2)
	fmt.Println(i1 != i2)

	//論理演算子
	s1 := "ABC"
	s2 := "AB"
	s3 := "abc"

	fmt.Println(s1 == s2 && s3 != "")
	fmt.Println(s1 == s2 || s3 != "")
	fmt.Println(s1 != s2 && s1 != s3)
	fmt.Println(s1 == s2 || s1 == s3)
}
