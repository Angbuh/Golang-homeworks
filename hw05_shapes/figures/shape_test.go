package figures

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateArea_Rectangle(t *testing.T) {
	testdata := Rectangle{
		Width:  10,
		Height: 5,
	}
	var expected float64 = 50

	res := testdata.CalculateArea()
	assert.Equal(t, expected, res)
}

func TestCalculateArea_Circle(t *testing.T) {
	const EPSILON float64 = 0.0000001
	testdata := Circle{
		Radius: 5,
	}
	expected := 78.5398163

	res := testdata.CalculateArea()
	assert.LessOrEqual(t, math.Abs(expected-res), EPSILON)
}

func TestCalculateArea_Triangle(t *testing.T) {
	testdata := Triangle{
		Base:   8,
		Height: 6,
	}
	expected := 24.000000

	res := testdata.CalculateArea()
	assert.Equal(t, expected, res)
}

func TestGetReportRectangle(t *testing.T) {
	testdata := Rectangle{
		Width:  10,
		Height: 5,
	}
	expected := `Прямоугольник: ширина: 10.000000, высота: 5.000000
Площадь: 50.000000`

	res := testdata.GetReport()
	assert.Equal(t, expected, res)
}

func TestGetReportCircle(t *testing.T) {
	testdata := Circle{
		Radius: 5,
	}
	expected := `Круг: радиус: 5.000000
Площадь: 78.539816`

	res := testdata.GetReport()
	assert.Equal(t, expected, res)
}

func TestGetReportTriangle(t *testing.T) {
	testdata := Triangle{
		Base:   8,
		Height: 6,
	}
	expected := `Треугольник: основание: 8.000000, высота: 6.000000
Площадь: 24.000000`

	res := testdata.GetReport()
	assert.Equal(t, expected, res)
}
