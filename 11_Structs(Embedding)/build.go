package main

import "fmt"

type Builder interface {
	Build()
}

type Person struct {
	Name string
	Age  int
}

type Building struct {
	Builder
	Name string
}

type WoodBuilder struct {
	Person
}

func (wb WoodBuilder) Build() {
	fmt.Println("Строю из дерева")
}

type BrickBuilder struct {
	Person
}

func (bb BrickBuilder) Build() {
	fmt.Println("Строю из камня")
}

type SuperBuilder struct {
	Person
}

func (s SuperBuilder) Build() {
	fmt.Println("Строю монолитный дом")
}

func main() {
	useless()
}

func useless() {
	woodenBuiding := Building{
		Builder: WoodBuilder{Person{
			Name: "Vladimir",
			Age:  44,
		}},
		Name: "Изба",
	}

	woodenBuiding.Build()

	brickBuilding := Building{
		Builder: BrickBuilder{
			Person{
				Name: "Egor",
				Age:  28,
			}},
		Name: "Скайскрепер",
	}

	brickBuilding.Build()

	superBuilding := Building{
		Builder: SuperBuilder{
			Person{
				Name: "Dasha",
				Age:  30,
			}},
		Name: "Монолит",
	}

	superBuilding.Build()
}
