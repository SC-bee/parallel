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

func make_sequence(number int) (search_list []int) {
	//初期化
	list := []int{}
	//2からnumberまでの数字の配列を作る
	for i := 2; i < number+1; i++ {
		list = append(list, i)
	}
	search_list = list
	return search_list
}

func get_prime(number int) ([]int, int) {
	prime_list := []int{}
	search_list := []int{}
	search_list = make_sequence(number)
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

//渡された配列を複数の配列にわける関数
func split_search(want_split []int, split_num int) (tmp [][]int) {
	raw := want_split
	thrice := [][]int{}
	sliceSize := len(raw)
	//元の配列を指定の数の配列にわける
	for i := 0; i < sliceSize; i += sliceSize / split_num {
		end := i + sliceSize/split_num
		if sliceSize < end {
			end = sliceSize
		}
		if i == 0 {
			thrice = append(thrice, raw[i:end])
		}
		if i >= 1 {
			thrice = append(thrice, raw[i:end+i])
		}
	}
	tmp = thrice
	return tmp
}

func main() {

	fmt.Printf("%d\n", split_search(make_sequence(100), 2))
	var p1 = split_search(make_sequence(100), 2)[0]
	var p2 = split_search(make_sequence(100), 2)[1]
	fmt.Printf("%d&%d\n", p1, p2)
	start := time.Now()
	list, count := get_prime(10)
	fmt.Println(list)
	fmt.Printf("count:%d個\n", count)
	end := time.Now()
	fmt.Printf("%fms", (end.Sub(start)).Seconds())
}
