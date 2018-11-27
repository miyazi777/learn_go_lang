package main

import (
  "log"
  "io/ioutil"
  . "fmt"
)

//func readFile() {
//  file, _ = os.Open("./file3.txt")
//  message, _ := ioutil.ReadAll(file)
//  Println(message)
//}

func readFile2() {
  message, _ := ioutil.ReadFile("./file3.txt")
  Println(string(message))
}

func readWriteFile() {
  Println("test")
  // ファイルにメッセージを書き込む
  hello := []byte("hello world\n")
  err := ioutil.WriteFile("./hello_file.txt", hello, 0666)
  if err != nil {
    log.Fatal(err)
  }

  // ファイルの中身を全て読みだす
  message, err := ioutil.ReadFile("./hello_file.txt")
  if err != nil {
    log.Fatal(err)
  }
  Println(string(message))
}

func main() {
  //readFile()
  readFile2()
  readWriteFile()
}


