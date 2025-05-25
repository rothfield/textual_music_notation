package logger 

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	FatalLogger *log.Logger
	logFile     *os.File
)

func InitLogger() {
	var err error
	logFile, err = os.OpenFile("log/application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Could not open log file: ", err)
	}

	// Create loggers
	DebugLogger = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(logFile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(logFile, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Log(level string, message string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	interpolatedMessage := fmt.Sprintf(message, args...)
	logEntry := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, interpolatedMessage)
	fmt.Print(logEntry)            // Console
	_, _ = logFile.WriteString(logEntry) // File
}

