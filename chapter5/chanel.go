package main

import (
  "fmt"
  "log"
  "net/http"
)

func async_process() {
  fmt.Println("--- async process ---")
  urls := []string {
    "http://example.com",
    "http://example.net",
    "http://example.org",
  }
  statusChan := make(chan string)

  for _, url := range urls {
    // 取得処理をゴルーチンで実行する
    go func(url string) {
      res, err := http.Get(url)
      if err != nil {
        log.Fatal(err)
      }
      defer res.Body.Close()
      statusChan <- res.Status
    }(url)
  }

  for i := 0; i < len(urls); i++ {
    fmt.Println(<-statusChan)
  }

  fmt.Println("------")
}

func main() {
  async_process()
}


