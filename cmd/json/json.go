package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

func main() {
	// --- 基本: 構造体 -> JSON ---
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age,omitempty"`
		Note string `json:"-"`
	}
	user := User{Name: "Alice", Age: 30}
	b, _ := json.Marshal(user)
	fmt.Println("基本 Marshal:", string(b))

	// --- JSON -> 構造体 ---
	jsonStr := []byte(`{"name":"Bob","age":25}`)
	var user2 User
	json.Unmarshal(jsonStr, &user2)
	fmt.Println("基本 Unmarshal:", user2.Name, user2.Age)

	// --- map[string]interface{} ---
	var m map[string]interface{}
	json.Unmarshal([]byte(`{"name":"Tom","age":22}`), &m)
	fmt.Println("動的 JSON:", m["name"], m["age"])

	// --- スライスのエンコード ---
	users := []User{{"Alice", 30, ""}, {"Bob", 25, ""}}
	b, _ = json.Marshal(users)
	fmt.Println("スライス Marshal:", string(b))

	// --- ネスト構造 ---
	type Address struct {
		City string `json:"city"`
	}
	type UserWithAddress struct {
		Name    string  `json:"name"`
		Address Address `json:"address"`
	}
	jsonStr = []byte(`{"name":"Taro","address":{"city":"Tokyo"}}`)
	var u3 UserWithAddress
	json.Unmarshal(jsonStr, &u3)
	fmt.Println("ネスト構造:", u3.Name, u3.Address.City)

	// --- 整形出力 ---
	b, _ = json.MarshalIndent(user, "", "  ")
	fmt.Println("整形出力:\n" + string(b))

	type Event struct {
		Title string `json:"title"`
		Day   Date   `json:"day"`
	}
	eventJSON := []byte(`{"title":"Meeting","day":"2025-06-01"}`)
	var event Event
	json.Unmarshal(eventJSON, &event)
	fmt.Println("カスタム型:", event.Title, time.Time(event.Day).Format("2006-01-02"))

	// --- RawMessage ---
	type Envelope struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}
	envStr := []byte(`{"type":"greeting","data":{"msg":"hello"}}`)
	var env Envelope
	json.Unmarshal(envStr, &env)
	fmt.Println("RawMessage 部分デコード:", env.Type, string(env.Data))

	// --- DisallowUnknownFields ---
	jsonStr = []byte(`{"name":"Bob","age":25,"extra":"value"}`)
	dec := json.NewDecoder(bytes.NewReader(jsonStr))
	dec.DisallowUnknownFields()
	var user4 User
	err := dec.Decode(&user4)
	if err != nil {
		fmt.Println("DisallowUnknownFields エラー:", err)
	}

	// --- JSON ストリーム処理 ---
	type Message struct {
		Msg string `json:"msg"`
	}
	stream := `{"msg":"first"}{"msg":"second"}`
	dec = json.NewDecoder(bytes.NewReader([]byte(stream)))
	for {
		var msg Message
		if err := dec.Decode(&msg); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("ストリーム エラー:", err)
			break
		}
		fmt.Println("ストリームから読み込み:", msg.Msg)
	}

	// --- null 対応: ポインタを使う ---
	type NullableUser struct {
		Name *string `json:"name"`
	}
	jsonStr = []byte(`{"name":null}`)
	var nu NullableUser
	json.Unmarshal(jsonStr, &nu)
	if nu.Name == nil {
		fmt.Println("null フィールドは nil として扱われる")
	}
}

// --- カスタム型の Marshal/Unmarshal ---
type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	*d = Date(t)
	return err
}
