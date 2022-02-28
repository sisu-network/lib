package log

import (
	"fmt"

	"github.com/logdna/logdna-go/logger"
)

const levelVerbose = "verbose"

var _ Logger = (*DNALogger)(nil)

type LogDNAConfig struct {
	Secret        string   `toml:"secret"`
	AppName       string   `toml:"app_name"`
	HostName      string   `toml:"host_name"`
	FlushInterval duration `toml:"flush_interval"`
	MaxBufferLen  int      `toml:"max_buffer_len"`
}

type DNALogger struct {
	logger *logger.Logger
}

func NewDNALogger(key string, opts logger.Options) *DNALogger {
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

	return &DNALogger{logger: log}
}

func (dna *DNALogger) Critical(a ...interface{}) {
	dna.logger.Critical(fmt.Sprintf("%v", a))
}

func (dna *DNALogger) Error(a ...interface{}) {
	dna.logger.Error(fmt.Sprintf("%v", a))
}

func (dna *DNALogger) Warn(a ...interface{}) {
	dna.logger.Warn(fmt.Sprintf("%v", a))
}

func (dna *DNALogger) Info(a ...interface{}) {
	dna.logger.Info(fmt.Sprintf("%v", a))
}

func (dna *DNALogger) Verbose(a ...interface{}) {
	dna.logger.LogWithLevel(fmt.Sprintf("%v", a), levelVerbose)
}

func (dna *DNALogger) Debug(a ...interface{}) {
	dna.logger.Debug(fmt.Sprintf("%v", a))
}

func (dna *DNALogger) Errorf(template string, a ...interface{}) {
	dna.logger.Error(fmt.Sprintf("%v", getMessage(template, a...)))
}

func (dna *DNALogger) Warnf(template string, a ...interface{}) {
	dna.logger.Warn(fmt.Sprintf("%v", getMessage(template, a...)))
}

func (dna *DNALogger) Infof(template string, a ...interface{}) {
	dna.logger.Info(fmt.Sprintf("%v", getMessage(template, a...)))
}

func (dna *DNALogger) Verbosef(template string, a ...interface{}) {
	dna.logger.LogWithLevel(fmt.Sprintf("%v", getMessage(template, a...)), levelVerbose)

}
func (dna *DNALogger) Debugf(template string, a ...interface{}) {
	dna.logger.Debug(fmt.Sprintf("%v", getMessage(template, a...)))
}
