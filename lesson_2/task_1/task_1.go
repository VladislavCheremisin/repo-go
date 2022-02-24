package main

import "fmt"

func main() {
	var rectangleLength, rectangleWidth float64

	fmt.Print("Введите длинну прямоугольника: ")
	fmt.Scanln(&rectangleLength)

	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scanln(&rectangleWidth)

	fmt.Println("Площадь прямоугольника =", (rectangleLength * rectangleWidth))
}
