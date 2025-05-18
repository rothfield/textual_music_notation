package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    log.Println("WebSocket request received, upgrading the connection...")
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("WebSocket Upgrade Error:", err)
        return
    }
    defer conn.Close()

    log.Println("Client Connected via WebSocket")

    for {
        log.Println("Waiting for message from client...")
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("WebSocket Read Error:", err)
            break
        }
        log.Printf("Message received from client: %s\n", string(msg))

        parsedOutput := ParseNotation(string(msg))
        log.Println("Parsed notation successfully.")

        log.Println("Sending parsed output back to client...")
        err = conn.WriteMessage(websocket.TextMessage, []byte(parsedOutput))
        if err != nil {
            log.Println("WebSocket Write Error:", err)
            break
        }
        log.Println("Message sent successfully.")
    }

    log.Println("WebSocket connection closed.")
}

func serveFiles() {
    fs := http.FileServer(http.Dir("./web"))
    http.Handle("/", fs)
    http.HandleFunc("/ws", handleWebSocket)
    log.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    serveFiles()
}

