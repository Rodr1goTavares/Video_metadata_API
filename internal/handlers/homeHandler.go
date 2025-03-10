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
    
    <form id="uploadForm" enctype="multipart/form-data">
        <label for="video">Escolha um vídeo para enviar:</label>
        <input type="file" id="video" name="video" accept="video/*" required>
        <br><br>
        <input type="submit" value="Enviar Vídeo">
    </form>

    <p id="mensagem"></p>

    <script>
        document.getElementById("uploadForm").addEventListener("submit", function(event) {
            event.preventDefault(); // Impede o redirecionamento

            const formData = new FormData();
            const fileInput = document.getElementById("video");

            if (fileInput.files.length === 0) {
                document.getElementById("mensagem").innerText = "Selecione um vídeo antes de enviar.";
                return;
            }

            formData.append("video", fileInput.files[0]);

            fetch("http://localhost:8080/video", {
                method: "POST",
                body: formData
            })
            .then(response => response.text())
            .then(data => {
                document.getElementById("mensagem").innerText = "Vídeo enviado com sucesso!";
            })
            .catch(error => {
                document.getElementById("mensagem").innerText = "Erro ao enviar o vídeo.";
                console.error("Erro:", error);
            });
        });
    </script>
</body>
</html>
`
