package solution

import "sort"

func Solution(A []int) int {
	sort.Ints(A)
	n := len(A)
	missing := 1
	max := 0
	for i := 0; i < n; i++ {
		if A[i] > 0 && missing == A[i] {
			missing = missing + 1
		}

		if max < A[i] {
			max = A[i]
		}
	}

	if missing == max {
		return missing + 1
	} else {
		return missing
	}
}
