package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprint(writer, HTML_PAGE)
}

const HTML_PAGE string = `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Upload de Vídeo</title>
</head>
<body>
    <h1>Enviar Vídeo</h1>
  <form action="http://localhost:8080/video" method="post" enctype="multipart/form-data">
        <label for="video">Escolha um vídeo para enviar:</label>
        <input type="file" id="video" name="video" accept="video/*" required>
        <br><br>
        <input type="submit" value="Enviar Vídeo">
    </form>
</body>
</html>
`
