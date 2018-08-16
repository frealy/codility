// https://app.codility.com/demo/results/trainingAT7PHT-E96/
package solution

func Gcd(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return Gcd(b, a%b)
	}
}
func HasSamePrime(a, b int) bool {
	gcd := Gcd(a, b)
	var gcdA, gcdB int
	for a != 1 {
		gcdA = Gcd(a, gcd)
		if gcdA == 1 {
			break
		}
		a = a / gcdA
	}
	if a != 1 {
		return false
	}
	for b != 1 {
		gcdB = Gcd(b, gcd)
		if gcdB == 1 {
			break
		}
		b = b / gcdB
	}
	return b == 1
}
func Solution(A []int, B []int) int {
	n := len(A)
	count := 0
	for i := 0; i < n; i++ {
		if HasSamePrime(A[i], B[i]) {
			count++
		}
	}
	return count
}
