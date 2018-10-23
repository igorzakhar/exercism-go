package summultiples

func SumMultiples(lim int, div ...int) int {
	sum := 0
	for i := 1; i < lim; i++ {
		for _, d := range div {
			if i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
