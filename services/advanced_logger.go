package services

import "fmt"

type AdvancedLogger struct {
	config       *Config
	formatter    *Formatter
	timeProvider *TimeProvider
}

// NewAdvancedLogger — конструктор с ТРЕМЯ параметрами!
func NewAdvancedLogger(cfg *Config, fmt *Formatter, tp *TimeProvider) *AdvancedLogger {
	return &AdvancedLogger{
		config:       cfg,
		formatter:    fmt,
		timeProvider: tp,
	}
}

func (l *AdvancedLogger) Log(msg string) {
	if l.config.LogLevel == "debug" {
		timestamp := l.timeProvider.Now()
		output := fmt.Sprintf("[%s] %s", timestamp, l.formatter.Format(msg))
		fmt.Println(output)
	}
}

var _ LoggerInterface = (*AdvancedLogger)(nil)
