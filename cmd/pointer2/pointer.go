package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// イミュータブル変数は実体で宣言する。
	// ポインタは、別の変数に値を渡した場合、その変数からも元の値を更新することができる。
	// 複数の変数間で1つの値を共有するミュータブル変数という扱いになる。
	v1 := &s1{Number: 1}
	fmt.Println("v1:", v1)

	v2 := v1
	fmt.Println("v2:", v2)

	v2.Number = 10
	fmt.Printf("v1: %v, v2: %v\n", v1, v2)

	// 実体を別の変数に渡した場合、実体から複製した別の値を生成する。
	// 値を管理する変数が1つのイミュータブル変数となる。
	v3 := s1{Number: 1}
	fmt.Println("v3:", v3)

	v4 := v3
	fmt.Println("v4:", v4)

	v4.Number = 10
	fmt.Printf("v3: %v, v4: %v\n", v3, v4)

	// golangにおいては、関数の引数やメソッドのレシーバをポインタで宣言することができる。
	// 前提として、関数をコールするときは、指定された引数の元値を複製して受け取る。
	// このとき、関数の引数をポインタにすればアドレス値を引数として読み込むことができる。
	v5 := &s2{Number: 100, Name: "Taro", Address: []byte("65535")}
	CheckSizeByPointer(v5)

	// 関数の引数にポインタではなく実体を渡す場合、実体そのものを複製するため、
	// ポインタを利用するときよりも、実体を渡すときのほうが複製にかかるメモリ領域が増大する。
	// そのため、関数の引数やメソッドのレシーバはパフォーマンス観点によりポインタで宣言すべき。
	v6 := s2{Number: 100, Name: "Taro", Address: []byte("65535")}
	CheckSize(v6)

	// ゼロ値/null値を見分ける必要がある場合はポインタを使用する。
	// ポインタで変数宣言する際に、初期値を指定しない場合、初期値はnilとなる。
	// nilはアドレス値が存在しないことを表す。
	var id1 *int
	var name1 *string
	fmt.Printf("Init Value: id=%v, name=%v\n", id1, name1)
	// id, name ともに初期値は nil となる

	var id2 int
	var name2 string
	fmt.Printf("Init Value: id=%v, name=%v\n", id2, name2)
	// id は初期値が 0、name は初期値が空文字となる
}

type s1 struct {
	Number int
}

type s2 struct {
	Number  int
	Name    string
	Address []byte
}

func CheckSizeByPointer(s *s2) {
	fmt.Printf("Size: %v bytes\n", unsafe.Sizeof(s))
}

func CheckSize(s s2) {
	fmt.Printf("Size: %v bytes\n", unsafe.Sizeof(s))
}
