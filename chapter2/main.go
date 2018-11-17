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

func main() {
  func_var()
  func_if()
  func_for()
}
