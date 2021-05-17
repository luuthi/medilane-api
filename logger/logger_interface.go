package logger

import (
	"github.com/sirupsen/logrus"
	"io"
)

const (
	TRACE = logrus.TraceLevel
	DEBUG = logrus.DebugLevel
	INFO  = logrus.InfoLevel
	WARN  = logrus.WarnLevel
	ERROR = logrus.ErrorLevel
	FATAL = logrus.FatalLevel
	PANIC = logrus.PanicLevel
)

func LogLevel(key string) logrus.Level {
	m := map[string]logrus.Level{
		"TRACE": logrus.TraceLevel,
		"DEBUG": logrus.DebugLevel,
		"INFO":  logrus.InfoLevel,
		"WARN":  logrus.WarnLevel,
		"ERROR": logrus.ErrorLevel,
		"FATAL": logrus.FatalLevel,
		"PANIC": logrus.PanicLevel,
	}

	lv, ok := m[key]
	if !ok {
		return logrus.InfoLevel
	}

	return lv
}

const timeFormat = "02-01-2006 15:04:05.000"

const (
	rotateMaxSize = 50 // megabytes
	rotateBackups = 10
	rotateDays    = 7
)

const EnvAssetTesting = "ENV_ASSET_TESTING"

type ConfigLogging struct {
	FileOut            string       `json:"file_out" yaml:"FILE_OUT"`                 // File path; default: stdout
	Level              string       `json:"level" yaml:"LEVEL"`                       // default: WARN
	FormatText         bool         `json:"format_text" yaml:"FORMAT_TEXT"`           // Log as ASCII formatter (log raw map[value]); default: JSON
	DisableCaller      bool         `json:"disable_caller" yaml:"DISABLE_CALLER"`     // Add the calling method as a field; default: always
	ShowCallerFullPath bool         `json:"show_caller_full_path" yaml:"SHOW_CALLER"` // Show method called with full path; default: only file name
	DisableLogFile     bool         `json:"disable_log_file" yaml:"DISABLE_LOG_FILE"`
	DisableConsoleLog  bool         `json:"disable_console_log" yaml:"DISABLE_CONSOLE_LOG"` // os.Stdout
	IOWriter           io.Writer    // add more writer; default os.Stdout & FileLog
	DisableRotate      bool         `json:"disable_rotate" yaml:"DISABLE_ROTATE"`
	Rotate             ConfigRotate `json:"rotate" yaml:"ROTATE"`
}

type ConfigRotate struct {
	MaxSize    int  `json:"max_size" yaml:"MAX_SIZE"`
	MaxBackups int  `json:"max_backups" yaml:"MAX_BACKUPS"`
	MaxDays    int  `json:"max_days" yaml:"MAX_DAYS"`
	Compress   bool `json:"compress" yaml:"COMPRESS"`
	LocalTime  bool `json:"local_time" yaml:"LOCAL_TIME"`
}

/*
type Fields logrus.Fields

type LoggerWrap interface {
	Init(conf Config) (LoggerWrap, error)
	Close()

	// implement
	// logrus.FieldLogger
	WithField(key string, value interface{}) //*logrus.Entry
	WithFields(fields Fields) *logrus.Entry
	WithError(err error) *logrus.Entry

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}
*/
