package utils

import (
	"log"
	"os"
)

// NewLogger initializes and returns a custom logger
func NewLogger(serviceName string) *log.Logger {
	return log.New(os.Stdout, "["+serviceName+"] ", log.LstdFlags|log.Lshortfile)
}
