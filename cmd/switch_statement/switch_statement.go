package main

import (
	"fmt"
	"time"
)

func main() {
	// 基本的な switch
	fmt.Println("--- 基本的な switch ---")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Friday":
		fmt.Println("Almost weekend")
	default:
		fmt.Println("Midweek day")
	}

	// 複数の条件
	fmt.Println("\n--- 複数の条件 ---")
	fruit := "apple"
	switch fruit {
	case "apple", "banana", "orange":
		fmt.Println("Fruit is available")
	default:
		fmt.Println("Unknown fruit")
	}

	// 条件なしの switch (if の代替)
	fmt.Println("\n--- 条件なしの switch ---")
	age := 20
	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teen")
	default:
		fmt.Println("Adult")
	}

	// 型 switch (interface の型判定)
	fmt.Println("\n--- 型 switch ---")
	var i interface{} = 3.14
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case float64:
		fmt.Println("Float:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}

	// switch の fallthrough (次のケースに強制的にする)
	fmt.Println("\n--- fallthrough ---")
	num := 1
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Pther number")
	}

	// time.Weekday を使った switch
	fmt.Println("\n--- 曜日による switch ---")
	today := time.Now().Weekday()
	switch today {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch の中で変数定義
	fmt.Println("\n--- switch 内で変数定義 ---")
	switch length := len("hello"); {
	case length < 5:
		fmt.Println("Short word")
	case length == 5:
		fmt.Println("Five-letter word")
	default:
		fmt.Println("Long word")
	}
}
