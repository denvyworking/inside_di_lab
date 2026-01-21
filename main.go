package main

import (
	"bufio"
	"fmt"
	"log"
	"my-di-lab/di"
	"my-di-lab/services"
	"os"
)

type App struct {
	logger services.LoggerInterface
}

func NewApp(logger services.LoggerInterface) *App {
	return &App{logger: logger}
}

func (a *App) Run() {
	a.logger.Log("Application started!")
	a.logger.Log("Dependency Injection works with dynamic constructor!")
}

// 2 - с кешированием
// ==========================================================

func main() {
	container := di.NewContainer()

	err := container.Register(services.NewConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Register(services.NewFormatter)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Register(services.NewTimeProvider)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1 — Создать обычный Logger (2 зависимости)")
		fmt.Println("2 — Создать AdvancedLogger (3 зависимости)")
		fmt.Println("3 — Проверить циклическую зависимость")
		fmt.Println("4 — Очистить кэш DI-контейнера")
		fmt.Println("0 — Выход")
		fmt.Print("Ваш выбор: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		switch input[0] {
		case '1':
			fmt.Println("\n Запрашиваем *Logger...")
			err = container.Register(services.NewLogger)
			if err != nil {
				log.Fatal(err)
			}
			var logger *services.Logger
			err := container.Resolve(&logger)
			if err != nil {
				log.Fatal("Ошибка:", err)
			}
			logger.Log("Привет от Logger!")

		case '2':
			fmt.Println("\n Запрашиваем *AdvancedLogger...")

			err = container.Register(services.NewAdvancedLogger)
			if err != nil {
				log.Fatal(err)
			}
			var logger *services.AdvancedLogger
			err := container.Resolve(&logger)
			if err != nil {
				log.Fatal("Ошибка:", err)
			}
			logger.Log("Привет от AdvancedLogger!")

		case '3':
			fmt.Println("\nТестируем цикл (отдельный контейнер)...")
			cycleContainer := di.NewContainer()
			cycleContainer.Register(services.NewServiceA)
			cycleContainer.Register(services.NewServiceB)
			var a *services.ServiceA
			err := cycleContainer.Resolve(&a)
			if err != nil {
				fmt.Printf("!!! Цикл обнаружен: %v !!!\n", err)
			}
		case '4':
			container.ClearCache()

		case '0':
			fmt.Println("👋")
			return

		default:
			fmt.Println("Неверный выбор.")
		}
	}
}
