package main

import "fmt"

func main() {
	// 基本的なif文
	num := 5
	if num > 0 {
		fmt.Println("num is positive")
	}

	// if-else文
	if num%2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}

	// if-else if-else文
	if num < 0 {
		fmt.Println("num is negative")
	} else if num == 0 {
		fmt.Println("num is zero")
	} else {
		fmt.Println("num is positive")
	}

	// if内での短絡変数定義 (スコープはif文内のみ)
	if val := num * 2; val > 5 {
		fmt.Println("val is greater than 5:", val)
	}

	// ネストされたif文
	if num > 0 {
		if num < 10 {
			fmt.Println("num is between 1 and 9")
		}
	}

	// ブール値を直接使用
	isActive := true
	if isActive {
		fmt.Println("Active flag is true")
	}

	// 複数条件 (AND)
	age := 25
	if age > 18 && age < 30 {
		fmt.Println("Age is between 19 and 29")
	}

	// 複数条件 (OR)
	grade := "B"
	if grade == "A" || grade == "B" {
		fmt.Println("Good grade")
	}

	// 否定 (NOT)
	if !isActive {
		fmt.Println("This won't print")
	}

	// if文で比較関数を利用
	if isEven(num) {
		fmt.Println("num is even (checked via function)")
	} else {
		fmt.Println("num is odd (checked via function)")
	}
}

func isEven(n int) bool {
	return n%2 == 0
}
