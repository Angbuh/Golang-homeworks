package main

import (
	"fmt"

	"github.com/angbuh/golang-homeworks/hw05_shapes/figures"
)

func main() {
	circle := figures.Circle{
		Radius: 5,
	}

	rectangle := figures.Rectangle{
		Width:  10,
		Height: 5,
	}

	triangle := figures.Triangle{
		Base:   8,
		Height: 6,
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

func calculateArea(s any) (float64, error) {
	if v, ok := s.(figures.Shape); ok {
		return v.CalculateArea(), nil
	}

	return 0, fmt.Errorf("%s", "Error: The passed object is not a figure")
}

func getReport(r figures.Repoter) string {
	return r.GetReport()
}
