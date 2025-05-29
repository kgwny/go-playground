package main

import (
	"errors"
	"fmt"
)

// Logger はログ出力用のインターフェース (依存対象)
type Logger interface {
	Info(msg string)
	Error(msg string)
}

// ConsoleLogger は Logger を実装する具体型
type ConsoleLogger struct{}

func (cl ConsoleLogger) Info(msg string) {
	fmt.Println("[INFO]", msg)
}

func (cl ConsoleLogger) Error(msg string) {
	fmt.Println("[ERROR]", msg)
}

// User はユーザー情報を持つ構造体
type User struct {
	Name   string
	Age    int
	logger Logger
}

// NewUser は User 構造体のファクトリ関数
// logger を依存として注入し、バリデーションを行う
func NewUser(name string, age int, logger Logger) (*User, error) {
	if age < 0 {
		logger.Error("invalid age provided")
		return nil, errors.New("age cannot be negative")
	}
	logger.Info("creating new user")

	return &User{
		Name:   name,
		Age:    age,
		logger: logger,
	}, nil
}

// Greet は User のメソッド
func (u *User) Greet() {
	u.logger.Info(fmt.Sprintf("User %s is greeting", u.Name))
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", u.Name, u.Age)
}

func main() {
	logger := ConsoleLogger{}

	user, err := NewUser("Taro", 25, logger)
	if err != nil {
		fmt.Println("Failed to create user:", err)
		return
	}

	user.Greet()
}
