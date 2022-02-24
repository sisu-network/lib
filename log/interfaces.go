package log

type Logger interface {
	Critical(a ...interface{})
	Error(a ...interface{})
	Warn(a ...interface{})
	Info(a ...interface{})
	Verbose(a ...interface{})
	Debug(a ...interface{})
	Errorf(template string, a ...interface{})
	Warnf(template string, a ...interface{})
	Infof(template string, a ...interface{})
	Debugf(template string, a ...interface{})
	Verbosef(template string, a ...interface{})
}
