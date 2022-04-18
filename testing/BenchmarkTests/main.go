package main

import (
	"BenchmarkTests/sorter/permutation"
	"BenchmarkTests/sorter/standard"
)

func main() {
	inputNums := []int{4, 9, 5, 4, 6}
	permutation.Permutation(inputNums)
	standard.Standard(inputNums)
}
