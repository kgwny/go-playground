package main

import (
	"fmt"
	"strconv"
)

// interface{} の変換
type Robot struct {
	name  string
	birth int
}

func main() {
	// int -> string
	i := 123

	itos1 := strconv.Itoa(i)
	fmt.Println(itos1)

	itos2 := string(i) // 123 が rune(Unicode)として認識される
	fmt.Println(itos2) // Unicode 123(10進数) -> "{"

	// string -> int
	s1 := "12345"
	s2 := "hoge"

	stoi1, _ := strconv.Atoi(s1)
	fmt.Println(stoi1)

	stoi2, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println("string to int の変換失敗 ->", err)
	} else {
		fmt.Println(stoi2)
	}

	// string -> bool
	s3 := "TRUEa"
	stob, err := strconv.ParseBool(s3)
	if err != nil {
		fmt.Println("string to bool の変換失敗 ->", err)
	} else {
		fmt.Println(stob)
	}
	// strconv.ParseBool によりboolean値へ変換できる文字列
	// true: "1", "t", "T", "true", "TRUE", "True"
	// false: "0", "f", "F", "false", "FALSE", "False"

	// bool -> string
	b1 := false
	btos := strconv.FormatBool(b1)
	fmt.Println(btos)

	asimo := Robot{
		name:  "ASIMO",
		birth: 1975,
	}

	Anything := map[string]interface{}{
		"stringVal": "文字列",
		"intVal":    12345,
		"boolVal":   true,
		"structVal": asimo, // Robot
	}

	// interface{} に包含される数値の更新(int)
	Anything["intVal"] = Anything["intVal"].(int) + 1
	fmt.Println("Anything-intVal =", Anything["intVal"])

	// interface{} に包含される数値を文字列として取得したいとき
	i1 := strconv.Itoa(Anything["intVal"].(int))
	fmt.Println("Anything-intVal(string) =", i1)

	// interface{} に包含される文字列の更新
	Anything["stringVal"] = "変更後"
	stringVal := Anything["stringVal"].(string)
	fmt.Println("Anything-stringVal =", stringVal)

	// interface{} に包含される構造体の取得
	robo := Anything["structVal"].(Robot)
	fmt.Println("Anything-structVal =", robo)

	fmt.Println("\n--- map の走査 ---")
	for i, v := range Anything {
		switch v := v.(type) { // 元の型がキャストされる
		case int:
			fmt.Printf("%s: %d (%T)\n", i, v, v)
		case string:
			fmt.Printf("%s: %s (%T)\n", i, v, v)
		case bool:
			fmt.Printf("%s: %t (%T)\n", i, v, v)
		case Robot:
			fmt.Printf("%s: %+v (%T)\n", i, v, v)
		}
	}
}
