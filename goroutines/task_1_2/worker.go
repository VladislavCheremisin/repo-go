package main

import (
	"fmt"
)

// this way is bad.

func worker(i int) {
	j := 1
	// start 1000 works
	for ; j <= 1000; j++ {
		println("gorutine", i, "is doing operation", j)
	}
}

func main() {
	num := 1000
	i := 1
	// start 1000 workers
	for ; i < num; i++ {
		go worker(i)
	}

	// Press enter in the end
	var input string
	fmt.Scanln(&input)

	println("all workers have done sum, last workers number, witch started working is =", i)
}
