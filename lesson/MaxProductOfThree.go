package solution

import "sort"

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Solution(A []int) int {
	sort.Ints(A)
	n := len(A)
	max := Max(A[0]*A[1]*A[n-1], A[n-1]*A[n-2]*A[n-3])
	return max
}
