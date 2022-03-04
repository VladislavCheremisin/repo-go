package main

import (
	"fmt"
)

func fibbonachi(number uint) uint {

	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	return fibbonachi(number-1) + fibbonachi(number-2)
}

func main() {

	var number uint
	fmt.Print("Введите порядковый номер числа Фибоначи начиная с 0 позиции: ")
	fmt.Scanln(&number)

	fmt.Println(fibbonachi(number))
}
