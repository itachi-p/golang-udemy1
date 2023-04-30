package main

import "fmt"

// 論理演算子

func main() {

	s1 := "ABC"
	s2 := "AB"
	s3 := "abc"

	fmt.Println(s1 == s2 && s3 != "")
	fmt.Println(s1 == s2 || s3 != "")
	fmt.Println(s1 != s2 && s1 != s3)
	fmt.Println(s1 == s2 || s1 == s3)
}
