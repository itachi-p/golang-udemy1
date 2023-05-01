package main

import "fmt"

// for文

func main() {
	/*
		i := 0
		for {
			i++
			if i == 3 {
				break
			}
			fmt.Println(i)
		}

		// 条件付きfor文
		point := 0
		for point < 10 {
			fmt.Println(point)
			point++
		}
	*/

	// 古典的for文
	/*
		for i := 0; i < 10; i++ {
			if i == 3 || i == 6 {
				continue // 処理をスキップしてfor文に戻る
			}
			if i == 8 {
				break // ループ処理を強制終了
			}
			fmt.Println(i)
		}
	*/

	// 配列(Goでは配列は固定長)
	/*
		arr := [3]string{"a", "b", "c"}
		for i := 0; i < len(arr); i++ {
			fmt.Printf("%d番目の要素: %s\n", i+1, arr[i])
		}

		// 範囲式for文
		// len等で配列の要素数を取り出さずとも同じことができる
		arr2 := [3]int{1, 2, 3}

		// 配列の要素のインデックス番号と値を一緒に取り出せる
		for i, v := range arr2 {
			fmt.Println(i, v)
		}

		// インデックス番号を使いたくない場合
		for _, v := range arr2 {
			fmt.Println(v)
		}

		// 逆にインデックス番号だけを使いたい場合
		for i := range arr2 {
			fmt.Println(i)
		}
	*/

	// Slices（可変長配列に似たデータ型）の場合
	sl := []string{"Python", "PHP", "Go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// Mapを使った連想配列(Key & Value)
	m := map[string]int{"Apple": 398, "Banana": 198}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
