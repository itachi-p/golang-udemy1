package main

import (
	"fmt"
	"strings"
)

//strings

func main() {
	//文字列を結合
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"N", "o", "A"}, "")
	fmt.Println(s1, s2)

	//文字列に含まれる部分文字列の最初の位置を検索
	i1 := strings.Index("ABCDA", "A")
	i2 := strings.Index("ABCDA", "D")
	i3 := strings.Index("ABCDA", "Z") //存在しない場合-1
	fmt.Println(i1, i2, i3)           //Index番号0から数えてn番目で発見

	i4 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i4)

	//第2引数の文字列のいずれかが最初・最後に現れるインデックス番号
	i5 := strings.IndexAny("XABCY", "ABC")
	i6 := strings.LastIndexAny("XABCY", "ABC")
	fmt.Println(i5, i6)

	//指定した文字で始まるか、終わるか
	b1 := strings.HasPrefix("ABC", "A")
	b2 := strings.HasSuffix("ABCZ", "Z")
	fmt.Println(b1, b2)

	//指定の文字・文字列のいずれかが含まれるか
	b3 := strings.Contains("ABC", "D")
	b4 := strings.ContainsAny("ABCDE", "XJCM")
	fmt.Println(b3, b4)

	//指定の文字が何回出現するか
	i7 := strings.Count("ABCABCABC", "B")
	i8 := strings.Count("ABCABC", "") //空文字を指定した場合、文字数+1が返る
	fmt.Println(i7, i8)

	//文字列を繰り返し結合
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("DEF", 0) //負数はランタイムエラー発生
	fmt.Println(s3, s4)

	//文字列の置換
	//第3引数は置換する回数 -1で全置換
	s5 := strings.Replace("AAAAAA", "A", "B", 3)
	s6 := strings.Replace("AAAAAA", "A", "B", -1)
	fmt.Println(s5, s6)

	//文字列分割(文字列型のsliceとして返る)
	s7 := strings.Split("A,B,C,D,E", ",")
	fmt.Printf("%s, %T\n", s7, s7)

	//大文字・小文字の置換
	s13 := strings.ToUpper("abc")
	fmt.Println(s13)

	//文字列の「前後」から空白(半角及び全角)/指定の文字を除去
	//取り除かれるのは前後だけである点に注意
	s15 := strings.TrimSpace("  　　- H o　g e F u 　g a  -　 ")
	s16 := strings.Trim("///*F/i/z/z//_/B/u//z//z*////", "/")
	fmt.Println(s15, s16)

	//文字列からスペースで区切られたフィールドを取り出しslice型で返す
	s18 := strings.Fields("a b    c  ")
	fmt.Printf("%s, %T\n", s18, s18)
	fmt.Println(s18[0], s18[1])
}
