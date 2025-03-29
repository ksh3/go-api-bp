package core

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ksh3/go-api/src/core/config"
	"github.com/ksh3/go-api/src/core/contract"
)

type LogFormat int

const (
	TextLogFormat LogFormat = iota
	JSONLogFormat
)

type LoggerConfig struct {
	OutputToFile bool
	LogFilePath  string
	LogFormat    LogFormat
}

type Logger struct {
	Config LoggerConfig
	File   *os.File
}

func NewLogger() (*Logger, error) {
	appEnv := os.Getenv("APP_ENV")

	// NOTE: Default to prod
	if appEnv == "" {
		appEnv = config.ProdEnvKey
	}

	logConf := LoggerConfig{
		OutputToFile: appEnv == config.DevEnvKey,
		LogFilePath:  "./debug.log",
		LogFormat:    TextLogFormat,
	}

	var logFile *os.File
	var err error

	if logConf.OutputToFile {
		logFile, err = os.OpenFile(
			logConf.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644,
		)
		if err != nil {
			return nil, err
		}
	}

	return &Logger{
		Config: logConf,
		File:   logFile,
	}, nil
}

func (l *Logger) logToOutput(logMsg string) {
	if l.Config.OutputToFile {
		fmt.Fprintln(l.File, logMsg)
	} else {
		log.Println(logMsg)
	}
}

func (l *Logger) InfoLog(msg string) {
	logMsg := fmt.Sprintf(
		"%s [INFO] %s", time.Now().Format("2006-01-02 15:04:05"), msg)
	l.logToOutput(logMsg)
}

func (l *Logger) ErrorLog(appErr *contract.AppError, formats ...LogFormat) {
	format := l.Config.LogFormat

	if len(formats) > 0 {
		format = formats[0]
	}

	switch format {
	case JSONLogFormat:
		l.JSONErrorLog(appErr)
	case TextLogFormat:
	default:
		l.TextErrorLog(appErr)
	}
}

func (l *Logger) TextErrorLog(appErr *contract.AppError) {
	logMsg := fmt.Sprintf(
		"%s [ERROR] IsCritical: %t Code: %d, Message: %s, Cause: %v",
		time.Now().Format("2006-01-02 15:04:05"),
		appErr.IsCritical,
		appErr.Code,
		appErr.Message,
		appErr.Cause,
	)
	l.logToOutput(logMsg)
}

func (l *Logger) JSONErrorLog(appErr *contract.AppError) {
	logEntry := map[string]interface{}{
		"timestamp":  time.Now().Format("2006-01-02 15:04:05"),
		"level":      "ERROR",
		"isCritical": appErr.IsCritical,
		"code":       appErr.Code,
		"message":    appErr.Message,
		"cause":      fmt.Sprintf("%v", appErr.Cause),
	}
	jsonLog, _ := json.Marshal(logEntry)
	l.logToOutput(string(jsonLog))
}

func (l *Logger) Close() error {
	if l.Config.OutputToFile {
		return l.File.Close()
	}
	return nil
}
