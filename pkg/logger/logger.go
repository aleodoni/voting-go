// Package logger provides logging functionality for the application.
package logger

import "log"

func Info(message string) {
	log.Printf("[INFO] %s", message)
}

func Error(message string) {
	log.Printf("[ERROR] %s", message)
}

func Warning(message string) {
	log.Printf("[WARNING] %s", message)
}

func Debug(message string) {
	log.Printf("[DEBUG] %s", message)
}
