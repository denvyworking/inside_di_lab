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

func NewLoggerAsInterface(cfg *services.Config, fmt *services.Formatter) services.LoggerInterface {
	return services.NewLogger(cfg, fmt)
}

func NewAdvancedLoggerAsInterface(
	cfg *services.Config,
	fmt *services.Formatter,
	tp *services.TimeProvider,
) services.LoggerInterface {
	return services.NewAdvancedLogger(cfg, fmt, tp)
}

// ЦИКЛИЧЕСКИЕ ЗАВИСИМОСТИ
type ServiceA struct {
	B *ServiceB
}

type ServiceB struct {
	A *ServiceA
}

func NewServiceA(b *ServiceB) *ServiceA {
	return &ServiceA{B: b}
}

func NewServiceB(a *ServiceA) *ServiceB {
	return &ServiceB{A: a}
}

// ==========================================================

// func main() {
// 	fmt.Println("Выберите режим:")
// 	fmt.Println("1 - Обычный Logger (2 зависимости)")
// 	fmt.Println("2 - AdvancedLogger (3 зависимости)")
// 	fmt.Println("3 - Проверка циклической зависимости")
// 	fmt.Print("Ваш выбор (1, 2 или 3): ")

// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')

// 	container := di.NewContainer()

// 	// Регистрируем общие зависимости (они нужны почти везде)
// 	container.Register(services.NewConfig)
// 	container.Register(services.NewFormatter)
// 	container.Register(services.NewTimeProvider)

// 	switch input[0] {
// 	case '1':
// 		fmt.Println("\n Выбран обычный Logger (2 параметра)")
// 		container.Register(NewLoggerAsInterface)
// 		container.Register(NewApp)

// 		var app *App
// 		err := container.Resolve(&app)
// 		if err != nil {
// 			log.Fatal("Ошибка создания приложения:", err)
// 		}
// 		fmt.Println("\n✅ Запуск приложения:")
// 		app.Run()

// 	case '2':
// 		fmt.Println("\n Выбран AdvancedLogger (3 параметра)")
// 		container.Register(NewAdvancedLoggerAsInterface)
// 		container.Register(NewApp)

// 		var app *App
// 		err := container.Resolve(&app)
// 		if err != nil {
// 			log.Fatal("Ошибка создания приложения:", err)
// 		}
// 		fmt.Println("\n✅ Запуск приложения:")
// 		app.Run()

// 	case '3':
// 		fmt.Println("\n🔄 Тестируем циклическую зависимость (ServiceA ↔ ServiceB)...")

// 		// Регистрируем ТОЛЬКО циклические зависимости
// 		container.Register(NewServiceA)
// 		container.Register(NewServiceB)

// 		var a *ServiceA
// 		err := container.Resolve(&a)
// 		if err != nil {
// 			fmt.Printf("\n✅ ЦИКЛ УСПЕШНО ОБНАРУЖЕН!\n")
// 			fmt.Printf("❌ Ошибка: %v\n", err)
// 		} else {
// 			fmt.Println("❌ Цикл НЕ обнаружен — это ошибка!")
// 		}

// 	default:
// 		log.Fatal("Неверный выбор")
// 	}
// }

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
			fmt.Println("\n🔄 Тестируем цикл (отдельный контейнер)...")
			cycleContainer := di.NewContainer()
			cycleContainer.Register(NewServiceA)
			cycleContainer.Register(NewServiceB)
			var a *ServiceA
			err := cycleContainer.Resolve(&a)
			if err != nil {
				fmt.Printf("✅ Цикл обнаружен: %v\n", err)
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
