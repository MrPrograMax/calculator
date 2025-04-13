# Вычислятор

## Описание

**Вычислятор** — это приложение для обработки списка арифметических инструкций с возможностью вывода значений переменных.

Каждая инструкция может быть одного из двух типов:

- `calc`: Выполняет арифметическую операцию (`+`, `-`, `*`) над двумя аргументами и сохраняет результат в переменную.
- `print`: Выводит значение указанной переменной.

### Особенности

- Каждая арифметическая операция считается "дорогой" (время выполнения — 50 мс).
- Каждую переменную можно вычислить только один раз.
- Задействует только необходимые для вывода переменные (оптимизация по зависимостям).
- Реализовано на Go, запускается в Docker-контейнере.

## Порты

gRPC :8090 \
REST :8080

## Примеры

### Входные данные

```json
[
  { "type": "calc", "op": "+", "var": "x", "left": 1, "right": 2 },
  { "type": "print", "var": "x" }
]
```

### Выходные данные
```json
{
  "items": [
    { "var": "x", "value": 3 }
  ]
}
```

## Makefile

Убедитесь, что у вас установлен `Docker` и `make`.

```bash
make build - билд 

make run - запуск

make stop - остановка

make logs - просмотр логов
```

```bash
make test - Показывает покрытие по каждому пакету

make test-total - Показает общий процент покрытия тестами:
```


