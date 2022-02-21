package main

import (
	"fmt"
	"math"
)

func main() {
	var circleArea float64

	fmt.Print("Введите площадь окружности: ")
	fmt.Scanln(&circleArea)

	fmt.Println("Диаметр окружности =", math.Sqrt(4*circleArea/math.Pi))
	fmt.Println("Длина окружности =", math.Sqrt(circleArea*4*math.Pi))
}
