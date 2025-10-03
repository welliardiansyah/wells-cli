package scaffold

import (
	"fmt"
	"time"
)

func Info(msg string, args ...interface{}) {
	fmt.Printf("[INFO] %s - %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(msg, args...))
}

func Success(msg string, args ...interface{}) {
	fmt.Printf("[SUCCESS] %s - %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(msg, args...))
}

func Warning(msg string, args ...interface{}) {
	fmt.Printf("[WARNING] %s - %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(msg, args...))
}

func Error(msg string, err error) {
	fmt.Printf("[ERROR] %s - %s: %v\n", time.Now().Format("15:04:05"), msg, err)
}
