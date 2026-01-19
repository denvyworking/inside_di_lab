package services

import "fmt"

type AdvancedLogger struct {
	config       *Config
	formatter    *Formatter
	timeProvider *TimeProvider
}

// NewAdvancedLogger - конструктор с 3 параметрами
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
