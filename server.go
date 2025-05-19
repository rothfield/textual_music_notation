package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("WebSocket Upgrade Error:", err)
        return
    }
    defer conn.Close()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("WebSocket Read Error:", err)
            break
        }

        // ✅ Parse the composition
        parsedComposition := ParseComposition(string(msg))

        // ✅ Generate the formatted tree
        formattedTree := GenerateFormattedTree(parsedComposition)

        // ✅ Console Output
        fmt.Println("=== Parse Tree Structure ===")
        fmt.Println(formattedTree)

        // ✅ WebSocket Output (plain string format)
        err = conn.WriteMessage(websocket.TextMessage, []byte(formattedTree))
        if err != nil {
            log.Println("WebSocket Write Error:", err)
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    log.Println("Server listening on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

