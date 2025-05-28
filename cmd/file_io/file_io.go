package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// 基本設定
	fileName := "sample.txt"
	appendFileName := "sample_append.txt"
	dir := "example_dir"
	os.MkdirAll(dir, 0755)

	// ファイルへの書き込み
	content := "Hello, Golang!\nThis is a test file.\n"
	err := os.WriteFile(fileName, []byte(content), 0644)
	checkError(err)
	fmt.Println("ファイルへの書き込みを完了")

	// ファイルからの読み込み (全体)
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("ファイルの内容:\n" + string(data))

	// os + bufio による行単位の読み込み
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println("行単位で読み込み:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	checkError(scanner.Err())

	// 書き込み (os + bufio + Flush)
	newFileName := filepath.Join(dir, "buffered_write.txt")
	newFile, err := os.Create(newFileName)
	checkError(err)
	writer := bufio.NewWriter(newFile)
	writer.WriteString("This is a buffered write.\n")
	writer.WriteString("Second line.\n")
	writer.Flush()
	newFile.Close()
	fmt.Println("Buffered write 完了")

	// 追記モードで書き込み
	f, err := os.OpenFile(appendFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)
	_, err = f.WriteString("Appended line.\n")
	checkError(err)
	f.Close()
	fmt.Println("追記書き込み完了")

	// io.Copy を使ったファイルコピー
	source, err := os.Open(fileName)
	checkError(err)
	defer source.Close()

	destFile := filepath.Join(dir, "copy.txt")
	dest, err := os.Create(destFile)
	checkError(err)
	defer dest.Close()

	_, err = io.Copy(dest, source)
	checkError(err)
	fmt.Println("ファイルコピー完了")

	// バイナリデータの書き込みと読み込み
	binaryFile := filepath.Join(dir, "binary.dat")
	err = os.WriteFile(binaryFile, []byte{0x00, 0x01, 0x02, 0x03}, 0644)
	checkError(err)
	binData, err := os.ReadFile(binaryFile)
	checkError(err)
	fmt.Printf("バイナリファイルの内容: %v\n", binData)

	// bytes.Buffer を使った仮想ファイル処理
	var buf bytes.Buffer
	buf.WriteString("This is a virtual file in memory.\n")
	io.Copy(os.Stdout, &buf)
	fmt.Println("メモリ上の仮想ファイル処理完了")
}

// エラー処理用関数
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
