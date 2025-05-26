package main

import (
	"errors"
	"fmt"
)

// 関数の書式
// func 関数名(パラメータ) 戻り値の型 {
//     ...
// }

// パラメータ、戻り値のいずれもない関数
func hello() {
	fmt.Println("Hello world!")
}

// 2つのパラメータを受け取って結果を返す関数
func add(a, b int) int {
	return a + b
}

// 複数の返却値を持つ関数 (多値関数)
// ひとつの変数で受け取ろうとした場合はエラー扱いとなる
func swap(a, b int) (int, int) {
	// これと同じ処理内容 -> b, a = a, b
	return b, a
}

// 配列から最大値を持つ要素を特定の上、インデックスと値を返却する
func findMaxVal(arr []int) (index int, value int) {
	value = arr[0]
	for i := 1; i < len(arr); i++ {
		if value < arr[i] {
			value = arr[i]
			index = i
		}
	}
	return index, value
}

// 可変長引数
// 関数のパラメータ定義で型名の前に ... というプレフィックスを指定すると、
// 可変長引数を扱うことができる
func sum(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}

// 以下の fibonacci 関数は、n番目のフィボナッチ数を返却する
// n が負の場合はエラーを返す
func fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("input must be a non-negative integer")
	}

	if n == 0 {
		return 0, nil
	} else if n == 1 {
		return 1, nil
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr, nil
}

func main() {
	hello()

	a := 10
	b := 20
	c := add(a, b)
	fmt.Println("c =", c)

	x, y := swap(a, b)              // 10, 20
	fmt.Println("x =", x, "y =", y) // 20, 10

	w, _ := swap(a, b)              // 返却値のうち1つめの値だけほしいとき
	_, z := swap(a, b)              // 返却値のうち2つめの値だけほしいとき
	fmt.Println("w =", w, "z =", z) // 20, 10

	arr1 := []int{1, 2, 4, 8, 16, 32, 64, 128}
	i, v := findMaxVal(arr1)
	fmt.Println("index =", i, "value =", v)

	s := sum(1, 2, 3, 4, 5)
	println(s)

	// 可変長引数を受け取る関数へパラメータ値としてスライスを指定したい場合は、
	// 以下のとおり、スライスの後ろに ... を付与する
	s1 := []int{1, 2, 3, 4, 5}
	r1 := sum(s1...)
	println(r1)

	// エラー発生パターン
	// golangには例外処理がないので、
	// 代わりに関数とその呼出元でエラーハンドリングしてあげる必要がある
	n := -10
	result, err := fibonacci(n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Printf("Fibonacci(%d) = %d\n", n, result)
	}
}
