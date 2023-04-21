package logger

import (
	"fmt"
	"github.com/gookit/color"
	"io"
	"os"
	"sync"
	"time"
)

var (
	InitDate string
	LogChan  = make(chan string)
	LogFile  *os.File
)

func init() {
	InitDate = time.Now().Format("2006-01-02")
	if !pathExists("logs/" + InitDate) {
		mkdirPath("logs/" + InitDate)
	}
	LogFile, _ = os.OpenFile(fmt.Sprintf("logs/%v/dapp.log", InitDate), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	go func() {
		for {
			select {
			case msg := <-LogChan:
				currentDate := time.Now().Format("2006-01-02")
				if currentDate != InitDate {
					rw := &sync.RWMutex{}
					rw.Lock()
					err := LogFile.Close()
					if err != nil {
						color.Danger.Println("Log file close error, " + err.Error())
						return
					}
					InitDate = currentDate
					LogFile, _ = os.OpenFile(fmt.Sprintf("logs/%v/dapp.log", InitDate), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
					rw.Unlock()
				}
				currentTime := time.Now().Format("2006-01-02 15:04:05")
				_, _ = io.WriteString(LogFile, fmt.Sprintf("[%v] - %v\n", currentTime, msg))
			}
		}
	}()
}

// Info Logger  info
func Info(message string, args ...any) {
	msg := fmt.Sprintf(message, args...)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	color.RGB(57, 121, 227).Println(fmt.Sprintf("[%v] - %v", currentTime, msg))
	LogChan <- msg
}

// Error Logger error
func Error(message string, args ...any) {
	msg := fmt.Sprintf(message, args...)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	color.Danger.Println(fmt.Sprintf("[%v] - %v", currentTime, msg))
	LogChan <- msg
}

// Warn Logger warning
func Warn(message string, args ...any) {
	msg := fmt.Sprintf(message, args...)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	color.Warn.Println(fmt.Sprintf("[%v] - %v", currentTime, msg))
	LogChan <- msg
}

// Success Logger success
func Success(message string, args ...any) {
	msg := fmt.Sprintf(message, args...)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	color.RGB(39, 229, 120).Println(fmt.Sprintf("[%v] - %v", currentTime, msg))
	LogChan <- msg
}

func mkdirPath(dir string) bool {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		color.Danger.Println("Create log folder, " + err.Error())
		return false
	}
	return true
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
