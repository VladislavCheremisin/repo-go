package main

import (
	"TestDrivenTests/fib"
	"fmt"
)

func main() {

	var number int
	fmt.Print("Введите порядковый номер числа Фибоначи начиная с 0 позиции: ")
	fmt.Scanln(&number)

	fmt.Println(fib.Fibonachi(number))
}
