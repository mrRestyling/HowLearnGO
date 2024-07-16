package main

import "fmt"

type Runner interface { // далее перечисляем сигнатуры наших методов
	Run() string
}

type Swimmer interface {
	Swim()
}

type Flyer interface {
	Fly()
}

// встраиваемый интерфейс
type Ducker interface {
	Runner
	Swimmer
	Flyer
}

// имплементация интерфейса
type Human struct {
	Name string
}

func (h Human) Run() string {
	return fmt.Sprintf("Человек %s бегает", h.Name)
}

func (h Human) writeCode() {
	fmt.Println("Человек пишет код")
}

type Duck struct {
	Name, Surname string
}

func (d Duck) Run() string {
	return "Утка бегает"
}

func (d Duck) Swim() string {
	return "Утка плавает"
}

func (d Duck) Fly() string {
	return "Утка летает"
}

// func interfaceValue() {
// 	var runner Runner
// 	fmt.Printf("\nrunner\nType: %T ; Value: %#v\n", runner, runner)

// 	if runner == nil {
// 		fmt.Printf("runner is nil\n\n")
// 	}

// 	// имплементация интерфейса
// 	// Имплементация интерфейса в Go происходит путем реализации методов, определенных в интерфейсе, в структуре данных.
// 	// Явной имплементации нет, нужно просто иметь такие же методы
// 	// Утиная типизация (duck typing)

// 	//  runner = int64(1)
// 	// runner.Run()

// 	var unnamedRunner *Human
// 	fmt.Printf("unnamedRunner\nType: %T ; Value: %#v\n\n", unnamedRunner, unnamedRunner)

// 	// присваивание интерфейсному значению
// 	runner = unnamedRunner
// 	fmt.Printf("runner\nType: %T ; Value: %#v\n\n", runner, runner)
// 	// значение уже не равно нил, так как у интерфейсного значения есть конкретный тип
// 	if runner == nil {
// 		fmt.Printf("Egor is nil\n\n")
// 	}

// 	namedRunner := &Human{Name: "Egor"}
// 	fmt.Printf("namedRunner\nType: %T ; Value: %#v\n\n", namedRunner, namedRunner)

// 	runner = namedRunner
// 	fmt.Printf("runner\nType: %T ; Value: %#v\n", runner, runner)

// 	// * значение интерфейса состоит из двух частей:
// 	// Конкретный тип и конкретное значение

// 	// Пустой интерфейс (у которого нет методов)
// 	var emptyInterface interface{} = unnamedRunner // кладем значение структуры
// 	fmt.Printf("\nemptyInterface\nType: %T ; Value: %#v\n", emptyInterface, emptyInterface)

// 	emptyInterface = runner
// 	fmt.Printf("1 emptyInterface\nType: %T ; Value: %#v\n", emptyInterface, emptyInterface)

// 	emptyInterface = int64(1)
// 	fmt.Printf("2 emptyInterface\nType: %T ; Value: %#v\n", emptyInterface, emptyInterface)

// 	emptyInterface = true
// 	fmt.Printf("3 emptyInterface\nType: %T ; Value: %#v\n", emptyInterface, emptyInterface)
// }

func main() {
	// interfaceValue()
	typeAssertionAndPolymorphism()
}

func typeAssertionAndPolymorphism() {

	var runner Runner
	fmt.Printf("\nrunner\nType: %T ; Value: %#v\n", runner, runner)

	Daria := &Human{Name: "Daria"}
	runner = Daria
	polymorphism(Daria)
	typeAssertion(Daria)

	Donald := &Duck{"Donald", "DK"}
	runner = Donald
	polymorphism(Donald)
	typeAssertion(Donald)
}

// Полиморфизм в Go - это свойство объектно-ориентированного программирования,
// которое позволяет использовать один и тот же интерфейс для различных типов данных.
func polymorphism(runner Runner) {
	fmt.Println(runner.Run())
}

// Из интерфесного значения извлекаем значение конкретного типа
func typeAssertion(runner Runner) {
	fmt.Printf("\nTypeAssertion\nType: %T ; Value: %#v\n\n", runner, runner)

	if human, ok := runner.(*Human); ok {
		fmt.Printf("Type: %T ; Value: %#v\n\n", human, human)
		human.writeCode()
	}
	if duck, ok := runner.(*Duck); ok {
		fmt.Printf("Type: %T ; Value: %#v\n\n", duck, duck)
		fmt.Println(duck.Fly())
	}

	switch v := runner.(type) {
	case *Human:
		fmt.Println(v.Run())
	case *Duck:
		fmt.Println(v.Swim())
	default:
		fmt.Printf("Type: %T, Value: %#v\n", v, v)
	}
}
