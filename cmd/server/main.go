package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"textual_music_notation/internal/logger"
	"textual_music_notation/pkg/parser"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleMessage(msg string) map[string]string {
	parsed := parser.ParseComposition(msg)
	formatter := &parser.StringFormatter{}
	parser.FormatComposition(parsed, formatter)
	treeOutput := formatter.Builder.String()
	htmlOutput := parser.CompositionToHTML(parsed)

	logger.Log("DEBUG", "Tree Output:\n%s", treeOutput)
	logger.Log("DEBUG", "HTML Output:\n%s", htmlOutput)

	return map[string]string{
		"tree": treeOutput,
		"html": htmlOutput,
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	logger.Log("DEBUG", "WebSocket request received, upgrading the connection...")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	logger.Log("DEBUG", "Client Connected via WebSocket")

	for {
		log.Print("Waiting for message from client...\n\n\n\n")
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket Read Error:", err)
			break
		}
		log.Printf("Message received from client: %s\n", string(msg))

		response := handleMessage(string(msg))

		if err := conn.WriteJSON(response); err != nil {
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
		webDir = "web"
	}

	logger.Log("DEBUG", "Serving from: %s", webDir)
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleWebSocket)
	logger.Log("INFO", "Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func runSelfTest() {
	//	sample := "S-- r- g | m- P D- N\n.\nhe-llo |"
	sample := "S "
	logger.Log("DEBUG", "🧪 Running self-test with input: %s", sample)
	handleMessage(sample)
	logger.Log("DEBUG", "------------  end self test ----\n\n\n")
}

func main() {
	logger.InitLogger()
	runSelfTest()
	serveFiles()
}
