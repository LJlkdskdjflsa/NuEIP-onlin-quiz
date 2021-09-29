/*
六、兩數總和
•給定一個整數陣列 nums 和一個整數 target ,當兩數總和為target時,返回兩數的索引。
•您可以假設每個輸入都只有一個解決方案,並且您可能不會兩次使用相同的元素。
•您可以按任何順序返回答案。
限制條件:
2 <= nums.length <= 103
-109 <= nums[j] <= 109
-109 <= target <= 109
僅存在一個有效答案。
*/

/**
* @param Integer[] $nums
* @param Integer $target
* @return Integer[]
**/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	//println(hashTable)
	for i, num := range nums {
		if idx, ok := hashTable[target-num]; ok {
			return []int{idx, i}
		}
		hashTable[num] = i
	}
	return nil
}
func main() {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{1, 2, 4, 5, 7}
	result1 := twoSum(nums1, 9)
	result2 := twoSum(nums2, 9)
	fmt.Println(result1)
	fmt.Println(result2)

}
