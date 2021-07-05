package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// init instance
func Init(conf ConfigLogging) *logrus.Logger {
	// new logger
	logWriter := logrus.New()

	writers := make([]io.Writer, 0)

	// set output file
	if !conf.DisableLogFile && conf.FileOut != "" {
		fPath, err := createFile(conf.FileOut)
		if err != nil {
			logWriter.WithFields(logrus.Fields{
				"path":  fPath,
				"error": err,
			}).Error("Failed to create logs folder")
		} else {
			file, err := os.OpenFile(fPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
			if err == nil {
				if conf.DisableRotate {
					// logWriter.Out = file
					writers = append(writers, file)
				} else {
					// setup rotate
					logRotate := NewLogRotate(fPath, &conf.Rotate)
					writers = append(writers, logRotate)
				}
				logWriter.WithField("path", fPath).Info("Config log to file ok")
			} else {
				logWriter.WithFields(logrus.Fields{
					"path":  fPath,
					"error": err,
				}).Error("Failed to open log file, using default stdout")
			}
		}
	}

	// more writer?
	if conf.IOWriter != nil {
		writers = append(writers, conf.IOWriter)
	}

	// default enable console log
	if !conf.DisableConsoleLog {
		writers = append(writers, os.Stdout)
	}

	// set Level
	if conf.Level != "" {
		level, err := logrus.ParseLevel(conf.Level)
		if err != nil {
			logWriter.SetLevel(WARN)
		} else {
			logWriter.SetLevel(level)
		}
	} else {
		logWriter.SetLevel(WARN)
	}

	// enable only Stdout if env ENV_ASSET_TESTING is on
	envTesting := os.Getenv(EnvAssetTesting)

	// Output default stdout
	if len(writers) == 0 || envTesting != "" {
		logWriter.SetOutput(os.Stdout)

		if envTesting != "" {
			logWriter.SetLevel(DEBUG)
			logWriter.WithField("TESTING", "on").Warn("Disable log to file for testing")
		}
	} else if len(writers) == 1 {
		// check 1 for highlight in tty (if log Stdout)
		logWriter.SetOutput(writers[0])
	} else {
		mWriter := io.MultiWriter(writers...)
		logWriter.SetOutput(mWriter)
	}

	// Add the calling method as a field
	logWriter.SetReportCaller(!conf.DisableCaller)

	// Log as JSON or the default ASCII formatter
	var formatter logrus.Formatter

	if !conf.FormatText {
		formatter = &logrus.JSONFormatter{
			TimestampFormat:   timeFormat,
			DisableTimestamp:  false,
			DisableHTMLEscape: true,
			DataKey:           "",
			FieldMap:          nil,
			PrettyPrint:       false,
		}
		// split method called to only name instead of full path
		if !conf.DisableCaller && !conf.ShowCallerFullPath {
			formatter.(*logrus.JSONFormatter).CallerPrettyfier = callerPrettyfier
		}
		logWriter.SetFormatter(formatter)
	} else {
		formatter = &logrus.TextFormatter{
			TimestampFormat:  timeFormat,
			DisableTimestamp: false,
		}
		if !conf.DisableCaller && !conf.ShowCallerFullPath {
			formatter.(*logrus.TextFormatter).CallerPrettyfier = callerPrettyfier
		}
	}
	logWriter.SetFormatter(formatter)

	return logWriter
}
