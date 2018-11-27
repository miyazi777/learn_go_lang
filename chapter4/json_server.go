package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
  "strconv"
  "io/ioutil"
  "html/template"
)

type Person struct {
  ID int `json: "id"`
  Name string `json:"name"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "hello world")
}

// テンプレートのコンパイル
var t = template.Must(template.ParseFiles("index.html"))

func PersonHandler(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()  // 処理の最後にBodyを閉じる

  if r.Method == "POST" {
    // リクエストボディをJSONに変換
    var person Person
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&person)
    if err != nil {
      log.Fatal(err)
    }
    // ファイル名を{id}.txtとする
    filename := fmt.Sprintf("%d.txt", person.ID)
    file, err := os.Create(filename)  // ファイルを生成
    if err != nil {
      log.Fatal(err)
    }
    defer file.Close()

    // ファイルにNameを書き込む
    _, err = file.WriteString(person.Name)
    if err != nil {
      log.Fatal(err)
    }

    // レスポンスとしてステータスコード201を送信
    w.WriteHeader(http.StatusCreated)
  } else if r.Method == "GET" {
    // パラメータを取得
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
      log.Fatal(err)
    }

    // ファイル読み込み
    filename := fmt.Sprintf("%d.txt", id)
    b, err := ioutil.ReadFile(filename)
    if err != nil {
      log.Fatal(err)
    }

    // personを生成
    person := Person{
      ID: id,
      Name: string(b),
    }

    // レスポンスにエンコーディングしたHTMLを書き込む
    t.Execute(w, person)
  }
}

func main() {
  fmt.Println("starting json server")
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/persons", PersonHandler)
  http.ListenAndServe(":3000", nil)
}
