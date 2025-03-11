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
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-gray-100 flex items-center justify-center min-h-screen">
    <div class="bg-white p-6 rounded-lg shadow-md w-96 text-center">
        <h2 class="text-xl font-semibold mb-4">Upload de Vídeo</h2>
        <form id="uploadForm" enctype="multipart/form-data" class="space-y-4">
            <input type="file" id="videoInput" name="video" accept="video/mp4" required
                   class="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50">
            <button type="submit" class="w-full bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600">Enviar</button>
        </form>
        <h3 class="text-lg font-medium mt-4" id="message" style="display: none;"></h3>
        <pre id="response" class="mt-2 p-3 bg-gray-200 rounded text-left text-sm" style="display: none;"></pre>
    </div>

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

            document.getElementById("message").style.display = 'none';
            document.getElementById("response").style.display = 'none';

            try {
                const response = await fetch("/video/upload", {
                    method: "POST",
                    body: formData
                });

                if (!response.ok) {
                    throw new Error("Erro ao enviar vídeo");
                }

                const jsonResponse = await response.json();
                
                document.getElementById("message").style.display = 'block';
                document.getElementById("response").style.display = 'block';

                document.getElementById("message").textContent = "Result";
                document.getElementById("response").textContent = JSON.stringify(jsonResponse, null, 2);
            } catch (error) {
                document.getElementById("message").style.display = 'block';
                document.getElementById("response").style.display = 'block';

                document.getElementById("message").textContent = "Failed to read file";
                document.getElementById("response").textContent = "Erro: " + error.message;
            }
        });
    </script>
</body>
</html>
`
