const socketUrl = "ws://localhost:8080/ws";
let socket;
let reconnectAttempts = 0;
let lastMessage = "";
let debounceTimeout;

function connectWebSocket() {
    console.log("Connecting to WebSocket...");
    socket = new WebSocket(socketUrl);

    socket.onopen = function() {
        console.log("Connected to WebSocket server");
        reconnectAttempts = 0;
        document.getElementById("connection-status").innerText = "Connected";
        document.getElementById("connection-status").style.backgroundColor = "green";

        // ✅ Resend the last message if there was one
        if (lastMessage) {
            console.log("Resending last message...");
            socket.send(lastMessage);

            // Log to the console
            console.log(`Request re-sent: ${lastMessage}`);
        }
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

    socket.onclose = function(event) {
        console.warn("WebSocket closed. Reason:", event.reason);
        document.getElementById("connection-status").innerText = "Disconnected";
        document.getElementById("connection-status").style.backgroundColor = "red";

        const maxReconnectAttempts = 10;
        if (reconnectAttempts < maxReconnectAttempts) {
            const reconnectDelay = Math.min(1000 * Math.pow(2, reconnectAttempts), 30000);
            console.log(`Attempting to reconnect in ${reconnectDelay / 1000} seconds...`);
            setTimeout(() => {
                reconnectAttempts++;
                connectWebSocket();
            }, reconnectDelay);
        } else {
            console.error("Max reconnect attempts reached.");
        }
    };

    socket.onerror = function(error) {
        console.error("WebSocket error:", error);
    };
}

// Initial connection
connectWebSocket();

// Handle network events
window.addEventListener("offline", () => {
    console.warn("You are offline. WebSocket will try to reconnect when you are back online.");
    document.getElementById("connection-status").innerText = "Offline";
    document.getElementById("connection-status").style.backgroundColor = "orange";
});

window.addEventListener("online", () => {
    console.log("You are back online. Attempting WebSocket reconnection...");
    connectWebSocket();
});

document.getElementById("notation-input").addEventListener("input", (event) => {
    clearTimeout(debounceTimeout);
    debounceTimeout = setTimeout(() => {
        lastMessage = event.target.value;
        if (socket.readyState === WebSocket.OPEN) {
            socket.send(lastMessage);
        }
    }, 300);
});

