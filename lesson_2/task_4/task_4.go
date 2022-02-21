package main

import (
	"fmt"
)

func main() {
	var number int
	var flagNumber bool = true
	fmt.Print("Введите целое число: ")
	fmt.Scanln(&number)

	for i := 2; i <= number; i++ {
		flagNumber = false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flagNumber = true
			}
		}
		if !flagNumber {
			fmt.Println(i)
		}
	}

}
