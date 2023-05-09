package main

import (
	"fmt"
	"regexp"
)

// regular_expression : 正規表現
func main() {
	//Goの正規表現の基本 regexp.MatchString
	//簡易版。同じ正規表現チェックを繰り返すのには向いていない
	match, _ := regexp.MatchString("BC", "ABCD")
	fmt.Println(match)

	//Compile 頻繁に検索条件を変える場合に有用
	re1, _ := regexp.Compile("AB")
	match = re1.MatchString("AbC")
	fmt.Println(match)

	//MustCompile 基本はこれを使う。
	//2連続エスケープを避ける為、バッククォート(`)の使用推奨
	//regexp.MustCompile("\\d")
	re2 := regexp.MustCompile(`\d`) // \d は任意の半角数字[0-9]
	match = re2.MatchString("ABc９")
	fmt.Println(match)

	// \w はアンダーバーを含む半角英数字 [_0-9a-zA-Z]
	re3 := regexp.MustCompile(`\w`)
	match = re3.MatchString("*+-'&%$#!あＧ９ｊ")
	fmt.Println(match)

	/* 正規表現による文字列の置換
	正規表現にマッチした部分を別の文字列に置き換える。
	regexp.Regexp型メソッドReplaceAllString
	対象の文字列に正規表現のパターンにマッチする部分がない場合は、
	元の文字列がそのまま返される。*/

	// \s:スペース文字[ \t\n\r\f] +:直前の正規表現が1個以上
	re4 := regexp.MustCompile(`\s+`)
	//1個以上の連続した半角スペース(やタブ・改行等の空白文字)を全て,に置き換える
	fmt.Println(re4.ReplaceAllString("I am    sam   ", ","))

	//全角文字も置換可能
	//置換文字列に空文字を指定することで、正規表現にマッチした部分を文字列から除去
	re4 = regexp.MustCompile(`、| |-|　|。`)
	fmt.Println(re4.ReplaceAllString("私　は、Go-langの  プログラマー。", ""))

	//正規表現による文字列の分割
	re5 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	//正規表現で文字列を分割する場合、regexp.Regexp型のメソッドSplitを使う
	//第2引数:分割する最大数。-1でマッチした全ての箇所で分割し、[]stringで返す。
	fmt.Println(re5.Split("ASHVJV<HabcXYZKNJBJVKABCxyz", -1))

	re5 = regexp.MustCompile(`\s+`)
	//半角スペース、タブ、改行などの空白で分割(半角空白1個挟んで表示)する。
	fmt.Println(re5.Split("aaa     bbbb	cccc　dd", -1))
	//一見わかりにくいが、全角スペースを挟んだ部分は分割されず一塊になっている。
	fmt.Println(re5.Split("aaa     bbbb	cccc　dd", -1)[2])
}
