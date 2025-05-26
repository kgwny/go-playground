package main

import (
	"fmt"
	"strings"
)

// 値レシーバの例
type User struct {
	Name string
	Age  int
}

// 値レシーバ: コピーされるため元の値は変更されない
func (u User) Greet() string {
	return fmt.Sprintf("Hello, my name is %s.", u.Name)
}

// ポインタレシーバ: 元の値を変更可能
func (u *User) Birthday() {
	u.Age++
}

// メソッドチェーンの例 (ポインタレシーバで自身を返す)
func (u *User) SetName(name string) *User {
	u.Name = name
	return u
}

func (u *User) SetAge(age int) *User {
	u.Age = age
	return u
}

// 構造体の埋め込みによるメソッドの継承
type Admin struct {
	User  // User のメソッドが Admin に引き継がれる
	Level int
}

// 構造体の匿名フィールドとメソッドのオーバーライド
func (a Admin) Greet() string {
	return fmt.Sprintf("Hi, I am admin %s with level %d.", a.Name, a.Level)
}

// インターフェースの定義と実装
type Greeter interface {
	Greet() string
}

type Dog struct {
	Name string
}

func (d Dog) Greet() string {
	return fmt.Sprintf("Woof! I'm %s the dog!", d.Name)
}

// メソッドの中で他のメソッドを呼び出す
type Text struct {
	Content string
}

func (t Text) WordCount() int {
	words := strings.Fields(t.Content)
	return len(words)
}

func (t Text) Summary() string {
	return fmt.Sprintf("Text contains %d words.", t.WordCount())
}

func main() {
	// User の例
	user := User{Name: "Alice", Age: 30}
	fmt.Println(user.Greet())
	user.Birthday() // 値レシーバなので user.Age は変わらない
	fmt.Printf("After birthday (value receiver): %d\n", user.Age)

	// ポインタレシーバでメソッドを呼び出す
	userPtr := &user
	userPtr.Birthday()
	fmt.Printf("After birthday (pointer receiver): %d\n", user.Age)

	// メソッドチェーン
	userPtr.SetName("Bob").SetAge(40)
	fmt.Printf("Updated user: %+v\n", user)

	// Admin の例 (構造体の埋め込み)
	admin := Admin{
		User:  User{Name: "Charlie", Age: 50},
		Level: 10,
	}
	fmt.Println(admin.Greet())      // Admin の Greet を呼び出し
	fmt.Println(admin.User.Greet()) // 埋め込まれた User の Greet を呼び出し

	// インターフェースを使ったポリモーフィズム
	var g Greeter

	g = user
	fmt.Println("Greeter (User):", g.Greet())

	g = Dog{Name: "Pochi"}
	fmt.Println("Greeter (Dog):", g.Greet())

	// Text 構造体の例
	text := Text{Content: "Go is expressive, concise, clean, and efficient."}
	fmt.Println(text.Summary())
}
