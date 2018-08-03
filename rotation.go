package solution

func Solution(A []int, K int) []int {
	if len(A) == 0 || len(A) == 1 || K == 0 || (len(A)%K == 0 && K != 1) {
		return A
	}
	var n int
	if len(A) < K {
		n = K % len(A)
	} else {
		n = K
	}

	var B []int
	var m int = len(A)
	for i := 0; i < len(A); i++ {
		if m-n+i < len(A) {
			B = append(B, A[m-n+i])
		} else {
			B = append(B, A[i-n])
		}
	}
	return B
}
