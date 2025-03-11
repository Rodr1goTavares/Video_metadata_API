package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	. "videomt/pkg"
)


func VideoFileHandler(writer http.ResponseWriter, request *http.Request) {
  if request.Method != "POST" {
    http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  request.ParseMultipartForm(10 << 20) // 10 MB

  file, _, err := request.FormFile("video")
  if err != nil {
    http.Error(writer, "Error to get file", http.StatusBadRequest)
    return
  }
  defer file.Close()

  videoMetadata, err := ExtractVideoMetadata(&file)
  if err != nil {
    http.Error(writer, "Error to extract metadata.", http.StatusInternalServerError)
    fmt.Println("Error to get metadata\n", err)
    return
  }
  
  writer.Header().Set("Content-Type", "application/json")
  writer.WriteHeader(http.StatusOK)
  
  if err := json.NewEncoder(writer).Encode(videoMetadata); err != nil {
		http.Error(writer, "Error serialize video metadata", http.StatusInternalServerError)
    fmt.Println("Error to serialize metadata\n", err)
	}
  
  fmt.Println("Video metadata:\n", videoMetadata)

}
