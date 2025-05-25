package main

import "fmt"

// 構造体の定義
type Person struct {
	name string
	age  int
}

func main() {
	// 基本的なポインター操作
	fmt.Println("--- 基本的なポインター操作 ---")
	x := 10
	p := &x // x のポインターを取得する
	fmt.Println("x =", x)
	fmt.Println("*p =", *p) // ポインターを使って x の値にアクセス

	*p = 20 // ポインター経由で x の値を変更
	fmt.Println("x =", x)

	// ポインター引数の関数
	fmt.Println("\n--- ポインター引数の関数 ---")
	val := 5
	fmt.Println("Before double:", val)
	double(&val)
	fmt.Println("After double:", val)

	// 配列とスライスとポインター
	fmt.Println("\n--- 配列とスライスとポインター ---")
	arr := [3]int{1, 2, 3}
	modifyArray(&arr)
	fmt.Println("Modified array:", arr)

	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println("Modified slice:", slice)

	// ポインターと構造体
	fmt.Println("\n--- ポインターと構造体 ---")
	p1 := Person{name: "Alice", age: 30}
	fmt.Println("Before update:", p1)
	updatePerson(&p1)
	fmt.Println("After update:", p1)

	// new キーワードでポインター生成
	fmt.Println("\n--- new キーワードでポインターを生成 ---")
	y := new(int) // int型のゼロ値へのポインターを取得
	fmt.Println("Value of y (new int):", *y)
	*y = 100
	fmt.Println("Updated value of y:", *y)

	// ポインターのポインター
	fmt.Println("\n--- ポインターのポインター ---")
	a := 42
	ptr1 := &a
	ptr2 := &ptr1
	fmt.Println("a:", a)
	fmt.Println("*ptr1:", *ptr1)
	fmt.Println("**ptr2:", **ptr2)

	// nilポインターの扱い
	fmt.Println("\n--- nilポインターの扱い ---")
	var np *int
	if np == nil {
		fmt.Println("np is nil")
	}
	// *np = 1 // 実行すると panic: nil pointer dereference になるので注意
}

// 値を2倍にする関数 (ポインター引数)
func double(n *int) {
	*n = *n * 2
}

// 配列の内容を変更する関数
func modifyArray(arr *[3]int) {
	for i := range arr {
		arr[i] += 10
	}
}

// スライスは参照型なのでポインター渡ししなくても良い
func modifySlice(s []int) {
	for i := range s {
		s[i] += 100
	}
}

// 構造体のポインターを受け取ってフィールドを更新
func updatePerson(p *Person) {
	p.name = "Bob"
	p.age += 5
}
