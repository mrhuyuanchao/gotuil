package goutil

import (
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Log 日志
type Log struct {
	logClient *logrus.Logger
}

// LogClient 日志
var LogClient *Log

// InitLog 初始化日志配置
func InitLog(logPath string) {
	logClient := logrus.New()
	if logPath == "" {
		return
	}
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	logClient.Out = src
	logClient.SetLevel(logrus.DebugLevel)
	logClient.SetFormatter(&logrus.JSONFormatter{})
	//baseLogPath := path.Join(logPath, "/logs")
	logWriter, err := rotatelogs.New(
		fmt.Sprintf("%s/%s", logPath, "%Y-%m-%d.log"),
		rotatelogs.WithLinkName(logPath+"/current.log"), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),           // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour),       // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)

	LogClient = &Log{logClient}
}

// Info 信息
func (l *Log) Info(msg string) {
	if l != nil && l.logClient != nil {
		l.logClient.Info(msg)
	}
}

// Debug 调试
func (l *Log) Debug(msg string) {
	if l != nil && l.logClient != nil {
		l.logClient.Debug(msg)
	}
}

// Error 错误
func (l *Log) Error(msg string) {
	if l != nil && l.logClient != nil {
		l.logClient.Error(msg)
	}
}

// Fatal 信息
func (l *Log) Fatal(msg string) {
	if l != nil && l.logClient != nil {
		l.logClient.Fatal(msg)
	}
}
