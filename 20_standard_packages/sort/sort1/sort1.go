package main

import (
	"fmt"
	"sort"
)

//sort

func main() {

	i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10, 1, 0}
	s := []string{"a", "z", "h", "n", "j", "b", "A", "X"}

	fmt.Println(i, s)

	sort.Ints(i)
	sort.Strings(s)

	fmt.Println(i, s)
}
