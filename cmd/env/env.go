package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの取得
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env が読み込めませんでした。: %v", err)
	}

	// 環境変数の取得
	m1 := os.Getenv("SAMPLE_MESSAGE")
	fmt.Println(m1)

	// 環境変数に追加して取得
	os.Setenv("ADD_MESSAGE", "HOGE")
	m2 := os.Getenv("ADD_MESSAGE")
	fmt.Println(m2)

	// 環境変数の削除
	os.Unsetenv("SAMPLE_MESSAGE")
	m3 := os.Getenv("SAMPLE_MESSAGE")
	fmt.Println(m3)

	// すべての環境変数を削除
	os.Clearenv()
	m4 := os.Getenv("ADD_MESSAGE")
	fmt.Println(m4)
}
