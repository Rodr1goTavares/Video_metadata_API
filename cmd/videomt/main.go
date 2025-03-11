package main

import (
	"fmt"
	"net/http"
	. "videomt/internal/handlers"
)

func main() {
  startServer()
}

func startServer() {
  http.HandleFunc("/", HomeHandler)
  http.HandleFunc("/video/upload", VideoFileHandler)
  PORT := "8080"
  fmt.Println("Server runnin on", PORT)
  if err := http.ListenAndServe(":" + PORT, nil); err != nil {
    fmt.Println("Error to run server: \n", err)
    return
  }
}
