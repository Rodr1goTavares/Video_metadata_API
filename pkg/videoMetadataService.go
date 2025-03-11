package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type VideoMetadata struct {
  FileName string `json: "fileName"`
  Size string `json:"size"`
  Height int16 `json:"height"`
  Width int16 `json:"width"`
  AspectRatio string `json:"aspectRatio"`
  Duration string `json:"durationSecounds"`
}




func ExtractVideoMetadata(fileHeader *multipart.FileHeader, file *multipart.File) (*VideoMetadata, error) {
	// Create temp file
	tempFile, err := os.CreateTemp("", "uploaded-*.mp4")
	if err != nil {
		fmt.Println("Failed to save file:", err)
		return nil, err
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	// Copy content to temp file
	_, err = io.Copy(tempFile, *file)
	if err != nil {
		fmt.Println("Failed to copy file:", err)
		return nil, err
	}

	tempFile.Sync()

	// get file data with ffmpeg
	cmd, err := ffmpeg_go.Probe(tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("error executing ffprobe: %v", err)
	}

	// Decode JSON
	var metadataMap map[string]interface{}
	if err := json.Unmarshal([]byte(cmd), &metadataMap); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	// Extract file size
	format, ok := metadataMap["format"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid format in metadata")
	}

	sizeStr, _ := format["size"].(string)
	sizeBytes, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse size: %v", err)
	}
	sizeMB := fmt.Sprintf("%.2f MB", float64(sizeBytes)/(1024*1024)) // Convertendo para MB

	// Extract duraction 
	durationStr, _ := format["duration"].(string)

	// Extract width and heigth
	streams, ok := metadataMap["streams"].([]interface{})
	if !ok || len(streams) == 0 {
		return nil, fmt.Errorf("no streams found in metadata")
	}

	var width, height int
	for _, stream := range streams {
		streamMap, ok := stream.(map[string]interface{})
		if !ok {
			continue
		}
		if codecType, exists := streamMap["codec_type"].(string); exists && codecType == "video" {
			w, _ := streamMap["width"].(float64)
			h, _ := streamMap["height"].(float64)
			width = int(w)
			height = int(h)
			break
		}
	}

	aspectRatio := calculateAspectRatio(width, height)

	return &VideoMetadata{
		FileName:    fileHeader.Filename,
		Size:        sizeMB,
		Height:      int16(height),
		Width:       int16(width),
		AspectRatio: aspectRatio,
		Duration:    durationStr,
	}, nil
}


func calculateAspectRatio(width int, height int) string {
	if height == 0 {
		return "Unknown"
	}
	divisor := gcd(width, height)
	return fmt.Sprintf("%d:%d", width/divisor, height/divisor)
}


func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
