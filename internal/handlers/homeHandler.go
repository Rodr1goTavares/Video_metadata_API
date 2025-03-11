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
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Upload de Vídeo</title>
</head>
<body>
    <h2>Upload de Vídeo</h2>
    <form id="uploadForm" enctype="multipart/form-data">
        <input type="file" id="videoInput" name="video" accept="video/mp4" required>
        <button type="submit">Enviar</button>
    </form>
    <h3>Resposta do Servidor:</h3>
    <pre id="response"></pre>

    <script>
        document.getElementById("uploadForm").addEventListener("submit", async function(event) {
            event.preventDefault();
            const formData = new FormData();
            const fileInput = document.getElementById("videoInput");
            
            if (fileInput.files.length === 0) {
                alert("Selecione um arquivo antes de enviar.");
                return;
            }
            
            formData.append("video", fileInput.files[0]);

            try {
                const response = await fetch("/video", {
                    method: "POST",
                    body: formData
                });

                if (!response.ok) {
                    throw new Error("Erro ao enviar vídeo");
                }

                const jsonResponse = await response.json();
                document.getElementById("response").textContent = JSON.stringify(jsonResponse, null, 2);
            } catch (error) {
                document.getElementById("response").textContent = "Erro: " + error.message;
            }
        });
    </script>
</body>
</html>
`
