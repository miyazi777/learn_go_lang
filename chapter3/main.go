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

// このように記述することでTask構造体にメソッド定義できる
func (task Task) String() string {
	str := Sprintf("%d) %s\n", task.ID, task.Detail)
	return str
}

// こちらもメソッド定義
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

// インターフェイス
//// 検証用のcarインターフェイス
type Car interface {
	run(int) string
	stop()
}

//// 検証用の構造体
type MyCar struct {
	name  string
	speed int
}

//// インターフェイスで定義されたメソッドを実装する
func (u *MyCar) run(speed int) string {
	u.speed = speed
	return Sprintf("%dkmで走ります", u.speed)
}

func (u *MyCar) stop() {
	Println("停止します")
	u.speed = 0
}

func func_interface() {
	Println("---interface---")

	myCar := &MyCar{name: "マイカー", speed: 0}
	Println(myCar.run(50))
	myCar.stop()

	Println("------")
}

func func_interface_any_type() {
	Println("---interface(any type)---")

	// 空のインターフェイス定義
	var x interface{}
	num := 0
	str := "hello"
	// どんな型でも代入可能
	x = num
	Println(x)
	x = str
	Println(x)

	Println("------")
}

// 空のインターフェイスに値を入れた場合の型判定しているサンプル
func func_interface_any_type2() {
	Println("---interface(any type)2---")
	type Element interface{}
	var element Element = "hello"

	if value, ok := element.(int); ok {
		Printf("int value:%d\n", value)
	} else if value, ok := element.(string); ok {
		Printf("string value:%s\n", value) // string value:hello
	} else {
		Println("other")
	}

	// switch文をつかった場合のやり方
	switch value := element.(type) {
	case int:
		Printf("int value:%d\n", value)
	case string:
		Printf("string value:%s\n", value) // string value:hello
	default:
		Println("other")
	}

	Println("------")
}

// 構造Taskに対してUserを埋め込む実験
type User struct {
  FirstName string
  LastName string
}

func (u *User) FullName() string {
  fullname := Sprintf("%s %s", u.FirstName, u.LastName)
  return fullname
}

func NewUser(firstName, lastName string) *User {
  return &User{
    FirstName: firstName,
    LastName: lastName,
  }
}

type Task2 struct {
  ID int
  Detail string
  done bool
  *User // Userを埋め込む
}

func NewTask2(id int, detail, firstName, lastName string) *Task2 {
  task := &Task2{
    ID: id,
    Detail: detail,
    done: false,
    User: NewUser(firstName, lastName),
  }
  return task
}

func func_extend() {
	Println("---Extend---")

  task := NewTask2(1, "buy the milk", "Jack", "Daniel")
  Println(task.FirstName) // Jack
  Println(task.LastName)  // Daniel
  Println(task.FullName())    // Jack Daniel
  Println(task.User.FirstName)  // Jack
  Println(task.User.LastName)   // Daniel

	Println("------")
}

// インターフェイスの埋め込み
type Arm interface {
  setArm(string)
}

type Body interface {
  setBody(string)
}

type Robot interface {
  Arm
  Body
}

//// 検証用の構造体
type MyRobot struct {
	armName string
	bodyName string
}

//// インターフェイスで定義されたメソッドを実装
func (r *MyRobot) setArm(arm string) {
	r.armName = arm
}

func (r *MyRobot) setBody(body string) {
	r.bodyName = body
}

func func_extend_interface() {
	robot := MyRobot{}
  robot.setArm("10t arm")
  robot.setBody("100t body")
  Println(robot.armName)    // 10t arm
  Println(robot.bodyName)   // 100t body
}

// 型の変換(キャストに失敗するとパニックが発生する
func func_cast() {
  Println("---cast---")
  var i uint8 = 3
  var j uint32 = uint32(i)  // unit8 > unit32
  Println(j)      // 3

  var s string = "abc"
  var b []byte = []byte(s)  // string > []byte
  Println(b)  // [97 98 99]
  Println("------")
}

func main() {
	var id ID = 1
	var priority Priority = 5
	custom_type(id, priority)
	//custom_type(priority, id)   // これはコンパイルエラーとなる
	func_struct()
	method_test()
	func_interface()
	func_interface_any_type()
	func_interface_any_type2()
  func_extend()
  func_extend_interface()
  func_cast()
}
