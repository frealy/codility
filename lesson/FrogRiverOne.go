package solution

func Solution(X int, A []int) int {
	n := len(A)
	x := make([]bool, X+1)
	for i := 0; i < n; i++ {
		if !x[A[i]] {
			x[A[i]] = true
			X--
		}

		if X == 0 {
			return i
		}
	}
	return -1
}
