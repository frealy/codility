package solution

import "math"

func Solution(A []int) int {
	n := len(A)
	ans := make(map[float64]bool)
	for i := 0; i < n; i++ {
		val := math.Abs(float64(A[i]))
		if ans[val] == false {
			ans[val] = true
		}
	}
	count := len(ans)
	return count
}
