package main

import (
	"fmt"
	"math"
)

type Shape interface {
	CalculateArea() float64
}

type Repoter interface {
	GetReport() string
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func (r Rectangle) CalculateArea() float64 {
	return r.height * r.width
}

func (r Rectangle) GetReport() string {
	return fmt.Sprintf("Прямоугольник: ширина: %f, высота: %f\nПлощадь: %f", r.width, r.height, r.CalculateArea())
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) GetReport() string {
	return fmt.Sprintf("Круг: радиус: %f\nПлощадь: %f", c.radius, c.CalculateArea())
}

func (c Triangle) CalculateArea() float64 {
	return c.base * c.height / 2
}

func (c Triangle) GetReport() string {
	return fmt.Sprintf("Треугольник: основание: %f, высота: %f\nПлощадь: %f", c.base, c.height, c.CalculateArea())
}

func calculateArea(s any) (float64, error) {
	if v, ok := s.(Shape); ok {
		return v.CalculateArea(), nil
	}

	return 0, fmt.Errorf("%s", "Ошибка: переданный объект не является фигурой")
}

func getReport(r Repoter) string {
	return r.GetReport()
}

func main() {
	circle := Circle{
		radius: 5,
	}

	rectangle := Rectangle{
		width:  10,
		height: 5,
	}

	triangle := Triangle{
		base:   8,
		height: 6,
	}

	area, err := calculateArea(&rectangle)
	if err != nil {
		panic(err)
	}

	if rectangle.CalculateArea() != area {
		panic("invalid rectangl area value")
	}

	area, err = calculateArea(&circle)
	if err != nil {
		panic(err)
	}

	if circle.CalculateArea() != area {
		panic("invalid circle area value")
	}

	area, err = calculateArea(&triangle)
	if err != nil {
		panic(err)
	}

	if triangle.CalculateArea() != area {
		panic("invalid triangle area value")
	}

	_, err = calculateArea(4)
	if err == nil {
		panic("err must be nonnil")
	}

	fmt.Println(getReport(&circle))
	fmt.Println(getReport(&triangle))
	fmt.Println(getReport(&rectangle))
	fmt.Println(err)
}
