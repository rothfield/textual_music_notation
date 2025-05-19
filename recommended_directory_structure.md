project-root/
│
├── web/                         # Front-end assets (HTML, CSS, JS)
│   ├── main.js                  
│   └── style.css                
│
├── cmd/                         # Main application entry points
│   └── server/                  # Server application
│       └── main.go              # Main server logic
│
├── internal/                    # Internal packages, not exposed externally
│   ├── parser/                  # Parsing logic for composition and paragraphs
│   │   ├── composition_parser.go
│   │   ├── paragraph_parser.go
│   │   ├── letter_line_parser.go
│   │   └── classify_lines.go
│   └── lexer/                   # Lexers for parsing
│       ├── upper_annotation_lexer.go
│       ├── lower_annotation_lexer.go
│       └── lyrics_lexer.go
│
├── pkg/                         # Shared packages across the project
│   ├── formatter.go
│   └── tree_renderer.go
│
├── scripts/                     # Utility scripts
│   └── bump.go                  # ✅ Version bumper script
│
├── assets/                      # Static assets like HTML, CSS, JS
│   ├── index.html
│   └── web/                     # Served by the server
│       ├── main.js
│       └── style.css
│
├── go.mod                       # Go module definition
├── go.sum                       # Go module dependencies
└── README.md                    # Project documentation

