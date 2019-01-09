package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Options 日志配置
type Options struct {
	Level          string // 级别
	WithCallerHook bool   // 是否输出打日志的文件名和行号
	Formatter      string // 输出格式 json | text

	Write    bool   // 是否输出到文件
	Path     string // 如果要输出到文件，指定目录路径
	FileName string // 日志文件名

	MaxAge        int           // 只保存多少天的日志
	RotationCount uint          // 要保存多少个日志文件 只在maxAge 为 -1 时用
	RotationTime  time.Duration // 每隔多久记录一个文件

	Debug bool // if set true, separate
}

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Print(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
}

var L Logger

var logLevelMap = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
}

const (
	defaultFilename  = "goapp"
	defaultLogLevel  = "info"
	defaultFormatter = "text"
	defaultMaxAge    = -1
)

var defaultPath = os.TempDir()

func defaultOptions() *Options {
	return &Options{
		Level:        defaultLogLevel,
		Formatter:    defaultFormatter,
		Write:        false,
		Path:         defaultPath,
		FileName:     defaultFilename,
		MaxAge:       defaultMaxAge,
		RotationTime: time.Duration(7*24) * time.Hour,
		Debug:        false,
	}
}

func Default(conf *Options) (Logger, error) {
	if conf == nil {
		conf = defaultOptions()
	}
	logLevel, ok := logLevelMap[conf.Level]
	if !ok {
		logLevel = logrus.DebugLevel
	}

	logDir := conf.Path
	if logDir == "" {
		logDir = defaultPath
	}

	logFileName := conf.FileName
	if logFileName == "" {
		logFileName = defaultFilename
	}

	var formatter logrus.Formatter
	if conf.Formatter == "json" {
		formatter = &logrus.JSONFormatter{}
	} else {
		formatter = &logrus.TextFormatter{}
	}

	logger := logrus.New()
	logger.SetLevel(logLevel)

	if conf.Write {
		var maxAge time.Duration
		if conf.MaxAge == -1 {
			maxAge = -1
		} else {
			maxAge = time.Duration(maxAge) * time.Hour
		}

		rotationCount := conf.RotationCount
		rotationTime := conf.RotationTime

		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			return nil, err
		}

		logPath := filepath.Join(logDir, conf.FileName)

		writer, err := rotatelogs.New(
			logPath+".%Y%m%d%H%M.log",
			rotatelogs.WithClock(rotatelogs.UTC),
			rotatelogs.WithMaxAge(maxAge),
			rotatelogs.WithRotationCount(rotationCount),
			rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour),
		)

		if err != nil {
			return nil, err
		}

		// 滚动日志
		logger.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.DebugLevel: writer,
				logrus.InfoLevel:  writer,
				logrus.WarnLevel:  writer,
				logrus.ErrorLevel: writer,
				logrus.FatalLevel: writer,
			},
			formatter,
		))

		defaultLogFilePrex := logFileName + "."
		pathMap := lfshook.PathMap{
			// 常规日志
			logrus.DebugLevel: fmt.Sprintf("%s/%sout", logDir, defaultLogFilePrex),
			logrus.InfoLevel:  fmt.Sprintf("%s/%sout", logDir, defaultLogFilePrex),
			logrus.WarnLevel:  fmt.Sprintf("%s/%sout", logDir, defaultLogFilePrex),

			// 错误级别日志
			logrus.ErrorLevel: fmt.Sprintf("%s/%serr", logDir, defaultLogFilePrex),
			logrus.FatalLevel: fmt.Sprintf("%s/%serr", logDir, defaultLogFilePrex),
		}
		logger.AddHook(lfshook.NewHook(
			pathMap,
			formatter,
		))
	}
	// console
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05.000",
	})
	logger.Out = os.Stdout

	if conf.Debug {
		logger.SetReportCaller(true)
	}

	L = logger

	return L, nil
}
