package main

import "fmt"

// func (receiver T) methodName()  {} // value receiver
// func (receiver *T) methodName() {} // pointer receiver

// type Ourtype string
type Square struct {
	Side int
}

func (s Square) Perimeter() {
	fmt.Printf("\n%T, %#v \n", s, s)
	fmt.Printf("Периметр фигуры: %d \n\n", s.Side*4)
}

func (s *Square) Scale(multiplier int) {
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}

func (s Square) WrongScale(multiplier int) {
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}

func definition() {
	squareRoma := Square{Side: 4}
	pSquareRoma := &squareRoma

	squareDima := Square{2}

	squareRoma.Perimeter()
	squareDima.Perimeter()

	pSquareRoma.Scale(2)

	pSquareRoma.Perimeter() // сам разыменовывает указатель,s
	// потом обычное значение передает в нашу функцию
	// (*pSquareRoma).Perimeter() Square{Slide:8}

	squareRoma.Scale(2)     // (&squareRoma).Scale
	pSquareRoma.Perimeter() // (*pSquareRoma).Perimeter()

	squareRoma.WrongScale(2)
	squareRoma.Perimeter()
}

func rules() {

}
func main() {
	definition()
}
