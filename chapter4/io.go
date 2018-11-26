package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

// ファイル生成のサンプル
func createFile() {
	// ファイルを生成
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// ファイルクローズ
	defer file.Close()
}

// ファイルへの書き込みサンプル
func writeFile() {
	// ファイルを生成
	file, err := os.Create("./file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	// ファイルクローズ
	defer file.Close()

	// 書き込むデータを[]byteで用意する
	message := []byte("hello world\n")

	// Writeを使って書き込む
	_, err = file.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

// ファイルへの書き込みサンプル2(WriteStringとFprint関数使用版)
func writeFile2() {
	// ファイルを生成
	file, err := os.Create("./file3.txt")
	if err != nil {
		log.Fatal(err)
	}
	// ファイルクローズ
	defer file.Close()

	// WriteStringを使って書き込む
	_, err = file.WriteString("hello world\n")
	if err != nil {
		log.Fatal(err)
	}

	// Fprintを使って書き込む
	_, err = fmt.Fprint(file, "hello world2\n")
	if err != nil {
		log.Fatal(err)
	}
}

// ファイル読み出しサンプル
func readFile() {
	file, err := os.Open("./file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 12byte格納可能なスライスを用意する
	message := make([]byte, 12)

	// ファイル内のデータをスライスに読み出す
	_, err = file.Read(message)
	if err != nil {
		log.Fatal(err)
	}

	// 文字列にして表示
	fmt.Print(string(message))
}

// jsonをファイルに書き込み
func writeJson() {
	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang memo",
	}

	// ファイル開く
	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// エンコーダの取得
	encoder := json.NewEncoder(file)

	// JSONエンコードしたデータの書き込み
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}
}

// jsonファイルを読み込み
func readJson() {
	file, err := os.Open("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var person Person

	// デコレータの取得
	decoder := json.NewDecoder(file)

	// JSONデコードしたデータの書き込み
	err = decoder.Decode(&person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(person) // {1 Gopher gopher@example.org 5  }
}

func main() {
	createFile()
	writeFile()
	writeFile2()
	readFile()

	writeJson()
	readJson()
}
