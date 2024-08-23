package repository11

// go mod vendor - скачивает все зависимости в пакет программы

// как нужно указывать github.com/mrRestyling/repository11
// пакет должен быть package repository11

// Код в сеть:
// git tag v1.0.0

// Вторая версия(изменения):
// в go.mod меняем название github.com/mrRestyling/repository11/v2
// git tag v2.0.0

func Sum(numbers ...int) int {

	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum

}
