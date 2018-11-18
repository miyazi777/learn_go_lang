package main

import (
  _ "strings"   // 使わないパッケージの先頭に_をつけてコンパイルエラーを抑止できる
  . "fmt"
  //f "fmt"     // aliasを指定することも可能
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
  func_switch_basic(15)  // FizzBuzz
  func_switch_basic(5)   // Buzz
  func_switch_basic(2)   // Fizz
  func_switch_failthrough() // 3\n2\ndefault
  func_switch_eval(15)   // FizzBuzz
  func_switch_eval(5)    // Buzz
  func_switch_eval(3)    // Fizz
  func_switch_eval(2)    // default
}

func main() {
  func_var()
  func_if()
  func_for()
  func_switch()
}
