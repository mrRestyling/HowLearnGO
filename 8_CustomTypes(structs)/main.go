package main

import (
	"fmt"
	"time"
)

type OurString string
type OurInt int

type Person struct {
	Name string
	Age  int
}

func main() {

	// create default
	var Egor Person
	fmt.Printf("%T, %#v \n", Egor, Egor)

	Egor = Person{}
	fmt.Printf("%T, %#v \n", Egor, Egor)

	// fields accessing
	Egor.Name = "Egor"
	Egor.Age = 28
	fmt.Println(Egor)

	// 2 person
	Daria := Person{
		Name: "Daria",
		Age:  29,
	}
	fmt.Println(Daria)

	// 3 person
	Vitalic := Person{"Vitalic", 40}
	fmt.Println(Vitalic)

	// 4 person (pointer to struct)
	pVova := &Person{"Vova", 44}
	fmt.Println(pVova)
	fmt.Println(*pVova)

	// field acceessing through the pointer
	pDaria := &Daria
	fmt.Println((*pDaria).Age)
	fmt.Println(pDaria.Age)

	// строковый тип
	var customString OurString
	fmt.Printf("\n%T %#v \n", customString, customString)
	customString = "hello"
	fmt.Printf("%T %#v \n\n", customString, customString)

	// целочисленный тип
	customInt := OurInt(5)
	fmt.Printf("%T %#v \n", customInt, customInt)
	customInt = 7
	fmt.Printf("%T %#v \n", customInt, customInt)
	fmt.Println(int(customInt))

	// создание анонимной сткрутуры

	unnamedStruct := struct {
		Name, LastName, BirthDate string
	}{
		Name:     "NoName",
		LastName: "NoLastNamee",
		// BirthDate: fmt.Sprintf("%s", time.Now()),
		// BirthDate: time.Now().String(),
		BirthDate: time.Now().Format("2006.01"),
	}
	fmt.Println(unnamedStruct)
}
