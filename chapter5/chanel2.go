package main

import (
  "fmt"
  "log"
  "net/http"
)

func getStatus(urls []string) <-chan string {
  // 関数でチャネルを生成
  statusChan := make(chan string)
  for _, url := range urls {
    go func(url string) {
      res, err := http.Get(url)
      if err != nil {
        log.Fatal(err)
      }
      defer res.Body.Close()
      statusChan <- res.Status
    }(url)
  }
  return statusChan   // チャネルを返す

}

func async_process() {
  fmt.Println("--- async process ---")
  urls := []string {
    "http://example.com",
    "http://example.net",
    "http://example.org",
  }
  statusChan := getStatus(urls)

  for i := 0; i < len(urls); i++ {
    fmt.Println(<-statusChan)
  }

  fmt.Println("------")
}

func main() {
  async_process()
}


