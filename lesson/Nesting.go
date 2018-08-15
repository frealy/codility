package solution

func Solution(S string) int {
	n := len(S)
	if n%2 == 1 {
		return 0
	}

	var open, close int
	for i := 0; i < n; i++ {
		a := S[i]
		if a == 40 || a == 91 || a == 123 {
			open++
		}
		if a == 41 || a == 93 || a == 125 {
			close++
		}
	}

	if open == close {
		return 1
	} else {
		return 0
	}
}
