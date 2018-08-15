package solution

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Solution(A []int) int {
	n := len(A)
	if n == 1 || n == 0 {
		return 0
	}
	max1 := 0
	max2 := 0
	minPrice := A[0]
	for i := 1; i < n; i++ {
		max1 = Max(0, A[i]-minPrice)
		minPrice = Min(minPrice, A[i])
		max2 = Max(max1, max2)
	}

	return max2
}
