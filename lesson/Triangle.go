package solution

import "sort"

func Solution(A []int) int {
	sort.Ints(A)
	n := len(A)
	check := false
	for i := 0; i < n-2; i++ {
		if (A[i]+A[i+1] > A[i+2]) && (A[i+1]+A[i+2] > A[i]) && (A[i]+A[i+2] > A[i+1]) {
			check = true
			break
		}
	}
	if check {
		return 1
	} else {
		return 0
	}
}
