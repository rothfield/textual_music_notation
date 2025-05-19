
# 🎵 Letter Music Notation Parser
A Go-based parser for a structured form of letter music notation that converts the notation into a parse tree, generates LilyPond source, and displays SVG renderings of the music score.

---

## 📌 Project Structure
```
project-root/
│
├── web/                         # Front-end assets
│   ├── main.js                  # WebSocket communication and UI updates
│   └── style.css                # Front-end styling
│
├── composition_parser.go        # Parses the entire composition into paragraphs
├── paragraph_parser.go          # Parses individual paragraphs, lines, and annotations
├── classify_lines.go            # Identifies the line types (LetterLine, Annotations, Lyrics)
├── letter_line_parser.go        # Parses individual letter lines into elements
├── upper_annotation_lexer.go    # Lexer for upper annotations
├── lower_annotation_lexer.go    # Lexer for lower annotations
├── lyrics_lexer.go              # Lexer for lyrics lines
├── tree_renderer.go             # Generates a formatted tree display of the parse structure
├── formatter.go                 # Handles formatted output for display and WebSocket
├── server.go                    # WebSocket server and HTTP file server
├── index.html                   # Main HTML interface
├── bump.go                      # Script to bump version of JS and CSS in index.html
│
├── go.mod                       # Go module definition
└── go.sum                       # Go module dependencies
```

---

## 📌 Specifications
### 🎯 Notation Structure
The notation consists of:
- **LetterLine** — Main musical line containing:
  - Pitches (`S, r, R, g, G, m, d, D, n, N, P`)
  - Dashes (`-`)
  - Barlines (`|`)
- **Upper Annotation Line** — Contains:
  - Octave markers (`.` for higher and `:` for highest)
  - Talas (`0 + 1 2 3 4 5 6 7 8`)
- **Lower Annotation Line** — Contains:
  - Octave markers (`.` for lower and `:` for lowest)
- **Lyrics Line** — Contains:
  - Only lyrics text

---

### 🎯 Paragraph Structure
A **Paragraph** consists of:
- **Upper Annotations** (optional, before the LetterLine)
- **LetterLine** (mandatory)
- **Lower Annotations** (optional, immediately after the LetterLine)
- **Lyrics Line** (optional, after Lower Annotations)

---

### 🎯 Parsing Logic
1. **Trim Leading and Trailing Empty Lines:**  
   - Empty lines at the start and end are stripped before parsing.

2. **Fold Consecutive Empty Lines:**  
   - Multiple empty lines are collapsed into a single one.

3. **No Empty Paragraphs Created:**  
   - Only non-empty paragraphs are appended to the `Composition`.

4. **LetterLine Identification:**  
   - If no `|` is found, the longest line with only valid pitches (`S, r, R, g, G, m, d, D, n, N, P`) is chosen.
   - If only one line is present, it is automatically classified as a `LetterLine`.

5. **Order of Classification:**  
   - Upper Annotations → LetterLine → Lower Annotations → Lyrics

---

## 📌 Setup Instructions
1. **Clone the repository**:
    ```sh
    git clone <repository-url>
    cd <repository-name>
    ```

2. **Install dependencies**:
    ```sh
    go mod tidy
    ```

3. **Prepare the directory structure** and run the server:
    ```sh
    go run server.go
    ```

4. **Visit the interface**:
    - Open [http://localhost:8080](http://localhost:8080) in your browser.
