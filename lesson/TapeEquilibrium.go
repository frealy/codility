package solution

import "math"

func Solution(A []int) int {
	n := len(A)
	sumRight := float64(0)
	for i := 1; i < n; i++ {
		sumRight = sumRight + float64(A[i])
	}

	sumLeft := float64(A[0])
	min := math.Abs(sumRight - sumLeft)

	for i := 1; i < n; i++ {
		if math.Abs(sumRight-sumLeft) < min {
			min = math.Abs(sumRight - sumLeft)
		}
		sumLeft = sumLeft + float64(A[i])
		sumRight = sumRight - float64(A[i])
	}
	return int(min)
}
