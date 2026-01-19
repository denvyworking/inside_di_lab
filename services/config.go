package services

// Config - простая структура, хранящая настройки.
// У неё нет зависимостей, она "листовой" узел в графе.
type Config struct {
	LogLevel string
}

// NewConfig - конструктор Config.
// Возвращает указатель на Config с дефолтными значениями.
func NewConfig() *Config {
	return &Config{
		LogLevel: "debug",
	}
}
