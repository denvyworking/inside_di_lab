package services

import "fmt"

// Logger - зависит от Config и Formatter.
// Это наш первый пример "структурной зависимости":
// чтобы создать Logger, нужны два других объекта.

type Logger struct {
	config    *Config
	formatter *Formatter
}

// NewLogger - конструктор Logger.
// Принимает зависимости как параметры.
// Именно по этой сигнатуре DI-контейнер поймёт,
// какие объекты нужно создать перед вызовом этой функции.
func NewLogger(cfg *Config, fmt *Formatter) *Logger {
	return &Logger{
		config:    cfg,
		formatter: fmt,
	}
}

// Log - метод, который использует зависимости.
func (l *Logger) Log(msg string) {
	if l.config.LogLevel == "debug" {
		fmt.Println(l.formatter.Format(msg))
	}
}

type LoggerInterface interface {
	Log(msg string)
}

var _ LoggerInterface = (*Logger)(nil) // компилятор проверит соответствие
