package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connectMySQL(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func getUsers(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, rows.Err()
}

func main() {
	// ユーザー名:パスワード@tcp(ホスト:ポート)/データベース名
	dsn := "user:password@tcp(127.0.0.1:3306)/testdb"

	// DB接続
	db, err := connectMySQL(dsn)
	if err != nil {
		log.Fatalf("MySQL への接続に失敗しました: %v", err)
	}
	defer db.Close()
	fmt.Println("MySQL に接続成功")

	// ユーザー一覧を取得
	users, err := getUsers(db)
	if err != nil {
		log.Fatalf("ユーザー情報の取得に失敗しました: %v", err)
	}

	// 出力
	fmt.Println("ユーザー一覧:")
	for _, name := range users {
		fmt.Println("- ", name)
	}
}
