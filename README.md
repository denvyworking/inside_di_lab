# Simple DI Container for Go

> Реализация собственного контейнера внедрения зависимостей (Dependency Injection) на языке Go 

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)](https://golang.org)

## Цель

Реализовать DI-контейнер, соответствующий трём требованиям:
1. **Определение типов в рантайме** через рефлексию.
2. **Выстраивание логики логгера через DI**.
3. **Обработка структурных зависимостей** с автоматическим определением порядка создания и защитой от циклов.

## Возможности

-  Автоматическое разрешение зависимостей любой глубины (DFS).
-  Поддержка любого числа параметров в конструкторах.
-  Кэширование (singleton): общие зависимости переиспользуются.
-  Защита от циклических зависимостей с читаемой ошибкой.
-  Управление жизненным циклом: очистка кэша в runtime.

## Архитектура
* di/container.go — ядро DI-контейнера.
* services/ — примеры компонентов (Logger, AdvancedLogger, Config, ...).
* main.go — интерактивная демонстрация.
- Используется только стандартная библиотека Go (reflect, fmt, errors).

## ▶️ Запуск

```bash
go run main.go
```

### Интерактивное меню:

```bash
1. Создать обычный Logger (2 зависимости)
2. Создать AdvancedLogger (3 зависимости)
3. Проверить циклическую зависимость
4. Очистить кэш DI-контейнера
0. Выход
```

### Пример вывода
```bash
При первом вызове:
    Запрашиваем *Logger...
    Creating new instance of *services.Config...
    Created *services.Config
    Creating new instance of *services.Formatter...
    Created *services.Formatter
    Creating new instance of *services.Logger...
    Created *services.Logger
    [LOG] Привет от обычного Logger!

При повторном вызове:
    Reusing existing instance of *services.Config
    Reusing existing instance of *services.Formatter
    Reusing existing instance of *services.Logger
``` 
