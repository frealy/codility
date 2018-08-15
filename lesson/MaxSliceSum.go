package solution

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Solution(A []int) int {
	n := len(A)
	max1 := A[0]
	maxAns := A[0]
	for i := 1; i < n; i++ {
		max1 = Max(A[i], A[i]+max1)
		maxAns = Max(maxAns, max1)
	}
	return maxAns
}
