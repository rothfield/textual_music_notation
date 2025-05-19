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

    Log("DEBUG", "Client Connected via WebSocket")

    for {
        log.Println("Waiting for message from client...")
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("WebSocket Read Error:", err)
            break
        }
        log.Printf("Message received from client: %s\n", string(msg))

        // Parse the composition
        parsedOutput := ParseComposition(string(msg))
        log.Println("Parsed composition successfully.")

        // Display the parse tree with the new format
        DisplayCompositionTree(parsedOutput)

        // Generate the formatted tree as a string
        formattedTree := GenerateFormattedTree(parsedOutput)

        // Send the formatted string instead of JSON
        err = conn.WriteMessage(websocket.TextMessage, []byte(formattedTree))
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
    InitLogger()
    defer logFile.Close()
    serveFiles()
}

// DisplayCompositionTree prints out the structure of the Composition
func DisplayCompositionTree(composition *Composition) {
    log.Println("Composition")
    for i, paragraph := range composition.Paragraphs {
        log.Printf("  Paragraph %d\n", i+1)
        DisplayParagraphTree(paragraph, "    ")
    }
}

// DisplayParagraphTree prints out the structure of a Paragraph
func DisplayParagraphTree(paragraph Paragraph, indent string) {
    log.Println(indent + "LetterLine")
    if paragraph.LetterLine != nil {
        for _, element := range paragraph.LetterLine.Elements {
            if element.IsBeat {
                log.Println(indent + "  - Beat:")
                for _, subElement := range element.SubElements {
                    log.Printf(indent + "    - %s: %s [Column=%d]\n", subElement.Token.Type, subElement.Token.Value, subElement.Column)
                }
            } else {
                log.Printf(indent + "  - %s: %s [Column=%d]\n", element.Token.Type, element.Token.Value, element.Column)
            }
        }
    }
}


