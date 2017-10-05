package main

import (
    "fmt"
    "net/http"
    "time"
    "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
  secs := time.Now().Unix()
  urlPathString := r.URL.Path[1:]
  urlPathInt, err := strconv.ParseInt(urlPathString, 10, 0)

  if err == nil{
    duration := time.Duration(urlPathInt)*time.Second
    time.Sleep(duration)
    fmt.Fprintf(w, "app resp time: %ds", time.Now().Unix() - secs)
  }else{
    fmt.Fprintf(w, "you didn't specify an int as the pause value, silly..")
  }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
