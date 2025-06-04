package main

import "fmt"

// 再帰(recursion)の例
// 階乗を再帰で計算する関数
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// 5の階乗 -> 1 * 2 * 3 * 4 * 5 = 120
	num := 5
	result := factorial(num)
	fmt.Printf("%d! = %d\n", num, result)
}
