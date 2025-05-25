package main

import "fmt"

func main() {
	// 配列 (固定長)
	var arr1 [3]int                    // 要素数3のint配列、初期値は0
	arr2 := [3]int{1, 2, 3}            // 明示的な初期化
	arr3 := [...]string{"A", "B", "C"} // 要素数は省略 (コンパイラが自動で判定する)

	fmt.Println("--- 配列 (固定長) ---")
	fmt.Println("配列 arr1:", arr1)
	fmt.Println("配列 arr2:", arr2)
	fmt.Println("配列 arr3:", arr3)

	// 配列の要素にアクセス
	fmt.Println("\n--- 配列の要素にアクセス arr2 の先頭を入れ替える ---")
	arr2[0] = 10
	fmt.Println("変更後 arr2:", arr2)

	// 配列の長さ
	fmt.Println("\n--- 配列の長さ ---")
	fmt.Println("arr3の長さ:", len(arr3))

	// スライス (可変長)
	slice1 := []int{1, 2, 3}    // リテラルで初期化
	slice2 := make([]string, 2) // 長さ2のスライス (空文字が入る)
	slice3 := []int{}           // 空っぽのスライス

	fmt.Println("\n--- スライス (可変長) ---")
	fmt.Println("スライス slice1 (リテラルで初期化):", slice1)
	fmt.Println("スライス slice2 (長さ2のスライス):", slice2)
	fmt.Println("スライス slice3 (空っぽのスライス):", slice3)

	// 要素の追加 (append)
	slice1 = append(slice1, 4, 5)

	fmt.Println("\n--- 要素の追加 ---")
	fmt.Println("append実行後 slice1:", slice1)

	// 他のスライスを結合
	slice3 = append(slice3, slice1...)

	fmt.Println("\n--- slice3 に slice1 を結合する ---")
	fmt.Println("結合後 slice3:", slice3)

	// スライスの部分参照
	part := slice1[1:4] // index 1〜3 (4は含まれない)
	fmt.Println("\n--- スライスの部分参照 ---")
	fmt.Println("slice1 の部分参照 (index 1〜3):", part)

	// スライスの容量と拡張
	fmt.Println("\n--- スライスの容量と拡張 ---")
	s := make([]int, 2, 5) // 長さ2, 容量5
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))
	s = append(s, 1, 2, 3) // 容量を超えると新しい配列が割り当てられる
	fmt.Println("append後 s:", s, "len:", len(s), "cap:", cap(s))

	// 配列とスライスの違い
	fmt.Println("\n--- 配列とスライスの違い ---")
	original := [3]int{1, 2, 3}
	copyA := original // 値コピー
	copyA[0] = 9
	fmt.Println("original から copyA を複製し、copyA の先頭の要素を 9 に入れ替える")
	fmt.Println("original:", original)
	fmt.Println("copyA:", copyA)

	orgSlice := []int{1, 2, 3}
	copyB := orgSlice // 参照コピー
	copyB[0] = 9
	fmt.Println("orgSlice から copyB を複製し、copyB の先頭の要素を 9 に入れ替える")
	fmt.Println("orgSlice:", orgSlice)
	fmt.Println("copyB:", copyB)
	fmt.Println("copyB の要素を入れ替えた結果、orgSlice の要素も入れ替わる")

	// forループによる配列/スライスの走査
	fmt.Println("\n--- forループによる配列/スライスの走査 ---")
	for i, v := range slice1 {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	// スライスのコピー
	fmt.Println("\n--- スライスのコピー ---")
	source := []int{1, 2, 3}
	dest := make([]int, len(source))
	copy(dest, source)
	fmt.Println("コピー元:", source)
	fmt.Println("コピー先:", dest)

	// 2次元スライス
	fmt.Println("\n--- 2次元スライスの例 ---")
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("2次元スライス:", matrix)

	// nilスライス vs 空スライス
	var nilSlice []int    // nilスライス
	emptySlice := []int{} // 空スライス

	fmt.Println("\n--- nilスライス ---")
	fmt.Println("nilSlice:", nilSlice, "nilか？", nilSlice == nil)

	fmt.Println("\n--- 空スライス ---")
	fmt.Println("emptySlice:", emptySlice, "nilか？", emptySlice == nil)

	// スライスを引数として関数に渡す
	nums := []int{1, 2, 3}
	incrementAll(nums)
	fmt.Println("関数呼び出し後 nums:", nums) // 元のスライスが変更される
}

// すべての要素を+1する関数 (スライスは参照渡しされる)
func incrementAll(arr []int) {
	for i := range arr {
		arr[i] += 1
	}
}
