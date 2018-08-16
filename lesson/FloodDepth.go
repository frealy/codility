// https://app.codility.com/demo/results/training7W3PT7-TU5/
package solution

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Solution(A []int) int {
	highestIdx := 0
	lowestIdx := 0
	max := 0

	for i := 1; i < len(A); i++ {
		current := A[i]
		highest := A[highestIdx]
		lowest := A[lowestIdx]
		if current > highest {
			max = Max(highest-lowest, max)
			highestIdx = i
			lowestIdx = highestIdx
		} else if current > lowest {
			max = Max(current-lowest, max)
		} else if current < lowest {
			lowestIdx = i
		}
	}
	return max
}
