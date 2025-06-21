package main

import (
	"fmt"
)

func init() {
	// init 関数は特殊な関数であり、パッケージの初期化に使われる
	// main パッケージに書くと main 関数より先に実行される
	// main パッケージでない場合は、import するだけで呼び出される
	fmt.Println("Hello", "init")
}

func hoge() {
	fmt.Println("Hello", "hoge")
}

func main() {
	// init() -> main 関数からは呼び出せない
	hoge()
	fmt.Println("Hello", "main")
}
