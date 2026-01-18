package logger

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// GormLogger 将 GORM 的日志转发到项目的 Logger（写入日志文件）
type GormLogger struct {
    level gormlogger.LogLevel
}

func NewGormLogger(level gormlogger.LogLevel) gormlogger.Interface {
    return &GormLogger{level: level}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
    nl := *l
    nl.level = level
    return &nl
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
    if Logger == nil {
        return
    }
    if l.level >= gormlogger.Info {
        Logger.SetPrefix("INFO: ")
        Logger.Printf(msg, data...)
    }
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
    if Logger == nil {
        return
    }
    if l.level >= gormlogger.Warn {
        Logger.SetPrefix("WARN: ")
        Logger.Printf(msg, data...)
    }
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
    if Logger == nil {
        return
    }
    // 过滤掉 record not found
    for _, d := range data {
        if err, ok := d.(error); ok && errors.Is(err, gorm.ErrRecordNotFound) {
            return
        }
    }
    if l.level >= gormlogger.Error {
        Logger.SetPrefix("ERROR: ")
        Logger.Printf(msg, data...)
    }
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
    if Logger == nil {
        return
    }
    // 如果是 record not found，忽略
    if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
        return
    }

    elapsed := time.Since(begin)
    sql, rows := fc()
    msg := fmt.Sprintf("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)

    if err != nil {
        if l.level >= gormlogger.Error {
            Logger.SetPrefix("ERROR: ")
            Logger.Println(msg, " error:", err)
        }
        return
    }

    if l.level >= gormlogger.Info {
        Logger.SetPrefix("TRACE: ")
        Logger.Println(msg)
    }
}
