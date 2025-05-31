package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// User は構造体タグを持つサンプル構造体です
type User struct {
	ID    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email,omitempty" validate:"email"`
}

func main() {
	// サンプルデータの作成
	user := User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}

	// JSONエンコード (struct tags の json 部分を確認)
	jsonBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON Output:")
	fmt.Println(string(jsonBytes))

	// reflect を使って struct tags を読み取る
	t := reflect.TypeOf(user)
	fmt.Println("\nStruct Tags:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %-5s  json tag: %-15s  validate tag: %s\n",
			field.Name,
			field.Tag.Get("json"),
			field.Tag.Get("validate"))
	}
}
