package solution

func Solution(A []int, K int) []int {
	n := len(A)
	var val []int
	if n == 0 || n == 1 || K%n == 0 {
		return A
	}
	if K > n {
		K = K % n
	}
	m := K
	for i := n - m; i < n; i++ {
		val = append(val, A[i])
	}
	for i := 0; i < n-m; i++ {
		val = append(val, A[i])
	}
	return val
}
