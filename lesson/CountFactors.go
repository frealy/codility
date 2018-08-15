package solution

func Solution(N int) int {
	count := 0
	i := 1
	for i*i <= N {
		if N%i == 0 {
			if i*i == N {
				count += 1
			} else {
				count += 2
			}
		}
		i++
	}
	return count
}
