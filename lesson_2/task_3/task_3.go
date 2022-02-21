package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64

	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&number)
	var numberModul = int(math.Abs(number))
	fmt.Println("Сотен =", (numberModul / 100))
	fmt.Println("Десятков =", (numberModul % 100 / 10))
	fmt.Println("Единиц=", (numberModul % 100 % 10))
}
