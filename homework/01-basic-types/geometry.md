## Упражнение
### Цель
В рамках упражнения учимся писать простые консольные команды и тесты.

### Требования
Необходимо реализовать консольную команду для расчета расстояния между двумя точкам, вычисления периметра и площади различных геометрических фигур.
Точки и геометрические фигуры должны быть отдельными структурами. Для всех методов должны быть написаны тесты.

Консольная команда должна запускаться при помощи флагов.
Возможные кейсы использования программы
```
$ ./geometry --distance --point=X,Y --point=X,Y
> Distance between (X1, Y1) and (X2, Y2) is 0.75
> Can't parse Point(...) with error: Error description

$ ./geometry [--perimeter | --square] [--circle --center=X,Y --radius=N] [--polygon --point=X,Y --point=X,Y --point=X,Y --point=X,Y]
> Perimeter of Figure is 2.45
> Square of Figure is 2.45

$ ./geometry --contains --checkpoint=X,Y [--circle --center=X,Y --radius=N] [--polygon --point=X,Y --point=X,Y --point=X,Y --point=X,Y]
> (X, Y) is inside Figure
> (X, Y) is outside Figure

$ ./geometry [--help]
Help:
     --distance --point=X,Y --point=X,Y
    [--perimeter | --square] [--circle --center=X,Y --radius=N] [--polygon --point=X,Y --point=X,Y --point=X,Y --point=X,Y]
    --contains --checkpoint=X,Y [--circle --center=X,Y --radius=N] [--polygon --point=X,Y --point=X,Y --point=X,Y --point=X,Y]
```

### Рекомендации по выполнению
Реализация геометрических фигур должна быть в отдельном пакете (не в main)

Для работы с флагами рекомендую использовать библиотеку pflag: https://github.com/spf13/pflag

Для работы с тестами лучше всего использовать библиотеку testify https://github.com/stretchr/testify

Все тесты должны быть в табличном формате.