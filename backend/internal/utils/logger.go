package utils

import (
    "log"
    "time"
)

func LogInfo(message string) {
    log.Printf("[INFO] %s %s", time.Now().Format("2006-01-02 15:04:05"), message)
}

func LogError(message string) {
    log.Printf("[ERROR] %s %s", time.Now().Format("2006-01-02 15:04:05"), message)
}

func LogWarning(message string) {
    log.Printf("[WARN] %s %s", time.Now().Format("2006-01-02 15:04:05"), message)
}