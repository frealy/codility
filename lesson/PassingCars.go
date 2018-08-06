package solution

func Solution(A []int) int {
	n := len(A)
	pass := 0
	x := 0
	for i := 0; i < n; i++ {
		if A[i] == 1 {
			pass = pass + 1*x
		} else if A[i] == 0 {
			x = x + 1
		}

		if pass > 1000000000 {
			return -1
		}
	}

	return pass
}
