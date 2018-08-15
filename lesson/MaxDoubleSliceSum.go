package solution

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Solution(A []int) int {
	n := len(A)
	a := make([]int, n)
	b := make([]int, n)
	for i := 1; i < n; i++ {
		a[i] = Max(a[i-1]+A[i], 0)
	}
	for i := n - 2; i > 0; i-- {
		b[i] = Max(b[i+1]+A[i], 0)
	}
	max := 0
	for i := 1; i < n-1; i++ {
		max = Max(a[i-1]+b[i+1], max)
	}
	return max
}
