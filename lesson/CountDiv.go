package solution

func Solution(A int, B int, K int) int {
	count1 := A / K
	count2 := B / K

	count := count2 - count1
	var mod int
	if A%K == 0 {
		mod = 1
	}
	return count + mod
}
