//go routinを使わない場合の処理
package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
)

func sync_process() {
  fmt.Println("--- sync process ---")
  urls := []string {
    "http://example.com",
    "http://example.net",
    "http://example.org",
  }

  for _, url := range urls {
    res, err := http.Get(url)
    if err != nil {
      log.Fatal(err)
    }
    defer res.Body.Close()
    fmt.Println(url, res.Status)
  }
  fmt.Println("------")
}

func async_process() {
  fmt.Println("--- async process ---")
  urls := []string {
    "http://example.com",
    "http://example.net",
    "http://example.org",
  }

  for _, url := range urls {
    go func(url string) {
      res, err := http.Get(url)
      if err != nil {
        log.Fatal(err)
      }
      defer res.Body.Close()
      fmt.Println(url, res.Status)
    }(url)
  }
  time.Sleep(time.Second)

  fmt.Println("------")
}

func main() {
  sync_process()
  async_process()
}


