package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `お問い合わせは example@example.com または support@example.jp までご連絡ください。`

	// 正規表現パターン (簡易的なメールアドレス形式)
	pattern := `[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`

	// 正規表現のコンパイル
	re := regexp.MustCompile(pattern)

	// 一致するすべての文字列を抽出
	matches := re.FindAllString(text, -1)

	// 結果の表示
	fmt.Println("見つかったメアド:")
	for _, match := range matches {
		fmt.Println(match)
	}
}
