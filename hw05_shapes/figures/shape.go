package figures

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
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) CalculateArea() float64 {
	return r.Height * r.Width
}

func (r Rectangle) GetReport() string {
	return fmt.Sprintf("Прямоугольник: ширина: %f, высота: %f\nПлощадь: %f", r.Width, r.Height, r.CalculateArea())
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) GetReport() string {
	return fmt.Sprintf("Круг: радиус: %f\nПлощадь: %f", c.Radius, c.CalculateArea())
}

func (c Triangle) CalculateArea() float64 {
	return c.Base * c.Height / 2
}

func (c Triangle) GetReport() string {
	return fmt.Sprintf("Треугольник: основание: %f, высота: %f\nПлощадь: %f", c.Base, c.Height, c.CalculateArea())
}
