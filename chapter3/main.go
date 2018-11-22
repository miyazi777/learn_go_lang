package main

import (
	. "fmt"
)

type ID int       // IDという型を定義
type Priority int // Priorityという型を定義

func custom_type(id ID, priority Priority) {
	Println("---custom type---")
	Println(id)       // 1
	Println(priority) // 3

	Println("------")
}

// 構造体
type Task struct { // 構造体の最初の文字が大文字の場合、publicとなり、小文字の場合はpackageに閉じたスコープとなる
	ID     int
	Detail string
	done   bool
}

func func_callByRef(task *Task) {
	task.done = true
}

// コンストラクタ
// 実際にはコンストラクタはないのでNewで始まる関数を定義してその内部で構造体を生成するのが、goの慣例
func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

func func_struct() {
	Println("---struct---")
	var task Task = Task{
		ID:     1,
		Detail: "buy the milk",
		done:   true,
	}
	Println(task.ID)     // 1
	Println(task.Detail) // buy the milk
	Println(task.done)   // true

	var task2 Task = Task{
		2, "buy the juice", false,
	} // 定義順に値を渡すことも可能
	Println(task2.ID)     // 2
	Println(task2.Detail) // buy th juce
	Println(task2.done)   // false

	// 初期値を指定しない場合はゼロ値が設定される
	task3 := Task{}
	Println(task3.ID)     // 0
	Println(task3.Detail) // ""
	Println(task3.done)   // true

	// 参照渡しが可能
	task4 := &Task{ID: 100, Detail: "ref test", done: false}
	func_callByRef(task4)
	Println(task4.ID)     // 100
	Println(task4.Detail) // "ref test"
	Println(task4.done)   // true

	// newで初期化することも可能
	var task5 *Task = new(Task)
	task5.ID = 5
	task5.Detail = "buy the milk"
	task5.done = false
	Println(task5.ID)     // 5
	Println(task5.Detail) // "buy the milk
	Println(task5.done)   // false

	// new関数を使った初期化
	task6 := NewTask(11, "test6")
	Println(task6.ID)     // 11
	Println(task6.Detail) // "test6"
	Println(task6.done)   // false

	Println("------")
}

func (task Task) String() string {
	str := Sprintf("%d) %s\n", task.ID, task.Detail)
	return str
}

func (task *Task) Finish() {
	task.done = true
}

func method_test() {
	Println("---method---")
	task := NewTask(1, "test method")
	Printf("%s", task.String()) // 1) test method
	Printf("%s", task)          // 上と同じ
	task.Finish()
	Println(task.done) // true

	Println("------")
}

func main() {
	var id ID = 1
	var priority Priority = 5
	custom_type(id, priority)
	//custom_type(priority, id)   // これはコンパイルエラーとなる
	func_struct()
	method_test()
}