/*
二、資料處理-字串
請寫出將字串「人易科技:上 機 測 驗 - 演算法」中的:
1. 字元「:」改成全型
2. 去掉中文字間的空白(保留一號二側空白)
3. 列印出符號:到-之間的字元
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	my_string := "人易科技:上 機 測 驗 - 演算法"
	my_string_bytes := []byte(my_string)
	fmt.Printf("原字串")
	fmt.Println(string(my_string_bytes))

	//1. 字元「:」改成全型
	//半形(:)=58 全形（：）=[239, 188, 154]
	new_string_1 := strings.Replace(my_string, ":", "：", 1)
	fmt.Printf("1. 字元「:」改成全型\n")
	fmt.Println(new_string_1)

	//2. 去掉中文字間的空白(保留一號二側空白)

	fmt.Printf("2. 去掉中文字間的空白(保留一號二側空白)\n")
	new_string_2 := strings.Replace(new_string_1, " ", "", 3)
	fmt.Println(new_string_2)

	//3. 列印出符號:到-之間的字元

	fmt.Printf("3. 列印出符號:到-之間的字元\n")
	string_slice_1 := strings.Split(new_string_2, "：")
	string_slice_2 := strings.Split(string_slice_1[1], "-")
	fmt.Println(string_slice_2[0])
}
