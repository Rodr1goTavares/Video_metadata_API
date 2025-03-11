# Video Metadata API Documentation

## Overview

The Video Metadata API allows users to upload video files and retrieve metadata information about the uploaded video, such as file name, size, dimensions, aspect ratio, and duration.

## API Endpoints

### 1. Upload Video

- **Endpoint:** `/video/upload`
- **Method:** `POST`
- **Content-Type:** `multipart/form-data`
- **Request Body:** A file input named `video` containing the video file.

#### Request Example using `curl`

```bash
curl -X POST http://localhost:8080/video/upload \
  -F "video=@/path/to/your/video.mp4"

Response

    Status Code: 200 OK
    Content-Type: application/json
    Response Body:

´´´json
{
  "fileName": "video.mp4",
  "size": "10MB",
  "height": 1080,
  "width": 1920,
  "aspectRatio": "16:9",
  "duration": "120" // <-- secounds
}
´´´
