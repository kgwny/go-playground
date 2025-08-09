package main

import "fmt"

/*
defer 文は、関数の実行を、その関数が返されるまで延期する
延期された呼び出しの引数はすぐに評価されるが、
関数呼び出しは、その関数が返されるまで実行されない
*/
func main() {
	defer fmt.Println("!")

	defer fmt.Println("world")

	defer fmt.Println(",")

	fmt.Println("Hello")
}
