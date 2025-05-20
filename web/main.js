const socketUrl = "ws://localhost:8080/ws";
let socket;
let lastMessage = "";
let debounceTimeout;

/**
 * ✅ Establishes the WebSocket connection and handles events
 */
function connectWebSocket() {
    console.log("Connecting to WebSocket...");
    socket = new WebSocket(socketUrl);

    /**
     * ✅ On successful connection, reset and re-send the last message if available
     */
    socket.onopen = function() {
        console.log("Connected to WebSocket server");
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

    /**
     * ✅ On message received, split the sections and render appropriately
     */
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

    /**
     * ✅ On connection close, attempt to reconnect every 30 seconds
     */
    socket.onclose = function(event) {
        console.warn("WebSocket closed. Reason:", event.reason);
        document.getElementById("connection-status").innerText = "Disconnected";
        document.getElementById("connection-status").style.backgroundColor = "red";

        console.log("Attempting to reconnect in 30 seconds...");
        setTimeout(() => {
            connectWebSocket();
        }, 5000); // Retry every 30 seconds
    };

    /**
     * ✅ On WebSocket error, log it to the console
     */
    socket.onerror = function(error) {
        console.error("WebSocket error:", error);
    };
}

/**
 * ✅ Handle network events (offline and online recovery)
 */
window.addEventListener("offline", () => {
    console.warn("You are offline. WebSocket will try to reconnect when you are back online.");
    document.getElementById("connection-status").innerText = "Offline";
    document.getElementById("connection-status").style.backgroundColor = "orange";
});

window.addEventListener("online", () => {
    console.log("You are back online. Attempting WebSocket reconnection...");
    connectWebSocket();
});

/**
 * ✅ Handle input changes with debouncing
 */
document.getElementById("notation-input").addEventListener("input", (event) => {
    clearTimeout(debounceTimeout);
    debounceTimeout = setTimeout(() => {
        lastMessage = event.target.value;
        if (socket.readyState === WebSocket.OPEN) {
            socket.send(lastMessage);
        }
    }, 300);
});

// ✅ Initial connection
connectWebSocket();

