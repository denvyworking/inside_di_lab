package di

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

/*  var x int = 42
t := reflect.TypeOf(x)   // Type: int
v := reflect.ValueOf(x)  // Value: 42
*/

// reflect.Value - это специальный контейнер, который позволяет
// вскрыть любое значение и работать с ним динамически.
/*
Он знает:

Какой у значения тип,
Какое у него значение,
Как его прочитать, изменить или (если это функция) вызвать.
*/

type registration struct {
	constructor reflect.Value
}

type Container struct {
	// хранит рецепт - способ создания (функция)
	registrations map[reflect.Type]registration
	// хранит уже готовый экземпляр (обьект)
	instances map[reflect.Type]reflect.Value // кэш для singleton
}

func NewContainer() *Container {
	return &Container{
		registrations: make(map[reflect.Type]registration),
		instances:     make(map[reflect.Type]reflect.Value),
	}
}

func (c *Container) Register(constructor interface{}) error {
	constructorValue := reflect.ValueOf(constructor)
	if constructorValue.Kind() != reflect.Func {
		return errors.New("can only register functions")
	}

	funcType := constructorValue.Type()
	if funcType.NumOut() != 1 {
		return errors.New("constructor must return exactly one value")
	}

	resultType := funcType.Out(0)
	c.registrations[resultType] = registration{
		constructor: constructorValue,
	}
	return nil
}

// Resolve создаёт экземпляр указанного типа.
func (c *Container) Resolve(target interface{}) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
	}

	// разыменование указателя (переается всегда указатель на тип)
	desiredType := targetValue.Type().Elem()

	path := []reflect.Type{} // порядок важен для сообщения об ошибке

	instance, err := c.resolveType(desiredType, path)
	if err != nil {
		return err
	}
	// устанавливаем тип. Конечный этап работы DI
	targetValue.Elem().Set(instance)
	return nil
}

// resolveType рекурсивно разрешает зависимости с защитой от циклов.
func (c *Container) resolveType(
	t reflect.Type,
	path []reflect.Type,
) (reflect.Value, error) {
	// Проверка на цикл: если тип уже в текущем пути — ошибка
	for _, p := range path {
		// наглядная обработка результата
		if p == t {
			var names []string
			for _, pt := range path {
				names = append(names, pt.String())
			}
			names = append(names, t.String()) // замыкаем цикл
			return reflect.Value{}, fmt.Errorf("circular dependency detected: %s",
				strings.Join(names, " -> "))
		}
	}

	// Singleton: если уже создан - возвращаем
	if instance, exists := c.instances[t]; exists {
		fmt.Printf("Reusing existing instance of %v\n", t)
		return instance, nil
	}

	// Ищем конструктор
	reg, exists := c.registrations[t]
	if !exists {
		return reflect.Value{}, fmt.Errorf("no constructor registered for type %v", t)
	}

	constructor := reg.constructor // метод нашего конструктора
	// logger (config и formatter)
	// !!!  funcType = func (cfg *Config, fmt *Formatter) *logger  !!!
	funcType := constructor.Type() // вся инфа о типе параметров конструктора
	numIn := funcType.NumIn()      // кол-во зависимостей

	// Добавляем текущий тип в путь
	newPath := append(path, t)

	// Рекурсивно разрешаем зависимости
	args := make([]reflect.Value, numIn)
	for i := 0; i < numIn; i++ {
		argType := funcType.In(i)
		argValue, err := c.resolveType(argType, newPath)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("failed to resolve dependency %v for %v: %w", argType, t, err)
		}
		args[i] = argValue
	}

	// Создаём новый экземпляр
	fmt.Printf("Creating new instance of %v...\n", t)

	// constructor.Call(args) — вызов функции через рефлексию.
	// В Go функция может возвращать несколько значений, поэтому Call возвращает срез []reflect.Value.
	// Но конструктор возвращает один объект -> results[0].

	results := constructor.Call(args)
	instance := results[0]

	// Кэшируем (singleton)
	c.instances[t] = instance
	fmt.Printf("Created %v\n", t)

	return instance, nil
}

func (c *Container) ClearCache() {
	c.instances = make(map[reflect.Type]reflect.Value)
	fmt.Println("Кэш DI-контейнера очищен.")
}
