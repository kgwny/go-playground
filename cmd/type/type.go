package main

import (
	"fmt"
	"reflect"
)

// 変数のデータ型を表示する
func main() {
	var i int = 65535
	var s string = "abcde"

	fmt.Printf("i のデータ型: %s\n", reflect.TypeOf(i))
	fmt.Printf("s のデータ型: %s\n", reflect.TypeOf(s))
}
