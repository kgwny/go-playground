package main

import (
	"encoding/json"
	"fmt"
)

// 1. 基本的な構造体定義
type Person struct {
	Name string
	Age  int
}

// 2. メソッド (値レシーバ)
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// 3. メソッド (ポインタレシーバ)
func (p *Person) Birthday() {
	p.Age++
}

// 4. 埋め込み構造体 (疑似継承)
type Employee struct {
	Person     // 埋め込み
	EmployeeID string
}

// 5. 構造体スライスとマップ
type Team struct {
	Name     string
	Members  []Person
	Projects map[string]string
}

// 6. JSONとの変換
type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// 7. インターフェースの実装
type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name + " says: Woof!")
}

func main() {
	// 基本的な構造体の使用
	p1 := Person{Name: "Alice", Age: 30}
	p1.Greet()

	// ポインタレシーバの使用
	p1.Birthday()
	p1.Greet()

	// 構造体のポインタを直接初期化
	p2 := &Person{Name: "Bob", Age: 25}
	p2.Birthday()
	p2.Greet()

	// 埋め込み構造体の使用
	e1 := Employee{
		Person:     Person{Name: "Carol", Age: 40},
		EmployeeID: "E123",
	}
	fmt.Printf("%s's employee ID is %s\n", e1.Name, e1.EmployeeID)

	// スライスとマップを含む構造体
	team := Team{
		Name: "Dev Team",
		Members: []Person{
			{Name: "Dave", Age: 22},
			{Name: "Eve", Age: 27},
		},
		Projects: map[string]string{
			"ProjectA": "Dave",
			"ProjectB": "Eve",
		},
	}
	fmt.Printf("Team: %s has %d members\n", team.Name, len(team.Members))

	// JSONエンコード
	product := Product{Name: "Laptop", Price: 1299.99}
	jsonData, _ := json.Marshal(product)
	fmt.Println("JSON:", string(jsonData))

	// JSONデコード
	var decodedProduct Product
	json.Unmarshal(jsonData, &decodedProduct)
	fmt.Printf("Decoded Product: %+v\n", decodedProduct)

	// インターフェースの使用
	var s Speaker = Dog{Name: "Fido"}
	s.Speak()
}
