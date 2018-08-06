package solution

func Solution(N int, A []int) []int {
	x := make([]int, N)
	n := len(A)
	max := 0
	add := 0
	for i := 0; i < n; i++ {
		if A[i] == N+1 {
			add = max
		} else {
			if x[A[i]-1] < add {
				x[A[i]-1] = add + 1
			} else {
				x[A[i]-1]++
			}

			if max < x[A[i]-1] {
				max = x[A[i]-1]
			}
		}
	}

	for i := 0; i < len(x); i++ {
		if x[i] < add {
			x[i] = add
		}
	}

	return x
}
