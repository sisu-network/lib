package log

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Ensure struct implements interface at compile-time
var _ Logger = (*DefaultLogger)(nil)

const (
	LOG_LEVEL_DEBUG = iota + 1
	LOG_LEVEL_VERBOSE
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
	LOG_LEVEL_CRITICAL
)

var globalLogger Logger
var logLevel int
var loggerLock = &sync.RWMutex{}

func init() {
	loggerLock.Lock()
	defer loggerLock.Unlock()

	globalLogger = newDefaultLogger()
	logLevel = LOG_LEVEL_INFO
}

func SetLogger(logger Logger) {
	loggerLock.Lock()
	defer loggerLock.Unlock()

	globalLogger = logger
}

func SetLogLevel(level int) {
	loggerLock.Lock()
	defer loggerLock.Unlock()

	logLevel = level
}

func getLogger() Logger {
	loggerLock.RLock()
	defer loggerLock.RUnlock()

	return globalLogger
}

func Info(a ...interface{}) {
	logger := getLogger()
	logger.Info(a...)
}

func Debug(a ...interface{}) {
	logger := getLogger()
	logger.Debug(a...)
}

func Warn(a ...interface{}) {
	logger := getLogger()
	logger.Warn(a...)
}

func Error(a ...interface{}) {
	logger := getLogger()
	logger.Error(a...)
}

func Verbose(a ...interface{}) {
	logger := getLogger()
	logger.Verbose(a...)
}

func Critical(a ...interface{}) {
	logger := getLogger()
	logger.Critical(a...)
}

func Infof(template string, a ...interface{}) {
	logger := getLogger()
	logger.Infof(template, a...)
}

func Debugf(template string, a ...interface{}) {
	logger := getLogger()
	logger.Debugf(template, a...)
}

func Warnf(template string, a ...interface{}) {
	logger := getLogger()
	logger.Warnf(template, a...)
}

func Errorf(template string, a ...interface{}) {
	logger := getLogger()
	logger.Errorf(template, a...)
}

func Verbosef(template string, a ...interface{}) {
	logger := getLogger()
	logger.Verbosef(template, a...)
}

// Default logging implementation. You can replace this logging module by another implementation
// of Logger interface.

type DefaultLogger struct {
	errorColor   *color.Color
	verboseColor *color.Color
	warningColor *color.Color
	debugColor   *color.Color
}

func newDefaultLogger() *DefaultLogger {
	return &DefaultLogger{
		errorColor:   color.New(color.FgRed),
		verboseColor: color.New(color.FgCyan),
		warningColor: color.New(color.FgYellow),
		debugColor:   color.New(color.FgMagenta),
	}
}

func (logger *DefaultLogger) Info(a ...interface{}) {
	logger.printWithTime(nil, "", a...)
}

func (logger *DefaultLogger) Debug(a ...interface{}) {
	logger.printWithTime(logger.debugColor, "", a...)
}

func (logger *DefaultLogger) Warn(a ...interface{}) {
	logger.printWithTime(logger.warningColor, "", a...)
}

func (logger *DefaultLogger) Error(a ...interface{}) {
	logger.printWithTime(logger.errorColor, "", a...)
}

func (logger *DefaultLogger) Verbose(a ...interface{}) {
	logger.printWithTime(logger.verboseColor, "", a...)
}

func (logger *DefaultLogger) HighVerbose(a ...interface{}) {
	logger.printWithTime(logger.verboseColor, "", a...)
}

func (logger *DefaultLogger) Critical(a ...interface{}) {
	logger.printWithTime(nil, "", a...)
}

func (logger *DefaultLogger) Infof(template string, a ...interface{}) {
	logger.printWithTime(nil, template, a...)
}

func (logger *DefaultLogger) Debugf(template string, a ...interface{}) {
	logger.printWithTime(logger.debugColor, template, a...)
}

func (logger *DefaultLogger) Warnf(template string, a ...interface{}) {
	logger.printWithTime(logger.warningColor, template, a...)
}

func (logger *DefaultLogger) Errorf(template string, a ...interface{}) {
	logger.printWithTime(logger.errorColor, template, a...)
}

func (logger *DefaultLogger) Verbosef(template string, a ...interface{}) {
	logger.printWithTime(logger.verboseColor, template, a...)
}

func (logger *DefaultLogger) HighVerbosef(template string, a ...interface{}) {
	logger.printWithTime(logger.verboseColor, template, a...)
}

func (logger *DefaultLogger) printWithTime(c *color.Color, template string, a ...interface{}) {
	now := time.Now().Format("15:04:05.00")
	msg := getMessage(template, a...)
	if c == nil {
		fmt.Println(now, msg)
	} else {
		c.Println(now, msg)
	}
}

func getMessage(template string, fmtArgs ...interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if len(template) > 0 {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}
