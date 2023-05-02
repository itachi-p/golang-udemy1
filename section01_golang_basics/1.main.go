package main

import (
	"fmt"
	"time"
)

// HelloWorld
/*
複数行のコメント
*/

func main() {
	fmt.Println("Hello World!")
	fmt.Println(time.Now())
	fmt.Printf("My Name is %v.\nI was born in %d.\tI'm %d years old.", "Gopher", 2009, time.Now().Year()-2009)
}
