package solution

func Solution(T *Tree) int {
	if T == nil {
		return -1
	}
	val1 := 1 + Solution(T.L)
	val2 := 1 + Solution(T.R)

	if val1 > val2 {
		return val1
	} else {
		return val2
	}
}
