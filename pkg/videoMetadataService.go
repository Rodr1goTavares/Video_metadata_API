package pkg

import (
	"fmt"
	"os"
	"github.com/3d0c/gmf"
)

type VideoMetadata struct {
  FileName string `json: "fileName"`
  Size int64 `json:"size"`
  Duration float32 `json:"duration"`
  Height int16 `json:"heigth"`
  Width int16 `json:"width"`
  AspectRatio string `json:"aspectRatio"`
}

func ExtractVideoMetadata(filePath string) (*VideoMetadata, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// get file info
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Open video format with FFmpeg
	formatCtx, err := gmf.NewInputCtx(filePath)
	if err != nil {
		return nil, err
	}
	defer formatCtx.CloseInput()

	// Extract duration
	duration := float32(formatCtx.Duration()) / 1000000.0

	// Find width and length stream
	var width, height int16
  for i := 0; i < formatCtx.StreamsCnt(); i++ {
	  stream, err := formatCtx.GetStream(i)
    if err != nil {
      return nil, err
    }
	  codecCtx := stream.CodecCtx()
    if int32(codecCtx.Codec().Type()) == gmf.AVMEDIA_TYPE_VIDEO {
		  width = int16(codecCtx.Width())
		  height = int16(codecCtx.Height())
		  break
	  }
  }

	// Calc aspect ratio
	aspectRatio := fmt.Sprintf("%.2f:1", float32(width)/float32(height))

	metadata := VideoMetadata{
    FileName:    fileInfo.Name(),
		Size:        fileInfo.Size(),
		Duration:    duration,
		Height:      height,
		Width:       width,
		AspectRatio: aspectRatio,
	}

	return &metadata, nil
}
