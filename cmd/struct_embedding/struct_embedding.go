package main

import "fmt"

// 埋め込まれる構造体
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

// 埋め込み先の構造体
type Employee struct {
	Person     // Person構造体を埋め込む
	EmployeeID string
	Position   string
}

func main() {
	// Employee のインスタンスを作成
	e := Employee{
		Person: Person{
			Name: "Alice",
			Age:  20,
		},
		EmployeeID: "E12345",
		Position:   "Engineer",
	}

	// 埋め込まれた構造体のフィールドやメソッドに直接アクセスできる
	fmt.Println("Employee Info:")
	fmt.Println("Name:", e.Name) // e.Person.Name と同じ
	fmt.Println("Age:", e.Age)   // e.Person.Age と同じ
	fmt.Println("Employee ID:", e.EmployeeID)
	fmt.Println("Position:", e.Position)

	e.Greet() // e.Person.Greet() と同じ
}
