package standard

import (
	"sort"
	"testing"
)

func BenchmarkStandard(b *testing.B) {
	var inputNums = []int{4, 9, 5, 4, 6}
	for i := 0; i < b.N; i++ {
		sort.Ints(inputNums)
	}
}
