package solution

import "math"

func Solution(A []int) int {
	n := len(A)
	sum := float64(0)
	for i := 0; i < n; i++ {
		sum = sum + float64(A[i])
	}

	add := float64(A[0])
	min := math.Abs((sum - add) - add)
	if n == 2 {
		return int(min)
	}
	for i := 1; i < n; i++ {
		add = add + float64(A[i])
		if math.Abs((sum-add)-add) < min {
			min = math.Abs((sum - add) - add)
		}
	}
	return int(min)
}
