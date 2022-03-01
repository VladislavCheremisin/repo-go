package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputNums := [5]int64{}
	fmt.Println("Введите последовательно 5 чисел: ")
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < len(inputNums); i++ {
		if scanner.Scan() {
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			inputNums[i] = num
			x := inputNums[i]
			j := i
			for ; j >= 1 && inputNums[j-1] > x; j-- {
				inputNums[j] = inputNums[j-1]
			}
			inputNums[j] = x

		} else {
			panic("you must input 5 numbers")
		}
	}
	fmt.Print("Введенные числа отсортированы по возрастанию: ")
	fmt.Println(inputNums)
}
