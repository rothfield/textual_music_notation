package main

import (
    "fmt" 
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

        // Generate the formatted tree including raw text
        //formattedTree := GenerateFormattedTree(parsedOutput)
        

        // Display the parse tree with the new format
        DisplayCompositionTree(parsedOutput)

        // ✅ Send the formatted string with raw text to the client
        err = conn.WriteMessage(websocket.TextMessage, []byte("hi"))
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

// DisplayCompositionTree prints out the structure of the Composition
func DisplayCompositionTree(composition *Composition) {
    fmt.Println("Composition")
    fmt.Println("  Raw Text:")
    fmt.Println("  " + composition.RawText)
    
    for i, paragraph := range composition.Paragraphs {
        fmt.Printf("  Paragraph %d\n", i+1)
        DisplayParagraphTree(paragraph, "    ")
    }
}

// DisplayParagraphTree prints out the structure of a Paragraph
func DisplayParagraphTree(paragraph Paragraph, indent string) {
    fmt.Println(indent + "LetterLine")
    if paragraph.LetterLine != nil {
        for _, element := range paragraph.LetterLine.Elements {
            if element.IsBeat {
                fmt.Println(indent + "  - Beat:")
                for _, subElement := range element.SubElements {
                    fmt.Printf(indent + "    - %s: %s [Column=%d]\n", subElement.Token.Type, subElement.Token.Value, subElement.Column)
                }
            } else {
                fmt.Printf(indent + "  - %s: %s [Column=%d]\n", element.Token.Type, element.Token.Value, element.Column)
            }
        }
    }
}


func main() {
    InitLogger()
    defer logFile.Close()
    serveFiles()
}

