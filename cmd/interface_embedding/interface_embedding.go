package main

import "fmt"

// Printer インターフェース
type Printer interface {
	Print(doc string)
}

// Scanner インターフェース
type Scanner interface {
	Scan() string
}

// MultiFunctionDevice は Printer と Scanner を埋め込んだインターフェース
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// MyDevice 構造体は MultiFunctionDevice を実装
type MyDevice struct {
	Name string
}

// Print メソッドの実装
func (d MyDevice) Print(doc string) {
	fmt.Printf("[%s] Printing document: %s\n", d.Name, doc)
}

// Scan メソッドの実装
func (d MyDevice) Scan() string {
	scanned := fmt.Sprintf("[%s] Scanned document content", d.Name)
	fmt.Println(scanned)
	return scanned
}

func main() {
	device := MyDevice{Name: "OfficeDevice"}

	// MultiFunctionDevice インターフェースとして使う
	var mfd MultiFunctionDevice = device

	// Print と Scan の使用
	mfd.Print("AnnualReport.pdf")
	content := mfd.Scan()

	fmt.Println("Scan result received:", content)
}
