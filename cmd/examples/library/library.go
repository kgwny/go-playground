package main

import (
	"errors"
	"fmt"
)

// BookGenre: iota を使った列挙型風の定義
type BookGenre int

const (
	Fiction BookGenre = iota
	NonFiction
	Science
	History
	Other
)

func (g BookGenre) String() string {
	return [...]string{"Fiction", "NonFiction", "Science", "History", "Other"}[g]
}

// Book: 基本的な構造体
type Book struct {
	Title  string
	Author string
	Genre  BookGenre
}

// Library: 書籍を管理する構造体（スライスの操作）
type Library struct {
	Catalog []Book
}

// AddBook: 書籍を追加
func (lib *Library) AddBook(book Book) {
	lib.Catalog = append(lib.Catalog, book)
}

// FindBookByTitle: 書名で検索（エラーハンドリング付き）
func (lib Library) FindBookByTitle(title string) (Book, error) {
	for _, book := range lib.Catalog {
		if book.Title == title {
			return book, nil
		}
	}
	return Book{}, errors.New("book not found")
}

// Printable: インタフェース
type Printable interface {
	Print()
}

// Print: BookがPrintableを実装
func (b Book) Print() {
	fmt.Printf("📘 %s by %s [%s]\n", b.Title, b.Author, b.Genre)
}

func main() {
	lib := Library{}

	// defer: 最後に表示されるメッセージ
	defer fmt.Println("✅ 処理が完了しました。")

	// 書籍追加
	lib.AddBook(Book{"The Go Programming Language", "Alan A. A. Donovan", Science})
	lib.AddBook(Book{"Sapiens", "Yuval Noah Harari", History})

	// 書籍検索と表示
	titlesToFind := []string{"Sapiens", "Unknown Book"}
	for _, title := range titlesToFind {
		book, err := lib.FindBookByTitle(title)
		if err != nil {
			fmt.Printf("❌ エラー: \"%s\" は見つかりませんでした。\n", title)
			continue
		}
		var p Printable = book
		p.Print()
	}
}
