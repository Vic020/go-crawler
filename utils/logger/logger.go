package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/natefinch/lumberjack"
)

const (
	LogErrorFileName = "go-crawler.error.log"
	LogStdFileName   = "go-crawler.log"

	DebugPrefix = "[DEBUG] "
	InfoPrefix  = "[INFO] "
	ErrorPrefix = "[ERROR]"
)

var (
	DebugMode = false
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	debugLogger = log.New(os.Stdout, DebugPrefix, log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(os.Stdout, InfoPrefix, log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, ErrorPrefix, log.Ldate|log.Ltime|log.Lshortfile)
}

func initLoggerHelper(logPath, name, prefix string) *log.Logger {
	fileName := path.Join(logPath, name)

	// f, err := os.OpenFile(path.Join(logPath, name),
	// 	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Printf("error opening file: %v", err)
	// 	os.Exit(1)
	// }
	// f.Close()

	return log.New(io.MultiWriter(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    500,
		MaxAge:     30,
		MaxBackups: 0,
	}, os.Stdout), prefix, log.Ldate|log.Ltime|log.Lshortfile)

}

type LoggerOptions struct {
	LogPath   string
	DebugMode bool
}

func InitLogger(ops LoggerOptions) {

	DebugMode = ops.DebugMode
	if DebugMode {
		Debug("Debug Enable")
	}

	infoLogger = initLoggerHelper(ops.LogPath, LogStdFileName, InfoPrefix)
	errorLogger = initLoggerHelper(ops.LogPath, LogErrorFileName, ErrorPrefix)

}

func Debug(v ...interface{}) {
	if DebugMode {
		debugLogger.Output(2, fmt.Sprintln(v...))
	}
}

func Debugf(format string, v ...interface{}) {
	if DebugMode {
		debugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Info(v ...interface{}) {
	infoLogger.Output(2, fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	infoLogger.Output(2, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	errorLogger.Output(2, fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	errorLogger.Output(2, fmt.Sprintf(format, v...))
}
