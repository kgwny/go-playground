package main

import (
	"fmt"
	"time"
)

func main() {
	// 現在の時刻を取得
	now := time.Now()
	fmt.Println("現在時刻:", now)

	// 日時の構成要素を取得
	fmt.Println("\n--- 日時の構成要素を取得 ---")
	fmt.Println("年:", now.Year())
	fmt.Println("月:", now.Month())
	fmt.Println("日:", now.Day())
	fmt.Println("時:", now.Hour())
	fmt.Println("分:", now.Minute())
	fmt.Println("秒:", now.Second())
	fmt.Println("ナノ秒:", now.Nanosecond())
	fmt.Println("曜日:", now.Weekday())

	// 指定した日付を作成
	fmt.Println("\n--- 指定した日付を作成 ---")
	t := time.Date(2025, 5, 25, 13, 30, 0, 0, time.Local)
	fmt.Println("指定日付:", t)

	/*
	 * 日付のフォーマット (レイアウト文字列を使う)
	 * time.Format や time.Parse のレイアウトは "2006-01-02 15:04:05" の形式に固定されており、
	 * これは Go における特有の仕様です（この日時がマジックナンバー）。
	 */
	fmt.Println("\n--- 日付のフォーマット ---")
	fmt.Println("フォーマット:", now.Format("2006-01-02 15:04:05"))

	// パース (文字列 -> time.Time)
	fmt.Println("\n--- 日時のパース ---")
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, "2025-05-25 13:30:00")
	if err != nil {
		fmt.Println("Parse error:", err)
	} else {
		fmt.Println("パース結果:", parsedTime)
	}

	// 時間の加算・減算
	fmt.Println("\n--- 時間の加算と減算 ---")
	fmt.Println("1時間後:", now.Add(1*time.Hour))
	fmt.Println("24時間前:", now.Add(-24*time.Hour))

	// 2つの時間の差分
	fmt.Println("\n--- 2つの時間の差分 ---")
	diff := t.Sub(now)
	fmt.Println("差分:", diff)

	// タイマー (1回だけ)
	fmt.Println("\n--- タイマー (1回だけ) ---")
	fmt.Println("3秒後にメッセージ表示")
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("3秒経過しました")
	})
	time.Sleep(4 * time.Second) // 実行を待機

	// 繰り返しタイマー (Ticker)
	fmt.Println("\n--- 繰り返しタイマー (Ticker) ---")
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for i := 0; i < 3; i++ {
			<-ticker.C
			fmt.Println("1秒ごとの処理")
		}
		ticker.Stop()
	}()
	time.Sleep(4 * time.Second)

	// 時間の比較
	fmt.Println("\n--- 時間の比較 ---")
	fmt.Println("指定日時 t は 現在時刻 now より後か？", t.After(now))
	fmt.Println("指定日時 t は 現在時刻 now より前か？", t.Before(now))
	fmt.Println("指定日時 t は 現在時刻 now と等しいか？", t.Equal(now))

	// UTCとローカル時刻
	fmt.Println("\n--- UTCとローカル時刻 ---")
	fmt.Println("UTC:", now.UTC())
	fmt.Println("ローカル:", now.Local())

	// 時間の丸め (Truncate, Round)
	fmt.Println("\n--- 時間の丸め (Truncate, Round) ---")
	fmt.Println("5分単位に丸める (切り捨て):", now.Truncate(5*time.Minute))
	fmt.Println("5分単位に丸める (四捨五入):", now.Round(5*time.Minute))

	// Unixタイム
	fmt.Println("\n--- Unixタイム ---")
	fmt.Println("Unix秒:", now.Unix())
	fmt.Println("Unixナノ秒:", now.UnixNano())
	fmt.Println("Unix秒から時刻へ:", time.Unix(now.Unix(), 0))

	// 月末を計算する例
	fmt.Println("\n--- 月末の計算例 ---")
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	firstOfNextMonth := firstOfMonth.AddDate(0, 1, 0)
	lastOfMonth := firstOfNextMonth.Add(-time.Nanosecond)
	fmt.Println("月末:", lastOfMonth)
}
