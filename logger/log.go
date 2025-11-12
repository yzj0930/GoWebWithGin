package logger

import (
	"log"
	"os"

	"github.com/yzj0930/GoWebWithGin/config"
)

type LoggerLevel int

const (
	LevelTrace LoggerLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

type LevelLogger struct {
	level LoggerLevel
	*log.Logger
}

var Logger *LevelLogger

func InitSysetmLogger() {
	logPath := config.GlobalConfig.Logging.File
	level := config.GlobalConfig.Logging.Level
	Logger = InitLogger(logPath, level, "SYSTEM: ")
}

func InitLogger(logPath string, level string, prefix string) *LevelLogger {
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	logger := log.New(file, prefix, log.LstdFlags|log.Lshortfile)
	newLogger := &LevelLogger{
		level:  LevelInfo, // 这里可以根据配置文件设置日志级别
		Logger: logger,
	}
	switch level {
	case "trace":
		newLogger.level = LevelTrace
	case "debug":
		newLogger.level = LevelDebug
	case "info":
		newLogger.level = LevelInfo
	case "warn":
		newLogger.level = LevelWarn
	case "error":
		newLogger.level = LevelError
	case "fatal":
		newLogger.level = LevelFatal
	default:
		newLogger.level = LevelInfo
	}
	return newLogger
}

func (l *LevelLogger) Info(v ...interface{}) {
	if l.level <= LevelInfo {
		l.Logger.SetPrefix("INFO: ")
		l.Logger.Println(v...)
	}
}

func (l *LevelLogger) Error(v ...interface{}) {
	if l.level <= LevelError {
		l.Logger.SetPrefix("ERROR: ")
		l.Logger.Println(v...)
	}
}

func (l *LevelLogger) Debug(v ...interface{}) {
	if l.level <= LevelDebug {
		l.Logger.SetPrefix("DEBUG: ")
		l.Logger.Println(v...)
	}
}

func (l *LevelLogger) Warn(v ...interface{}) {
	if l.level <= LevelWarn {
		l.Logger.SetPrefix("WARN: ")
		l.Logger.Println(v...)
	}
}

func (l *LevelLogger) Fatal(v ...interface{}) {
	if l.level <= LevelFatal {
		l.Logger.SetPrefix("FATAL: ")
		l.Logger.Println(v...)
		os.Exit(1)
	}
}

func (l *LevelLogger) Trace(v ...interface{}) {
	if l.level <= LevelTrace {
		l.Logger.SetPrefix("TRACE: ")
		l.Logger.Println(v...)
	}
}

func SetLogLevel(level LoggerLevel) {
	if Logger != nil {
		Logger.level = level
	}
}

func Info(v ...interface{}) {
	if Logger != nil {
		Logger.Info(v...)
	}
}

func Error(v ...interface{}) {
	if Logger != nil {
		Logger.Error(v...)
	}
}

func Debug(v ...interface{}) {
	if Logger != nil {
		Logger.Debug(v...)
	}
}

func Warn(v ...interface{}) {
	if Logger != nil {
		Logger.Warn(v...)
	}
}

func Fatal(v ...interface{}) {
	if Logger != nil {
		Logger.Fatal(v...)
	}
}

func Trace(v ...interface{}) {
	if Logger != nil {
		Logger.Trace(v...)
	}
}

func SetLevel(level LoggerLevel) {
	if Logger != nil {
		Logger.level = level
	}
}
