package logs

// updated
import (
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logprocess string
	logger     *logrus.Logger
)

func init() {
	logprocess = "instasafe"
	logLevel := "5"
	logLevelLogrus := logrus.WarnLevel
	switch logLevel {
	case "1":
		logLevelLogrus = logrus.FatalLevel
	case "2":
		logLevelLogrus = logrus.ErrorLevel
	case "3":
		logLevelLogrus = logrus.WarnLevel
	case "4":
		logLevelLogrus = logrus.InfoLevel
	case "5":
		logLevelLogrus = logrus.DebugLevel
	default:
		logLevelLogrus = logrus.WarnLevel
	}
	logFileName := strings.ToLower(logprocess)
	logFileName = strings.ReplaceAll(logFileName, " ", "_")
	logFileName = strings.ReplaceAll(logFileName, "-", "_")
	logger = logrus.New()
	logger.SetLevel(logLevelLogrus)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "logs/" + logFileName + ".log",
		MaxSize:    100,  // megabytes
		MaxBackups: 50,   //files
		MaxAge:     7,    //days
		Compress:   true, //.tar.gz
	})

}

func LogToFile(level string, fields logrus.Fields) {
	switch level {
	case "fatal":
		logger.WithFields(fields).Fatal()
	case "error'":
		logger.WithFields(fields).Error()
	case "warn":
		logger.WithFields(fields).Warn()
	case "info":
		logger.WithFields(fields).Info()
	case "debug":
		logger.WithFields(fields).Debug()
	default:
		logger.WithFields(fields).Info()
	}
}

func Fatal(log_message string, comment string, uuid string) {
	funcName, _, lineNo := getCallerInfo()
	fields := logrus.Fields{
		"logfile":      funcName,
		"logline":      lineNo,
		"loglevel":     "fatal",
		"logmessage":   log_message,
		"logprocess":   logprocess,
		"logtimestamp": time.Now().UTC().Format(time.RFC3339),
		"loguuid":      uuid,
		"comment":      comment,
	}
	LogToFile("fatal", fields)
}

func Error(log_message string, comment string, uuid string) {
	funcName, _, lineNo := getCallerInfo()
	fields := logrus.Fields{
		"logfile":      funcName,
		"logline":      lineNo,
		"loglevel":     "error",
		"logmessage":   log_message,
		"logprocess":   logprocess,
		"logtimestamp": time.Now().UTC().Format(time.RFC3339),
		"loguuid":      uuid,
		"comment":      comment,
	}
	LogToFile("error", fields)
}

func Warn(log_message string, comment string, uuid string) {
	funcName, _, lineNo := getCallerInfo()
	fields := logrus.Fields{
		"logfile":      funcName,
		"logline":      lineNo,
		"loglevel":     "warn",
		"logmessage":   log_message,
		"logprocess":   logprocess,
		"logtimestamp": time.Now().UTC().Format(time.RFC3339),
		"loguuid":      uuid,
		"comment":      comment,
	}
	LogToFile("warn", fields)
}

func Info(log_message string, comment string, uuid string) {
	funcName, _, lineNo := getCallerInfo()
	fields := logrus.Fields{
		"logfile":      funcName,
		"logline":      lineNo,
		"loglevel":     "info",
		"logmessage":   log_message,
		"logprocess":   logprocess,
		"logtimestamp": time.Now().UTC().Format(time.RFC3339),
		"loguuid":      uuid,
		"comment":      comment,
	}
	LogToFile("info", fields)
}

func Debug(log_message string, comment string, uuid string) {
	funcName, _, lineNo := getCallerInfo()
	fields := logrus.Fields{
		"logfile":      funcName,
		"logline":      lineNo,
		"loglevel":     "debug",
		"logmessage":   log_message,
		"logprocess":   logprocess,
		"logtimestamp": time.Now().UTC().Format(time.RFC3339),
		"loguuid":      uuid,
		"comment":      comment,
	}
	LogToFile("debug", fields)
}

func getCallerInfo() (function string, file string, line int) {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		details := runtime.FuncForPC(pc)
		function = details.Name()
	} else {
		function, file, line = "unknown", "unknown", 0
	}
	return
}
