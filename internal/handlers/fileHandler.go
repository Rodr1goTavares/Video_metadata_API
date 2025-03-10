package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	. "videomt/pkg"
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

  tempFile, err := os.CreateTemp("", "uploaded-*.mp4")
  if err != nil {
    http.Error(writer, "Error to proccess file", http.StatusInternalServerError)
		return
  }
  
  fmt.Printf("Temp file name: %s\n", tempFile.Name())

  _, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(writer, "Error has occurred. Try again latter.", http.StatusInternalServerError)
		return
	}

  videoMetadata, err := ExtractVideoMetadata(tempFile.Name())
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

}
