/*
四、資料排序-正序
今有一陣列資料 77,5,5,22,13,55,97.4.796.1.0.9 請寫出正序排列處理程式碼
| 本題需自行完成演算法,不可使用現成排序函式
*/
package main

import (
	"fmt"
)

func quickSort(nums []int) []int {
	if len(nums) < 2 { //停止遞迴條件
		return nums
	}
	left, right := 0, len(nums)-1 //左右界
	//center := rand.Int() % (len(nums) - 1)                // random sernder index
	//nums[center], nums[right] = nums[right], nums[center] //switch

	for i, _ := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left] //switch
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]

	quickSort((nums[:left]))
	quickSort((nums[left+1:]))

	return nums
}
func main() {

	my_array := []int{77, 5, 5, 22, 13, 55, 97, 4, 796, 1, 0, 9}
	fmt.Printf("陣列排序前 \n")
	fmt.Println(my_array)
	quickSort(my_array)
	fmt.Printf("陣列排序 \n")

	fmt.Println(my_array)
}
