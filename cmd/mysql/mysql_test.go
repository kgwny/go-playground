package main

import (
	"testing"
)

var testDSN = "user:password@tcp(127.0.0.1:3306)/testdb"

func TestConnectMySQL(t *testing.T) {
	db, err := connectMySQL(testDSN)
	if err != nil {
		t.Fatalf("DB接続に失敗しました: %v", err)
	}
	defer db.Close()
}

func TestGetUsers(t *testing.T) {
	db, err := connectMySQL(testDSN)
	if err != nil {
		t.Fatalf("DB接続に失敗しました: %v", err)
	}
	defer db.Close()

	names, err := getUsers(db)
	if err != nil {
		t.Fatalf("ユーザー取得に失敗しました: %v", err)
	}

	if len(names) == 0 {
		t.Log("ユーザーが0件でした。テスト用データが必要です。")
	} else {
		t.Logf("取得したユーザー数: %d", len(names))
	}
}
