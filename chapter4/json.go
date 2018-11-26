package main

import (
	"encoding/json"
	. "fmt"
	"log"
)

type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

// 構造体 -> jsonの基本
func basic() {
	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
	}

	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	Println(string(b)) // {"ID":1,"Name":"Gopher","Email":"gopher@example.org","Age":5,"Address":""}
}

type Person2 struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"-"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"`
	memo    string
}

// タグ付けのサンプル
func tag() {
	person := &Person2{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
	}

	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	Println(string(b)) // {"id":1,"name":"Gopher","age":5}
}

// JSON文字列から構造体
func unmarshal() {
	var person Person
	b := []byte(`{"id":1, "name":"Gopher", "age":5}`)
	err := json.Unmarshal(b, &person)
	if err != nil {
		log.Fatal(err)
	}
	Println(person) // {1 Gopher  5  }
}

func main() {
	basic()
	tag()
	unmarshal()
}
