package main

import "fmt"

// map : key-value形式の配列

func main() {
	// 明示的宣言
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	// 暗黙的宣言
	m2 := map[string]int{"C": 300, "D": 400}
	fmt.Println(m2)

	// 改行して見やすく宣言
	m3 := map[int]string{
		1: "fugafuga",
		2: "hogehoge",
	}
	fmt.Println(m3)

	// make関数を使って一旦空のmapを生成
	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "Japan"
	m4[2] = "USA"
	m4[3] = "Panama"
	fmt.Println(m4)

	// 値の取り出し
	fmt.Println(m["A"])
	fmt.Println(m4[2])
	// goではインデックスが存在しなくても型の初期値(stringなら空文字)が入る
	// エラー終了にはならないが、逆にミスにも気付きにくいので要注意
	fmt.Printf("存在しないインデックスの値:%s←\n", m4[len(m4)+1])

	// mapの値を取り出す際にエラーハンドリング機能も付いている
	//mapから値を取り出す際に第2引数でエラーハンドラーを受け取る
	s, ok := m4[1] //値の取り出しに成功したか否かの戻り値を受け取る
	fmt.Println(s, ok)

	s, ok = m4[len(m4)+1] //存在しないインデックス
	if !ok {
		fmt.Println("キーが存在しません")
	}
	fmt.Println(s, ok)

	// 値の更新や追加は普通に代入式で可能
	m4[2] = "America"
	m4[4] = "China"
	fmt.Println(m4, len(m4))

	// 要素の削除
	delete(m4, 3)
	fmt.Println(m4, len(m4))
}
