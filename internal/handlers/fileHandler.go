package handlers

import (
	"fmt"
	"net/http"
)


func VideoFileHandler(writer http.ResponseWriter, request *http.Request) {
  if request.Method != "POST" {
    http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  request.ParseMultipartForm(10 << 20) // 10 MB

  file, handler, err := request.FormFile("video")
  if err != nil {
    http.Error(writer, "Error to get file", http.StatusBadRequest)
    return
  }
  defer file.Close()

  fmt.Printf("File name: %s\n", handler.Filename)
}
