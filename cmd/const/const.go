package main

import (
	"fmt"
	"math"
	"time"
)

// グローバルスコープでの定数定義
const AppName = "GoConstExample"
const Version = "1.0.0"

// 型付き定数
const Pi float64 = 3.1415926535
const Greeting string = "Hello, world!"

// iota を使った列挙型的な定数
const (
	_  = iota             // 0 をスキップする場合はこのように記述する
	KB = 1 << (10 * iota) // 1 << 10
	MB                    // 1 << 20
	GB                    // 1 << 30
	TB                    // 1 << 40
)

// iota を使った曜日の定義
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 型なし定数と型推論の例
const DefaultPort = 8080
const PiTimes2 = Pi * 2

// 定数式での計算
const CircleRadius = 5
const CircleArea = Pi * CircleRadius * CircleRadius

func main() {
	fmt.Println("App:", AppName)
	fmt.Println("Version:", Version)
	fmt.Println("Greeting:", Greeting)

	fmt.Println("\n--- 定数による演習 ---")
	fmt.Printf("円の半径: %d, 面積: %.2f\n", CircleRadius, CircleArea)

	fmt.Println("\n--- 型なし定数の使用 ---")
	printPort(DefaultPort)

	fmt.Println("\n--- iota を使った容量定数 ---")
	fmt.Println("1 KB =", KB)
	fmt.Println("1 MB =", MB)
	fmt.Println("1 GB =", GB)
	fmt.Println("1 TB =", TB)

	fmt.Println("\n--- iota を使った曜日 ---")
	fmt.Println("Monday =", Monday)
	fmt.Println("Friday =", Friday)

	fmt.Println("\n--- シャドウイングの例 ---")
	const LocalMessage = "Main Scope"
	fmt.Println("LocalMessage:", LocalMessage)
	{
		const LocalMessage = "Inner Scope"
		fmt.Println("LocalMessage in inner block:", LocalMessage)
	}

	fmt.Println("\n--- math パッケージを使った定数式 ---")
	const Angle = 30.0
	radian := Angle * (math.Pi / 180)
	fmt.Printf("Angle: %.2f degrees = %.4f radians\n", Angle, radian)

	fmt.Println("\n--- time.Duration による定数の利用 ---")
	const Timeout = 5 * time.Second
	fmt.Println("Timeout duration:", Timeout)
}

func printPort(port int) {
	fmt.Printf("Default server port is %d\n", port)
}
