package fib

var (
	cache map[int]int
)

func init() {
	cache = make(map[int]int)
}

// Fibonacсi returns fibonacci number int
func Fibonacсi(number int) int {

	if number < 2 {
		return number
	}

	if r, ok := cache[number]; ok {
		return r
	}
	sum := Fibonacсi(number-1) + Fibonacсi(number-2)

	cache[number] = sum
	return sum
}
