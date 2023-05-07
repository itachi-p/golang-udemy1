package main

import "fmt"

// ラベル付きfor文

func main() {

	// breakで一気にネストしたループを抜けたい場合
	/*
		Loop:
			for {
				for {
					for {
						fmt.Println("START")
						// ただのbreakだと一番近いfor文を抜けるだけで無限ループ
						break Loop
					}
					fmt.Println("処理をしたくない部分")
				}
				fmt.Println("処理をしたくない部分")
			}
			fmt.Println("END")
	*/

	// continueと組み合わせる例
	// j=1の時だけ表示したい
Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i, j, i*j)
		}
		// ただのcontinueでは一番近いforを抜けるだけで、ここに入ってしまう
		fmt.Println("処理をしたくない部分")
	}
}
