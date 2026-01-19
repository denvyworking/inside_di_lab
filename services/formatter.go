package services

import "fmt"

// Formatter - форматирует сообщения для логгера.
// Тоже без зависимостей.
type Formatter struct{}

// NewFormatter - конструктор Formatter.
func NewFormatter() *Formatter {
	return &Formatter{}
}

// Format - метод, который добавляет префикс к сообщению.
func (f *Formatter) Format(msg string) string {
	return fmt.Sprintf("[LOG] %s", msg)
}
