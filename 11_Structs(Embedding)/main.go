// package main

// import (
// 	"fmt"
// )

// type Person struct {
// 	Name string
// 	Age  int
// }

// type Builder interface {
// 	Build()
// }

// type WorkExpirience struct {
// 	Name string
// 	Age  int
// }

// type WoodBuilder struct {
// 	Person
// 	// Name string
// 	// WorkExpirience
// }

// func (p Person) printName() {
// 	fmt.Println(p.Name)
// }

// type Building struct {
// 	Builder
// 	Name string
// }

// func main() {
// 	explanation()
// }

// func explanation() {
// 	// builder := WoodBuilder{Person{Name: "Egor", Age: 28}}
// 	// builder := WoodBuilder{Person{Name: "Egor", Age: 28}, "newEgor"}
// 	// builder := WoodBuilder{
// 	// 	Person{Name: "Egor", Age: 28},
// 	// 	"newEgor",
// 	// 	WorkExpirience{"Таксист", 3},
// 	// }

// 	fmt.Printf("Type: %T, Value: %#v\n", builder, builder)

// 	fmt.Println(builder.Person.Age)
// 	// свойство колизии, не знает какое свойство достать (на одном уровне)
// 	// нужно написать полный путь
// 	fmt.Println(builder.WorkExpirience.Age)

// 	// shadowing builder.Name
// 	fmt.Println(builder.Name)
// 	fmt.Println(builder.Person.Name)

// 	// ищет ближайший метод builder.printName()
// 	builder.printName()
// }
