package log

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	log         *slog.Logger
	filePath    = "/stack.log"
	aLog        *slog.Logger
	aFilePath   = "/access.log"
	logDir      = "./log"
	logSize     = 10   // MB
	logBucket   = 5    // 個数
	logAge      = 60   // days
	logCompress = true // 圧縮
)

func Initialize() {
	NewLogger()
}

func SetLogDir(dir string) {
	logDir = dir
}

func SetLogBucket(bucket int) {
	logBucket = bucket
}

func SetLogAge(age int) {
	logAge = age
}

func SetLogCompress(compress bool) {
	logCompress = compress
}

func NewLogger() {
	log = slog.New(slog.NewJSONHandler(NewWriter(filepath.Join(logDir, filePath)), nil))
	aLog = slog.New(slog.NewJSONHandler(NewWriter(filepath.Join(logDir, aFilePath)), nil))
}

func NewWriter(filepath string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    logSize,
		MaxBackups: logBucket,
		MaxAge:     logAge,
		Compress:   logCompress,
	}
}

func Access(args ...interface{}) {
	aLog.Info("Access log", args...)
}

func Info(msg string, args ...interface{}) {
	log.Info(msg, args...)
}

func Warning(msg string, args ...interface{}) {
	log.Warn(msg, args...)
}

func Error(msg string, args ...interface{}) {
	log.Error(msg, args...)
}

// Fatal exit application
func Fatal(msg string, args ...interface{}) {
	log.Error(msg, args...)
	os.Exit(1)
}

func Debug(msg string, args ...interface{}) {
	log.Debug(msg, args...)
}

func With(args ...interface{}) {
	log = log.With(args...)
}
