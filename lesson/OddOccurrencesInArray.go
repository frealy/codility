package solution

func Solution(A []int) int {
	val := make(map[int]int)
	value := 0
	for i := 0; i < len(A); i++ {
		if val[A[i]] == 0 {
			val[A[i]] = val[A[i]] + 1
			value = A[i]
		} else {
			delete(val, A[i])
		}
	}
	for i, _ := range val {
		value = i
	}
	return value

}
