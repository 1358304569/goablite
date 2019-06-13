/*
做一些小的测试
*/

package main

import (
	"fmt"
)

/*1. 测试短横杠 _ */
func main() {
	nums := []int{0, 1, 2, 4}

	// 正常操作
	// for i, n := range nums {
	// 	fmt.Println("index:", i)
	// 	fmt.Println("value", n)
	// }

	for n := range nums {
		fmt.Println("value", n)
	}
}
