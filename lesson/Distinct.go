package solution

func Solution(A []int) int {
	ans := make(map[int]int)
	for _, val := range A {
		ans[val] += 1
	}
	index := len(ans)
	return index
}
