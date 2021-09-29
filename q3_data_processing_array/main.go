/*
三、資料處理-
陣列
今有陣列資料 0.1.2.3.4.5.6.7.8.9 請寫出資料處理程式碼
1. 計算出「奇數值總和」減去「偶數值總和」
2. 分割成二陣列,分別存放「偶數值」及「奇數值」
*/

package main

import "fmt"

func odd_even_subtract(nums [10]int) int {
	odd_sum := 0
	even_sum := 0
	for _, num := range nums {
		if num%2 == 0 {
			even_sum += num
		} else {
			odd_sum += num
		}
	}

	return (odd_sum - even_sum)
}
func odd_even_seperate(nums [10]int) ([]int, []int) {
	odd_sum := []int{}
	even_sum := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			even_sum = append(even_sum, num)
		} else {
			odd_sum = append(odd_sum, num)
		}
	}

	return odd_sum, even_sum
}

func main() {
	nums := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("原陣列\n")
	fmt.Println(nums)
	fmt.Printf("1. 計算出「奇數值總和」減去「偶數值總和」\n")
	fmt.Println(odd_even_subtract(nums))
	fmt.Printf("2. 分割成二陣列,分別存放「偶數值」及「奇數值」\n")
	fmt.Println(odd_even_seperate(nums))
}
