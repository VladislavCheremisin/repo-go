package permutation

func Permutation(InputNums []int) {
	for i := 0; i < len(InputNums); i++ {
		x := InputNums[i]
		j := i
		for ; j >= 1 && InputNums[j-1] > x; j-- {
			InputNums[j] = InputNums[j-1]
		}
		InputNums[j] = x
	}
}
