package logger

import (
	logRotate "gopkg.in/natefinch/lumberjack.v2"
)

func NewLogRotate(fPath string, conf *ConfigRotate) *logRotate.Logger {
	ins := &logRotate.Logger{
		Filename:  fPath,
		Compress:  false,
		LocalTime: false,
	}

	// check default
	if conf.MaxSize != 0 {
		ins.MaxSize = conf.MaxSize
		ins.MaxBackups = conf.MaxBackups
		ins.MaxAge = conf.MaxDays
	} else {
		ins.MaxSize = rotateMaxSize
		ins.MaxBackups = rotateBackups
		ins.MaxAge = rotateDays
	}

	return ins
}
