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

func get_prime(list []int) ([]int, int) {
	prime_list := []int{}
	search_list := []int{}
	search_list = list
	//探索リストの先頭値が√numberを超えたら終了
	limit := int(math.Sqrt(float64(search_list[len(search_list)-1])))
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

//渡された配列を2つの配列にわける関数
func split_search(want_split []int) (tmp [][]int) {
	raw := want_split
	list1 := []int{}
	list2 := []int{}
	thrice := [][]int{}
	sliceSize := len(raw)
	//3で割ってあまりが1の数列,3で割ってあまりが2の数列にわける
	for i := 0; i < sliceSize; i += 1 {
		if raw[i]%3 == 1 {
			list1 = append(list1, raw[i])
		}
		if raw[i]%3 == 2 {
			list2 = append(list2, raw[i])
		}
	}
	thrice = append(thrice, list1, list2)
	tmp = thrice
	return tmp
}

func main() {
	var k = make_sequence(1000000)
	var p1 = split_search(k)[0]
	var p2 = split_search(k)[1]
	//channel作成
	counter1 := make(chan int, 3)
	counter2 := make(chan int, 3)
	//処理時間計測開始
	start := time.Now()
	//並列処理
	go func() {
		list, count := get_prime(p1)
		_ = list
		fmt.Printf("1つ目のリストの素数は%d個\n", count)
		for {
			time.Sleep(2 * time.Second)
			counter1 <- count
		}
	}()
	go func() {
		list2, count2 := get_prime(p2)
		_ = list2
		fmt.Printf("2つ目のリストの素数は%d個\n", count2)
		for {
			time.Sleep(2 * time.Second)
			counter2 <- count2
		}

	}()
	<-counter1
	<-counter2
	//処理時間計測終了
	end := time.Now()
	fmt.Printf("合計%d個\n", <-counter1+<-counter2)
	fmt.Printf("処理時間は%fs", (end.Sub(start)).Seconds())
}
