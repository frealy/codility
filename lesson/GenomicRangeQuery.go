package solution

func Solution(S string, P []int, Q []int) []int {
	n := len(S)
	x := make([][]int, 3)
	for i := 0; i < 3; i++ {
		x[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		a, c, g := 0, 0, 0
		if string(S[i]) == "A" {
			a = 1
		}
		if string(S[i]) == "C" {
			c = 1
		}
		if string(S[i]) == "G" {
			g = 1
		}
		x[0][i+1] = x[0][i] + a
		x[1][i+1] = x[1][i] + c
		x[2][i+1] = x[2][i] + g
	}
	n = len(P)
	var ans []int
	for i := 0; i < n; i++ {
		from := P[i]
		to := Q[i] + 1
		if x[0][to]-x[0][from] > 0 {
			ans = append(ans, 1)
		} else if x[1][to]-x[1][from] > 0 {
			ans = append(ans, 2)
		} else if x[2][to]-x[2][from] > 0 {
			ans = append(ans, 3)
		} else {
			ans = append(ans, 4)
		}
	}

	return ans
}
