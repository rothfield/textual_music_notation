
let socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = function() {
    console.log("Connected to WebSocket server");
};

socket.onmessage = function(event) {
    const sections = event.data.split("=== Staff Notation (SVG) ===");
    const treeAndLily = sections[0];
    const svgData = sections[1] ?? "";

    const parts = treeAndLily.split("=== LilyPond Source ===");
    document.getElementById("output-content").innerText = parts[0];
    document.getElementById("lilypond-content").innerText = parts[1];

    const svgContainer = document.getElementById("svg-container");
    svgContainer.innerHTML = svgData;
};

socket.onerror = function(error) {
    console.error("WebSocket error:", error);
};

let debounceTimeout;
document.getElementById("notation-input").addEventListener("input", (event) => {
    clearTimeout(debounceTimeout);
    debounceTimeout = setTimeout(() => {
        socket.send(event.target.value);
    }, 300);
});
    