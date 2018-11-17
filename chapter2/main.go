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

func func_if() {
  Println("---if---")



  Println("------")
}

func main() {
  func_var()
  func_if()
}
