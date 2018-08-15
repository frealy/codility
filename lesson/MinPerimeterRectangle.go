package solution

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Solution(N int) int {
	min := 1000000000000
	i := 1
	for i*i <= N {
		if N%i == 0 {
			if i*i == N {
				min = Min(min, 4*i)
			} else {
				min = Min(min, 2*((N/i)+i))
			}
		}
		i++
	}
	return min
}
