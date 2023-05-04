package main

import "fmt"

// interface型
// 最もポピュラーな使い方。異なる型に共通の性質を付与する。

// 異なる型に共通するインターフェースの定義
type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

// 異なる型に同名のメソッドを用意する
func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {
	//共通のStringfyという型でスライスとしてまとめることができる
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}
	fmt.Println(vs)
	for _, v := range vs {
		//異なるデータ型に共通のメソッドを使える
		fmt.Println(v.ToString())
	}
}
