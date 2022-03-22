package main

import (
	"fmt"
)

func main() {
	entersValues()
}

var number int
var fibonacciNumbersMap = map[int]int{
	0: 0,
	1: 1}

func countFibbonachi(number int) int {
	return ((number - 1) + (number - 2))
}

func entersValues() {
	fmt.Print("Введите порядковый номер числа Фибоначи начиная с 0 позиции: ")
	fmt.Scanln(&number)
	checkFibonacci(number)
}

func checkFibonacci(number int) {
	if v, found := fibonacciNumbersMap[number]; found {
		fmt.Println(fibonacciNumbersMap[v])
	} else {
		fibonacciNumbersMap[number] = countFibbonachi(number)
		fmt.Println(countFibbonachi(number))
	}
	entersValues()
}
