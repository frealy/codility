package solution

func Solution(A []int) int {
	n := len(A)
	index := 0
	min := float64(A[0]+A[1]) / float64(2)
	for i := 1; i < n-2; i++ {
		if min > float64(A[i]+A[i+1])/2.0 {
			min = float64(A[i]+A[i+1]) / 2.0
			index = i
		}
		if min > float64(A[i]+A[i+1]+A[i+2])/3.0 {
			min = float64(A[i]+A[i+1]+A[i+2]) / 3.0
			index = i
		}
	}
	if min > float64(A[n-1]+A[n-2])/2.0 {
		min = float64(A[n-1]+A[n-2]) / 2.0
		index = n - 2
	}
	return index
}
