package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var a, b, res float64
	var op string

	firstOperandUsual := "Введите первое число: "
	secondOperandUsual := "Введите второе число: "
	firstOperandDegree := "Введите число, которое возводим в степень: "
	secondOperandDegree := "Введите степень: "
	firstOperandLog := "Введите основание: "
	secondOperandLog := "Введите число: "

	fmt.Print("Введите арифметическую операцию (+, -, *, /, **, log): ")
	fmt.Scanln(&op)

	if op == "+" || op == "-" || op == "*" || op == "/" {
		fmt.Print(firstOperandUsual)
		fmt.Scanln(&a)

		fmt.Print(secondOperandUsual)
		fmt.Scanln(&b)
	} else if op == "**" {
		fmt.Print(firstOperandDegree)
		fmt.Scanln(&a)

		fmt.Print(secondOperandDegree)
		fmt.Scanln(&b)
	} else if op == "log" {
		fmt.Print(firstOperandLog)
		fmt.Scanln(&a)

		fmt.Print(secondOperandLog)
		fmt.Scanln(&b)
	} else {
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "**":
		res = math.Pow(a, b)
	case "log":
		res = getLog(a, b)
	case "/":
		if b == 0 {
			fmt.Println("На 0 делить нельзя")
			return
		} else {
			res = a / b
		}

	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res) // округляем до 2 знков после запятой

}
func getLog(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}
