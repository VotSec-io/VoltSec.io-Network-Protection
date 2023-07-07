package firewall

import (
	"log"
	"os"
)

// Logger represents a logging utility.
type Logger struct {
	file *os.File
}

// NewLogger creates a new instance of Logger.
func NewLogger(logFile string) (*Logger, error) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &Logger{
		file: file,
	}, nil
}

// Log writes a log entry to the log file.
func (logger *Logger) Log(entry string) {
	log.SetOutput(logger.file)
	log.Println(entry)
}

// Close closes the log file.
func (logger *Logger) Close() error {
	return logger.file.Close()
}
