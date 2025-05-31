package main

import (
	"fmt"
	"math"
)

// ----------------------------
// 基本的な interface の定義と使用
// ----------------------------

// 1. インターフェース定義
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. struct の定義
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

// 3. Shape インターフェースを実装
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// -----------------------------
// 空インターフェース (interface{})
// -----------------------------

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ----------------------
// 型アサーションと型スイッチ
// ----------------------

func doSomething(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}
}

// ----------------------------
// インターフェースの埋め込み (合成)
// ----------------------------

type Reader interface {
	Read() string
}

type Writer interface {
	Write(s string)
}

type ReadWriter interface {
	Reader
	Writer
}

type Console struct{}

func (c Console) Read() string {
	return "reading from console"
}

func (c Console) Write(s string) {
	fmt.Println("writing to console:", s)
}

// ------------------------------
// インターフェースを引数・戻り値に使う
// ------------------------------

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// ---------------------------
// インターフェースの nil チェック
// ---------------------------

type Stringer interface {
	String() string
}

type MyType struct{}

func (m *MyType) String() string {
	return "MyType instance"
}

func printStringer(s Stringer) {
	if s == nil {
		fmt.Println("Stringer is nil")
	} else {
		fmt.Println(s.String())
	}
}

func main() {
	// 基本的なインターフェースの使用
	var s Shape

	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	s = c
	printShapeInfo(s)

	s = r
	printShapeInfo(s)

	// 空インターフェースと型アサーション
	describe(42)
	describe("hello")
	describe(3.14)

	doSomething(123)
	doSomething("Gopher")
	doSomething(true)

	// 埋め込みインターフェースの使用
	var rw ReadWriter = Console{}
	fmt.Println(rw.Read())
	rw.Write("hello interface!")

	// インターフェースが nil かどうかのチェック
	var mt *MyType = nil
	printStringer(mt) // ポインタが nil の場合、実装していても nil 扱いにされる
	var st Stringer = mt
	printStringer(st) // インターフェース型は nil でない (mt型としてnil)

	// インターフェースのスライス
	shapes := []Shape{
		Circle{Radius: 2.5},
		Rectangle{Width: 3, Height: 4},
	}
	for _, shape := range shapes {
		printShapeInfo(shape)
	}
}
