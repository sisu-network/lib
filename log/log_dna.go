package log

import (
	"fmt"

	"github.com/logdna/logdna-go/logger"
)

const levelVerbose = "verbose"
const levelHighVerbose = "highverbose"

var _ Logger = (*DNALogger)(nil)

type LogDNAConfig struct {
	Secret        string   `toml:"secret" json:"secret,omitempty"`
	AppName       string   `toml:"app_name" json:"app_name,omitempty"`
	HostName      string   `toml:"host_name" json:"host_name,omitempty"`
	FlushInterval Duration `toml:"flush_interval" json:"flush_interval"`
	MaxBufferLen  int      `toml:"max_buffer_len" json:"max_buffer_len,omitempty"`
	LogLocal      bool     `toml:"log_local" json:"log_local,omitempty"`
}

type DNALogger struct {
	logger      *logger.Logger
	localLogger *DefaultLogger
}

func NewDNALogger(key string, opts logger.Options, logLocal bool) *DNALogger {
	if len(opts.Hostname) == 0 {
		panic("Hostname is required. For example: node0")
	}

	if len(opts.App) == 0 {
		panic("App is required. For example: sisu_main_app")
	}

	log, err := logger.NewLogger(opts, key)
	if err != nil {
		panic(err)
	}

	var localLogger *DefaultLogger
	if logLocal {
		localLogger = newDefaultLogger()
	}

	return &DNALogger{logger: log, localLogger: localLogger}
}

func (dna *DNALogger) Critical(a ...interface{}) {
	dna.logger.Critical(fmt.Sprintf("%v", a))

	if dna.localLogger != nil {
		dna.localLogger.Error(a...)
	}
}

func (dna *DNALogger) Error(a ...interface{}) {
	dna.logger.Error(fmt.Sprintf("%v", a))

	if dna.localLogger != nil {
		dna.localLogger.Error(a...)
	}
}

func (dna *DNALogger) Warn(a ...interface{}) {
	dna.logger.Warn(fmt.Sprintf("%v", a))

	if dna.localLogger != nil {
		dna.localLogger.Warn(a...)
	}
}

func (dna *DNALogger) Info(a ...interface{}) {
	dna.logger.Info(fmt.Sprintf("%v", a))

	if dna.localLogger != nil {
		dna.localLogger.Info(a...)
	}
}

func (dna *DNALogger) Verbose(a ...interface{}) {
	dna.logger.LogWithLevel(fmt.Sprintf("%v", a), levelVerbose)

	if dna.localLogger != nil {
		dna.localLogger.Verbose(a...)
	}
}

func (dna *DNALogger) HighVerbose(a ...interface{}) {
	dna.logger.LogWithLevel(fmt.Sprintf("%v", a), levelHighVerbose)

	if dna.localLogger != nil {
		dna.localLogger.HighVerbose(a...)
	}
}

func (dna *DNALogger) Debug(a ...interface{}) {
	dna.logger.Debug(fmt.Sprintf("%v", a))

	if dna.localLogger != nil {
		dna.localLogger.Debug(a...)
	}
}

func (dna *DNALogger) Errorf(template string, a ...interface{}) {
	dna.logger.Error(fmt.Sprintf("%v", getMessage(template, a...)))

	if dna.localLogger != nil {
		dna.localLogger.Errorf(template, a...)
	}
}

func (dna *DNALogger) Warnf(template string, a ...interface{}) {
	dna.logger.Warn(fmt.Sprintf("%v", getMessage(template, a...)))

	if dna.localLogger != nil {
		dna.localLogger.Warnf(template, a...)
	}
}

func (dna *DNALogger) Infof(template string, a ...interface{}) {
	dna.logger.Info(fmt.Sprintf("%v", getMessage(template, a...)))

	if dna.localLogger != nil {
		dna.localLogger.Infof(template, a...)
	}
}

func (dna *DNALogger) Verbosef(template string, a ...interface{}) {
	dna.logger.LogWithLevel(fmt.Sprintf("%v", getMessage(template, a...)), levelVerbose)

	if dna.localLogger != nil {
		dna.localLogger.Verbosef(template, a...)
	}
}

func (dna *DNALogger) HighVerbosef(template string, a ...interface{}) {
	dna.logger.LogWithLevel(fmt.Sprintf("%v", getMessage(template, a...)), levelHighVerbose)

	if dna.localLogger != nil {
		dna.localLogger.HighVerbosef(template, a...)
	}
}

func (dna *DNALogger) Debugf(template string, a ...interface{}) {
	dna.logger.Debug(fmt.Sprintf("%v", getMessage(template, a...)))

	if dna.localLogger != nil {
		dna.localLogger.Debugf(template, a...)
	}
}
