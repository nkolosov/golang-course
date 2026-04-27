# Типы данных, структуры и интерфейсы

## Описание задачи

Реализовать структуру Point (точка с координатами X, Y) и функцию Distance (вычисляющую расстояние между двумя точками).

Реализовать структуру Polygon (набор точек) и структуру Circle (окружность: задается Point центра и радиус)

Структуры Polygon и Circle должны соответствовать указанным ниже интерфейсам

```go
package geometry

type Figure interface {
    Area() float64
    Perimeter() float64
    Contains(point Point) bool
}

type Point interface {
    DistanceTo(other Point) float64
}
```

Весь функционал должен быть покрыт тестами

### Запуск проекта
Проект должен запускаться в консольном режиме, должны быть реализованы все сценарии 

Примеры консольных команд для запуска (формат флагов, их названия и прочее могут отличаться в зависимости от деталей реализации)

Подсчет расстояния
```shell
$ /your/app --distance --point=X,Y --point=X,Y
```

Подсчет площади фигуры
```shell
$ /your/app --area [--circle=X,Y,R] [--center=X,Y --radius=R] [--polygon --point=X,Y, --point=X,Y --point=X,Y]
```

Проверка вхождения точки в фигуру
```shell
$ /your/app --contains --point=X,Y [--circle=X,Y,R] [--center=X,Y --radius=R] [--polygon --point=X,Y, --point=X,Y --point=X,Y]
```

При неправильном запуске должны отображаться корректные ошибки

### Детали реализации
- Проект должен соответствовать Go Project Layout
- В проекте должен быть .gitignore, Makefile, go.mod
- К проекту должен быть подключен линтер https://golangci-lint.run/ и в проекте должен находиться файл конфигурации .golangci-lint.yaml со стандартным набором линтероа
```yaml
version: "2"

linters:
  default: standard
```
- В пакете main не должно быть бизнес-логики
- Для работы с флагами можно использовать пакет `flag` из стандартной или более продвинутый [pflag](https://github.com/spf13/pflag)
- Для тестирования необходимо использовать пакет [testify](https://github.com/stretchr/testify)
- Тесты необходимо сделать в формате табличных тестов, где это возможно