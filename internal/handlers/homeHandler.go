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
  <title>AI SaaS Startup</title>
  <script src="https://cdn.tailwindcss.com"></script>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css"></link>
  <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap" rel="stylesheet">
  <style>
    body {
      font-family: 'Inter', sans-serif;
    }
  </style>
</head>
<body class="bg-gray-50">
  <!-- Header -->
  <header class="bg-white shadow">
    <div class="container mx-auto px-6 py-4">
      <div class="flex justify-between items-center">
        <div class="text-2xl font-bold text-gray-800">AI SaaS</div>
        <nav class="space-x-4">
          <a href="#" class="text-gray-700 hover:text-gray-900">Home</a>
          <a href="#" class="text-gray-700 hover:text-gray-900">Features</a>
          <a href="#" class="text-gray-700 hover:text-gray-900">Pricing</a>
          <a href="#" class="text-gray-700 hover:text-gray-900">Contact</a>
        </nav>
      </div>
    </div>
  </header>

  <!-- Hero Section -->
  <section class="bg-gradient-to-r from-blue-500 to-indigo-600 text-white py-20">
    <div class="container mx-auto px-6 text-center">
      <h1 class="text-5xl font-bold mb-4">Revolutionize Your Business with AI</h1>
      <p class="text-xl mb-8">Unlock the power of artificial intelligence to drive growth and efficiency.</p>
      <a href="#" class="bg-white text-blue-600 px-8 py-3 rounded-full font-semibold hover:bg-gray-100">Get Started</a>
    </div>
  </section>

  <!-- Features Section -->
  <section class="py-16">
    <div class="container mx-auto px-6">
      <h2 class="text-3xl font-bold text-center mb-12">Features</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="bg-white p-6 rounded-lg shadow-lg text-center">
          <i class="fas fa-brain text-4xl text-blue-500 mb-4"></i>
          <h3 class="text-xl font-semibold mb-2">AI-Powered Insights</h3>
          <p class="text-gray-600">Get actionable insights from your data with advanced AI algorithms.</p>
        </div>
        <div class="bg-white p-6 rounded-lg shadow-lg text-center">
          <i class="fas fa-cogs text-4xl text-blue-500 mb-4"></i>
          <h3 class="text-xl font-semibold mb-2">Automation</h3>
          <p class="text-gray-600">Automate repetitive tasks and focus on what matters most.</p>
        </div>
        <div class="bg-white p-6 rounded-lg shadow-lg text-center">
          <i class="fas fa-chart-line text-4xl text-blue-500 mb-4"></i>
          <h3 class="text-xl font-semibold mb-2">Real-Time Analytics</h3>
          <p class="text-gray-600">Monitor your business performance in real-time with our analytics dashboard.</p>
        </div>
      </div>
    </div>
  </section>

  <!-- Video Upload Section -->
  <section class="bg-gray-100 py-16">
    <div class="container mx-auto px-6">
      <h2 class="text-3xl font-bold text-center mb-8">Upload Your Video</h2>
      <form id="uploadForm" action="/video/upload" method="POST" enctype="multipart/form-data" class="max-w-lg mx-auto bg-white p-8 rounded-lg shadow-lg">
        <div class="mb-6">
          <label for="video" class="block text-gray-700 font-semibold mb-2">Select a video file:</label>
          <input type="file" id="video" name="video" accept="video/*" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" required>
        </div>
        <button type="submit" class="w-full bg-blue-500 text-white px-6 py-3 rounded-lg font-semibold hover:bg-blue-600">Upload Video</button>
      </form>
      <div id="message" class="mt-4 text-center text-lg font-medium" style="display: none;"></div>
      <pre id="response" class="mt-2 p-3 bg-gray-200 rounded text-left text-sm" style="display: none;"></pre>
    </div>
  </section>


  <!-- Pricing Section -->
  <section class="py-16">
    <div class="container mx-auto px-6">
      <h2 class="text-3xl font-bold text-center mb-12">Pricing</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="bg-white p-6 rounded-lg shadow-lg text-center">
          <h3 class="text-xl font-semibold mb-4">Starter</h3>
          <p class="text-gray-600 mb-4">Perfect for small businesses</p>
          <p class="text-3xl font-bold mb-4">$29<span class="text-lg text-gray-500">/mo</span></p>
          <a href="#" class="bg-blue-500 text-white px-8 py-3 rounded-full font-semibold hover:bg-blue-600">Get Started</a>
        </div>
        <div class="bg-white p-6 rounded-lg shadow-lg text-center">
          <h3 class="text-xl font-semibold mb-4">Professional</h3>
          <p class="text-gray-600 mb-4">Ideal for growing businesses</p>
          <p class="text-3xl font-bold mb-4">$99<span class="text-lg text-gray-500">/mo</span></p>
          <a href="#" class="bg-blue-500 text-white px-8 py-3 rounded-full font-semi-bold hover:bg-blue-600">Get Started</a>
        </div>
        <div class="bg-white p-6 rounded-lg shadow-lg text-center">
          <h3 class="text-xl font-semibold mb-4">Enterprise</h3>
          <p class="text-gray-600 mb-4">For large organizations</p>
          <p class="text-3xl font-bold mb-4">$299<span class="text-lg text-gray-500">/mo</span></p>
          <a href="#" class="bg-blue-500 text-white px-8 py-3 rounded-full font-semibold hover:bg-blue-600">Get Started</a>
        </div>
      </div>
    </div>
  </section>

  <!-- Footer -->
  <footer class="bg-gray-800 text-white py-8">
    <div class="container mx-auto px-6">
      <div class="flex justify-between items-center">
        <div class="text-lg font-bold">AI SaaS</div>
        <div class="flex space-x-4">
          <a href="#" class="text-gray-400 hover:text-white"><i class="fab fa-twitter"></i></a>
          <a href="#" class="text-gray-400 hover:text-white"><i class="fab fa-linkedin"></i></a>
          <a href="#" class="text-gray-400 hover:text-white"><i class="fab fa-github"></i></a>
        </div>
      </div>
      <div class="mt-4 text-center text-gray-400">
        &copy; 2023 AI SaaS. All rights reserved.
      </div>
    </div>
  </footer>

  <script>
    document.getElementById("uploadForm").addEventListener("submit", async function(event) {
      event.preventDefault();
      const formData = new FormData();
      const fileInput = document.getElementById("video");

      if (fileInput.files.length === 0) {
        alert("Please select a video file before uploading.");
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
          throw new Error("Error uploading video");
        }

        const jsonResponse = await response.json();

        document.getElementById("message").style.display = 'block';
        document.getElementById("response").style.display = 'block';

        document.getElementById("message").textContent = "Success!";
        document.getElementById("response").textContent = JSON.stringify(jsonResponse, null, 2);
      } catch (error) {
        document.getElementById("message").style.display = 'block';
        document.getElementById("response").style.display = 'block';

        document.getElementById("message").textContent = "Error";
        document.getElementById("response").textContent = "Error: " + error.message;
      }
    });
  </script>
</body>
</html>
`
