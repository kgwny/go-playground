package main

import (
	"fmt"
)

func main() {
	// 基本のfor文 (C言語スタイル)
	fmt.Println("--- Basic for loop ---")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 条件だけのfor文 (while のような使い方)
	fmt.Println("\n--- For with only condition (like while) ---")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// 無限ループ
	fmt.Println("\n--- Infinite loop (break after 3 interations) ---")
	j := 0
	for {
		if j >= 3 {
			break
		}
		fmt.Println("j =", j)
		j++
	}

	// range を使った配列のループ
	fmt.Println("\n--- For range over slice ---")
	slice := []string{"apple", "banana", "cherry"}
	for index, value := range slice {
		fmt.Printf("index = %d, value = %s\n", index, value)
	}

	// range を使った map のループ
	fmt.Println("\n--- For range over string (runes) ---")
	str := "Go言語"
	for i, r := range str {
		fmt.Printf("Index %d, Rune %q\n", i, r)
	}

	// continue の使用
	fmt.Println("\n--- Using continue ---")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// break の使用
	fmt.Println("\n--- Using break ---")
	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// ネストされたfor文
	fmt.Println("\n--- Nested loops ---")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}

	// ラベル付き break
	fmt.Println("\n--- Using Labeled break ---")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// ラベル付き continue
	fmt.Println("\n--- Using Labeled continue ---")
row:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue row
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// forでフィボナッチ数列 (10個)
	fmt.Println("\n--- Fibonacci sequence ---")
	a, b := 0, 1
	for n := 0; n < 10; n++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}
