package solution

func Solution(S string) int {
	if S == "" {
		return -1
	}
	var ind int = 0
	var n int = 0
	if len(S)%2 == 1 {
		n = len(S)/2 + 1
	} else {
		return -1
	}
	for i := 0; i < n; i++ {
		if S[i] == S[len(S)-1-i] {
			ind = i
		} else {
			return -1
		}
	}
	return ind
}
