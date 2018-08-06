package solution

func merge(A []int, left []int, right []int) int {
	i, j, count := 0, 0, 0
	for i < len(left) || j < len(right) {
		if i == len(left) {
			A[i+j] = right[j]
			j++
		} else if j == len(right) {
			A[i+j] = left[i]
			i++
		} else if left[i] <= right[j] {
			A[i+j] = left[i]
			i++
		} else {
			A[i+j] = right[j]
			count = count + len(left) - i
			j++
		}
	}
	return count
}

func Solution(A []int) int {
	if len(A) < 2 {
		return 0
	}

	var m int = (len(A) + 1) / 2
	left := []int{}
	right := []int{}

	for i, val := range A {
		if i < m {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	return merge(A, left, right) + Solution(left) + Solution(right)
}
