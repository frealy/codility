package solution

func Gcd(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return Gcd(b, a%b)
	}
}
func Solution(N int, M int) int {
	x := Gcd(N, M)
	i := 0
	count := 0
	if x == 1 {
		return N
	}
	for i < N {
		i = i + x
		count++
	}
	return count
}
