package main

import "fmt"

//ビット演算

const (
	//1ビットを63箇所左にずらして巨大な数を作る(1の後に0が63個続く2進数)
	//
	Big = 1 << 63
	//これをまた60箇所右にずらすと、1<<3、つまり2進数で1000、10新数では8になる
	Small = Big >> 60
)

func needInt(x int) int { return x }
func needFloat(x float64) float64 {
	return x
}

func main() {
	fmt.Printf("Type:%T, Value:%d\n", needInt(Small), needInt(Small))
	//fmt.Println(Big)     //そのままでは境界値ギリギリアウトでoverflowエラーになる
	fmt.Println(Big - 1) //-1すればintのMAX値として表示可能（int64かint32か等は環境依存）
	fmt.Println(-Big)    //頭にマイナス符号を付ければintのMIN値として表示可能

	//int最大値(境界値)に+1すると符号が反転する(ビット演算で考えると理解できる)
	fmt.Println(needInt(Big-1) + 1) //結果的に=Bigに等しいが、エラーにならない
	fmt.Printf("Type:%T, Value:%f\n", needFloat(Small), needFloat(Small))
	fmt.Printf("Type:%T, Value:%g\n", 8.0, 8.0) //書式指定子%vや%gだと整数の場合intと同様の表示
	fmt.Println(needFloat(Big))                 //浮動小数点表記法で正の整数も1の位以下×桁数で表される
}
