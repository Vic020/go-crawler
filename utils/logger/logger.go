package logger

import (
	"log"
	"os"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	debugLogger = log.New(os.Stdout, "[DEBUG]", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(os.Stdout, "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
}

func Debug(v ...interface{}) {
	debugLogger.Println(v...)
}

func Info(v ...interface{}) {
	infoLogger.Println(v...)
}

func Error(v ...interface{}) {
	errorLogger.Println(v...)
}
