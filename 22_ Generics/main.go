package main

import (
	"fmt"
)

// Некоторое обобщение типов
// функции/типы

type Number interface {
	~int64 | float64
}

func main() {
	// ShowSum()
	// ShowContaints()
	// ShowAny()
	// UnionInterfaceAndType()
	TypeApproximation()

}

func ShowSum() {
	fl := []float64{1.0, 2.5, 0.5}

	in := []int64{1, 2, 3, 4, 5}

	// muchVAR := []interface{}{"string", struct{}{}, true}

	fmt.Println(sum(fl))
	fmt.Println(sum[int64](in))
	// fmt.Println(sum(muchVAR))
}

// [] - type parametrs
func sum[V int64 | float64](numbers []V) V {
	var sum V

	for _, n := range numbers {
		sum += n
	}
	return sum
}

func ShowContaints() {

	type Person struct {
		Name     string
		Age      int64
		JobTitle string
	}

	ints := []int64{2, 4, 5, 6, 7, 8, 48, 44}
	fmt.Println(contains(ints, 4))

	strings := []string{"sql", "postgre", "json"}
	fmt.Println(contains(strings, "rabbit"))

	people := []Person{
		{
			Name:     "Egor",
			Age:      28,
			JobTitle: "NatC",
		},
		{
			Name:     "Daria",
			Age:      30,
			JobTitle: "GA",
		},
	}

	// fmt.Println(people)

	fmt.Println(contains(people, Person{Name: "Egor", Age: 28, JobTitle: "NatC"}))
}

func contains[T comparable](elements []T, searchEl T) bool {
	for _, el := range elements {
		if el == searchEl {
			return true
		}
	}
	return false
}

func ShowAny() {
	show(1, 2, 3)
	show("Egor", "Daria", "Lucky")
	show([]float64{1.0, 2.0}, []float64{3.0, 4.0})
	show(map[int]string{1: "one", 2: "two"}, map[int]string{}, map[int]string{4: "four"})
	show(interface{}(1), interface{}(!true), any(struct{ name string }{name: "Egor"}))
}

func show[T any](objects ...T) {
	fmt.Println(objects)
}

// обобщенный тип:
type Numbers[T Number] []T

func UnionInterfaceAndType() {
	var ints Numbers[int64]
	ints = append(ints, []int64{1, 2, 3, 4, 5}...)

	// одной строкой тоже можно:
	floats := Numbers[float64]{1.1, 2.2, 3.5}

	fmt.Println(sumUnionInterface(ints))
	fmt.Println(sumUnionInterface(floats))

}

// Number - интерфейс с типами (чтобы не перечислять их здесь)
func sumUnionInterface[V Number](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// тип основаный на int64 (~)
type CustomInt int64

func (ci CustomInt) IsPositive() bool {
	return ci > 0
}
func TypeApproximation() {
	customInts := []CustomInt{5, 5, 4, 6}
	// // нельзя так делать без второго способа.
	// (тильда в интерфейсе)
	fmt.Println(sumUnionInterface(customInts))

	// Так можно.
	// 1 способ:
	castedInts := make([]int64, len(customInts))

	for i, val := range customInts {
		castedInts[i] = int64(val)
	}

	fmt.Println(sumUnionInterface(castedInts))

	// 2 способ:
	// Приближение типа: "~"(тильда) в Интерфейсе

	// type Number interface {
	// 	~int64 | float64
	// }

}
