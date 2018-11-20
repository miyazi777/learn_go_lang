package main

import (
  _ "strings"   // 使わないパッケージの先頭に_をつけてコンパイルエラーを抑止できる
  . "fmt"
  //f "fmt"     // aliasを指定することも可能
  "errors"
)


func func_var() {
  Println("--- variable ---")
  var message string = "hello world"  // var 変数名 データ型が変数定義の基本
  Println(message)  // message

  // 同じ型の変数を以下のように複数帝具することもできる
  var (
    a string = "aaa"
    b = "bbb"
  )
  Println(a) // aaa
  Println(b) // bbb

  // 省略記法
  msg := "hello"
  Println(msg) // hello

  // 定数
  const Hello string = "Hello"
  Println(Hello)  // Hello

  // 初期化しなかった場合のデフォルト値はデータ型によって決まっている
  var i int
  Println(i)    // 0

  Println("------")
}

func func_if_smaple(a int, b int) {
  // if文の基本
  if a > b {
    Println("a is larger than b")
  } else if a < b {
    Println("a is smaller than b")
  } else {
    Println("a equals b")
  }
}

func func_if() {
  Println("---if---")
  func_if_smaple(100, 10)   // a is larger than b
  func_if_smaple(10, 100)   // a is smaller than b
  func_if_smaple(10, 10)    // a eauls b
  Println("------")
}

func func_for() {
  Println("---for---")
  // 基本
  for i := 0; i < 3; i++ {
    Println(i)
  }
  // whileっぽいやり方
  n := 0
  for n < 3 {
    Printf("n = %d\n", n)
    n++
  }
  // continueとbreak
  m := 0
  for {
    if m < 3 {
      Printf("m = %d\n", m)
      m++
      continue
    }
    break
  }
  Println("------")
}

// switchの基本
func func_switch_basic(n int) {
  switch n {
  case 15:
    Println("FizzBuzz")
  case 5, 10:
    Println("Buzz")
  case 3,6,9:
    Println("Fizz")
  default:
    Println(n)
  }
}

// switch処理内にfailthroughを付けると次のswitch文へ処理を
func func_switch_failthrough() {
  n := 3
  switch n {
  case 3:
    Println(n)
    n = n - 1
    fallthrough
  case 2:
    Println(n)
    n = n - 1
    fallthrough
  default:
    Println("default")
  }
}

// switchには式が使える
func func_switch_eval(n int) {
  switch {
  case n%15 == 0:
    Println("FuzzBuzz")
  case n%5 == 0:
    Println("Buzz")
  case n%3 == 0:
    Println("Fizz")
  default:
    Println(n)
  }
}

func func_switch() {
  Println("---switch---")
  func_switch_basic(15)  // FizzBuzz
  func_switch_basic(5)   // Buzz
  func_switch_basic(2)   // Fizz
  func_switch_failthrough() // 3\n2\ndefault
  func_switch_eval(15)   // FizzBuzz
  func_switch_eval(5)    // Buzz
  func_switch_eval(3)    // Fizz
  func_switch_eval(2)    // default
  Println("------")
}

// 関数基本
func func_func_basic() {
  Println("hello")
}

// 引数の型が同じ場合は以下のように型を指定可能
func func_sum(i, j int) {
  Println(i+j)
}

// 複数の戻り値
func func_swap(i, j int) (int, int) {
  return j, i
}

// エラーを返却する関数
func func_error_rtn(i, j int) (int, error) {
  if j == 0 {
    return 0, errors.New("divied by zero")
  }
  return i / j, nil
}

// 名前付き戻り値
func func_name_rtn(v int) (result int) {
  result = v + 1
  return
}

// 関数リテラル
func func_literal() {
  func(i, j int) {
    Println(i+j)
  }(3, 5)
}

// 無名関数を変数に入れることも可能
var sum_func func(i, j int) = func(i, j int) {
  Println(i + j)
}

func func_func() {
  Println("---func---")
  func_func_basic() // hello
  func_sum(1, 3)  // 3
  x, y := 3, 4
  x, y = func_swap(x, y)
  Println(x, y)   // 4 3
  n, err := func_error_rtn(10, 0)
  Println(n, err) // 0 divied by zero
  n2, err2 := func_error_rtn(10, 2)
  Println(n2, err2)  // 5 <nil>
  Println(func_name_rtn(5))  // 6
  func_literal() // 8
  sum_func(3, 4) // 7
  Println("------")
}

func main() {
  func_var()
  func_if()
  func_for()
  func_switch()
  func_func()
}
