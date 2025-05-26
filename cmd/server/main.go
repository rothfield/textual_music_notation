package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"textual_music_notation/internal/logger"
	"textual_music_notation/pkg/parser"
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

	logger.DebugLogger.Println("Client Connected via WebSocket")

	for {
		log.Println("Waiting for message from client...")
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket Read Error:", err)
			break
		}
		log.Printf("Message received from client: %s\n", string(msg))

		parsed := parser.ParseComposition(string(msg))
		formatter := &parser.StringFormatter{}
		parser.FormatComposition(parsed, formatter)
		logger.DebugLogger.Printf("\n%s", formatter.Builder.String())

		err = conn.WriteMessage(websocket.TextMessage, []byte(formatter.Builder.String()))
		if err != nil {
			log.Println("WebSocket Write Error:", err)
			break
		}
		log.Println("Message sent successfully.")
	}

	log.Println("WebSocket connection closed.")
}

func serveFiles() {
	var webDir string
	if exePath, err := os.Executable(); err == nil {
		root := filepath.Dir(exePath)
		candidate := filepath.Join(root, "../../web")
		if _, err := os.Stat(candidate); err == nil {
			webDir = candidate
		}
	}
	if webDir == "" {
		// fallback to working directory for go run
		webDir = "web"
	}

	Log("DEBUG", "Serving from:%s", webDir)
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func xkserveFiles() {
	exePath, _ := os.Executable()
	root := filepath.Dir(exePath)
	webDir := filepath.Join(root, "../../web")
	log.Println("Serving web directory from:", webDir)
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	logger.InitLogger()
	serveFiles()
}
