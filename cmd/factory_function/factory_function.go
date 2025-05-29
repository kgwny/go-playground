package main

import "fmt"

// User はユーザーを表す構造体です
type User struct {
	ID   int
	Name string
	Age  int
}

// NewUser は User 構造体のファクトリ関数（コンストラクタのようなもの）です
func NewUser(id int, name string, age int) User {
	return User{
		ID:   id,
		Name: name,
		Age:  age,
	}
}

// ファクトリ関数にポインタを使う場合
func RegisterUser(id int, name string, age int) *User {
	return &User{
		ID:   id,
		Name: name,
		Age:  age,
	}
}

func main() {
	// ファクトリ関数を使って User を生成
	user1 := NewUser(1, "Taro", 25)
	fmt.Printf("ID: %d, Name: %s, Age: %d\n", user1.ID, user1.Name, user1.Age)

	user2 := RegisterUser(2, "Jiro", 20)
	fmt.Printf("ID: %d, Name: %s, Age: %d\n", user2.ID, user2.Name, user2.Age)
}
