package main

import "fmt"

// interface型
// fmt.Stringer型インターフェース
//文字列を返すString()メソッドだけが定義されている
/*
type Stringer interface {
	String() string
}
*/

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	//通常は&{100 ABC}と表示されるところを、指定のフォーマットに変えられる
	fmt.Println(p)
}
