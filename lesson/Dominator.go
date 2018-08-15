package solution

func Solution(A []int) int {
	n := len(A)
	dom := 0
	max := 0
	val := make(map[int]int)
	ans := make(map[int]int)
	for i, v := range A {
		val[v]++
		if ans[v] == 0 {
			ans[v] = i
		}
	}
	for i, v := range val {
		if max < v {
			max = v
			dom = i
		}
	}
	if max > n/2 {
		return ans[dom]
	} else {
		return -1
	}
}
