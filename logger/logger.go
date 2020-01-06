package logger

import (
	"github.com/mgutz/logxi/v1"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

var myLogger log.Logger

func Init(name string, logFileName string, logFileDir string, maxLogFileSize, maxLogFileBackups, maxLogFileAge int) {
	myLogger = log.NewLogger(os.Stdout, name)

	if len(logFileDir) >= 0 {
		dirExists, _ := dirExists(logFileDir)
		if dirExists {
			maxSize := maxLogFileSize
			maxBackups := maxLogFileBackups
			maxAge := maxLogFileAge

			if maxSize <= 0 {
				maxSize = 100
			}

			if maxBackups <= 0 {
				maxBackups = 3
			}

			if maxAge <= 0 {
				maxAge = 30
			}

			logFileOutput := &lumberjack.Logger{
				Filename:   filepath.Join(logFileDir, logFileName),
				MaxSize:    maxSize, // megabytes
				MaxBackups: maxBackups,
				MaxAge:     maxAge, //days
			}

			output := io.MultiWriter(os.Stdout, logFileOutput)
			myLogger = log.NewLogger(output, name)
		} else {
			myLogger.Warn("logFileDir does not exit", "dir", logFileDir)
		}
	} else {
		myLogger.Info("logFileDir is empty")
	}

}

func dirExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else {
		if os.IsNotExist(err) {
			return false, err
		} else {
			return true, err
		}
	}
}

func Trace(msg string, args ...interface{}) {
	myLogger.Trace(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	myLogger.Debug(msg, args...)
}

func Info(msg string, args ...interface{}) {
	myLogger.Info(msg, args...)
}

func Warn(msg string, args ...interface{}) error {
	return myLogger.Warn(msg, args...)
}

func Error(msg string, args ...interface{}) error {
	return myLogger.Error(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	myLogger.Fatal(msg, args...)
}

func Log(level int, msg string, args []interface{}) {
	myLogger.Log(level, msg, args)
}

func SetLevel(level int) {
	myLogger.SetLevel(level)
}

func IsTrace() bool {
	return myLogger.IsTrace()
}

func IsDebug() bool {
	return myLogger.IsDebug()
}

func IsInfo() bool {
	return myLogger.IsInfo()
}

func IsWarn() bool {
	return myLogger.IsWarn()
}
