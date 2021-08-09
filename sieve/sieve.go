package main

import (
	"fmt"
	"math"
	"time"
)

//配列の特定の要素を削除する関数
func remove(s_list []int, index int) (tmp []int) {
	tmp = append(s_list[:index], s_list[(index+1):]...)
	return
}

func get_prime(number int) ([]int, int) {
	//初期化
	prime_list := []int{}
	search_list := []int{}
	//2からnumberまでの数字の配列を作る
	for i := 2; i < number+1; i++ {
		search_list = append(search_list, i)
	}
	//探索リストの先頭値が√numberを超えたら終了
	limit := int(math.Sqrt(float64(number)))
	for search_list[0] <= limit {
		//探索リストの先頭を素数リストに移動
		p_num := search_list[0]
		prime_list = append(prime_list, p_num)
		//探索リストの先頭を削除
		search_list = remove(search_list, 0)
		//p_numの倍数を探索リストから篩落とす
		search_list_length := len(search_list)
		tmp := []int{}
		for i := 0; i < search_list_length; i++ {
			if search_list[i]%p_num != 0 {
				tmp = append(tmp, search_list[i])
			}
		}
		search_list = tmp
	}
	//探索リストの残りを素数リストに追加
	prime_list = append(prime_list, search_list...)
	return prime_list, len(prime_list)
}

func main() {
	start := time.Now()
	list, count := get_prime(10000000)
	_ = list
	fmt.Printf("合計:%d個\n", count)
	end := time.Now()
	fmt.Printf("%fs", (end.Sub(start)).Seconds())
}
