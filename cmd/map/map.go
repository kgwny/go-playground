package main

import (
	"fmt"
	"sort"
)

func main() {
	// map の作成と初期化
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("初期化された map:", ages)

	// 要素の追加・更新
	ages["Charlie"] = 20 // 追加
	ages["David"] = 12   // 追加
	ages["Alice"] = 31   // 上書き
	fmt.Println("map における要素の追加・更新:", ages)

	// 要素の取得
	aliceAge := ages["Alice"]
	fmt.Println("map における要素の取得 (Alice の年齢):", aliceAge)

	// 存在チェック (ゼロ値との区別)
	age, exists := ages["David"]
	if exists {
		fmt.Println("map における要素の取得 (David の年齢):", age)
	} else {
		fmt.Println("想定外の挙動 (David は存在しない)")
	}

	// 要素の削除
	delete(ages, "Bob")
	fmt.Println("map から Bob を削除:", ages)

	// 要素数
	fmt.Println("map の長さ:", len(ages))

	// 空の map を make で作成
	scores := make(map[string]float64)
	scores["math"] = 85.5
	scores["english"] = 92.0
	fmt.Println("make で作成した map:", scores)

	// ループ (for range) で map を走査
	fmt.Println("map のループ:")
	for name, age := range ages {
		fmt.Printf(" %s is %d years old\n", name, age)
	}

	// map のキーだけ、または値だけを取り出す
	fmt.Println("map のキーと値を分離:")
	for name := range ages {
		fmt.Println("名前:", name)
	}
	for _, age := range ages {
		fmt.Println("年齢:", age)
	}

	// map のキーでソートして出力
	fmt.Println("ソートされた map の出力:")
	var keys []string
	for key := range ages {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf(" %s: %d\n", key, ages[key])
	}

	// map のネスト (map の中に map がある状態)
	students := map[string]map[string]int{
		"Tom": {
			"Math":    90,
			"Science": 80,
		},
		"Jerry": {
			"Math":    70,
			"Science": 75,
		},
	}
	fmt.Println("ネストされた map:")
	for name, subjects := range students {
		fmt.Println("  ", name)
		for subject, score := range subjects {
			fmt.Printf("    %s: %d\n", subject, score)
		}
	}

	// map の値としてスライス
	fruits := map[string][]string{
		"red":   {"apple", "cherry"},
		"green": {"grape", "melon"},
	}
	fmt.Println("値がスライスで構成された map:")
	for color, list := range fruits {
		fmt.Printf("  %s: %v\n", color, list)
	}

	// map を関数に渡す
	fmt.Println("map を関数に渡す:")
	printMap(ages)

	// map は参照型 (変更は関数に渡した map に反映される)
	updateMap(ages)
	fmt.Println("関数で map を変更した後:", ages)

	// nil map を使う場合の注意点
	var nilMap map[string]int
	// nilMap["a"] = 1 // panic: assignment to entry in nil map
	fmt.Println("nil map の初期値:", nilMap)
}

// map を受け取る関数
func printMap(m map[string]int) {
	for k, v := range m {
		fmt.Printf("  %s: %d\n", k, v)
	}
}

// map を変更する関数 (参照渡し)
func updateMap(m map[string]int) {
	m["NewPerson"] = 100
}
