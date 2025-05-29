const output = document.getElementById("output-content");
const lilypond = document.getElementById("lilypond-content");
const editor = document.getElementById("editor");
const svgContainer = document.getElementById("svg-container");
const htmlOutput = document.getElementById("html-output");
const htmlRaw = document.getElementById("html-raw");
const connectionStatus = document.getElementById("connection-status");

let socket;

function setConnectionStatus(connected) {
  if (connected) {
    connectionStatus.textContent = "Connected";
    connectionStatus.style.backgroundColor = "#2e7d32"; // green
  } else {
    connectionStatus.textContent = "Disconnected";
    connectionStatus.style.backgroundColor = "#c62828"; // red
  }
}

function connect() {
  socket = new WebSocket("ws://localhost:8080/ws");

  socket.onopen = () => {
    console.log("✅ WebSocket connected");
    setConnectionStatus(true);
    sendNotation();
  };

  socket.onmessage = (event) => {
    let data;
    try {
      data = JSON.parse(event.data);
    } catch (err) {
      console.error("❌ JSON parse error:", event.data);
      return;
    }

    if (data.tree) output.textContent = data.tree;
    if (data.lilypond) lilypond.textContent = data.lilypond;
    if (data.svg) svgContainer.innerHTML = data.svg;
    if (data.html) htmlOutput.innerHTML = data.html;
    if (data.html) htmlRaw.innerText = formatHTML(data.html);
  };

  socket.onerror = (error) => {
    console.error("WebSocket error:", error);
  };

  socket.onclose = () => {
    console.warn("⚠️ WebSocket closed. Reconnecting in 3s...");
    setConnectionStatus(false);
    setTimeout(connect, 3000);
  };
}

function sendNotation() {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(editor.value);
  }
}

// Debounced user input
let debounceTimer;
editor.addEventListener("input", () => {
  clearTimeout(debounceTimer);
  debounceTimer = setTimeout(() => {
    sendNotation();
  }, 200);
});

function formatHTML(html) {
  const div = document.createElement("div");
  div.innerHTML = html;
  return div.innerHTML;
}


connect();

