/*
五、邏輯處理-交集、差集、聯集
今有二陣列,請寫出資料處理程式碼 陣列a:775,5,22,13,55,974,796,1,09
陣列b:0.1,2,3,4,5,6,7,8,9
1. 陣列c = 陣列a交集陣列b 2. 陣列d = 陣列a 差集陣列b 3. 陣列e = 陣列a聯集陣列b
| 本題需自行完成演算法,不可使用現成交集/差集/聯集函式
intersection, union, difference
*/
package main

import "fmt"

func intersection(a []int, b []int) []int {
	var result []int

	for _, num_a := range a {
		for _, num_b := range b {
			if num_a == num_b {
				result = append(result, num_b)
			} else {

			}
		}
	}
	return result
}
func difference(a []int, b []int) []int {
	intersaction := intersection(a, b)
	union := union(a, b)

	var result []int
	m := make(map[int]bool)
	for _, item := range intersaction {
		m[item] = true
	}
	for _, item := range union {
		if _, ok := m[item]; !ok {
			result = append(result, item)
		}
	}
	return result
}
func union(a []int, b []int) []int {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}

func main() {
	a := []int{775, 5, 22, 13, 55, 974, 796, 1, 0, 9}

	b := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("1. 陣列c = 陣列a交集陣列b 2. 陣列d = 陣列a 差集陣列b 3. 陣列e = 陣列a聯集陣列b\n")
	fmt.Printf("交集\n")
	fmt.Println(difference(b, a))
	fmt.Printf("差集\n")
	fmt.Println(intersection(b, a))
	fmt.Printf("聯集\n")
	fmt.Println(union(b, a))

}
