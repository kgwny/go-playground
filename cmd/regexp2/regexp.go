package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Learn Regular Expressions in Go Language."

	// 単語全体をマッチする正規表現パターン
	pattern := `(\w+)`

	// 正規表現をコンパイルする
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// マッチングの実行
	matches := re.FindAllString(text, -1)
	fmt.Println("Matches found:", matches)
}
