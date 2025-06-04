package main

import "fmt"

// カウンター関数を生成するクロージャー (無名関数)
func createCounter() func() int {
	counter := 0 // この変数がクロージャー内で保持される
	return func() int {
		counter++
		return counter
	}
}

func main() {
	// 1つ目のカウンターを作成
	counterA := createCounter()

	fmt.Println("Counter A:", counterA()) // 1
	fmt.Println("Counter A:", counterA()) // 2
	fmt.Println("Counter A:", counterA()) // 3

	// 2つ目のカウンターを作成
	counterB := createCounter()

	fmt.Println("Counter B:", counterB()) // 1
	fmt.Println("Counter B:", counterB()) // 2

	// 再び Counter A を呼び出す -> 状態が保持されていることがわかる
	fmt.Println("Counter A:", counterA()) // 4
}
