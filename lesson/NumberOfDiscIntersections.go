package solution

func Solution(A []int) int {
	n := len(A)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		right := 0
		//if i+A[i]<= n-1, that's it, extract this i+A[i], let ans[i+A[i]]++, means there is one disk that i+A[i]
		if n-i-1 >= A[i] {
			right = i + A[i]
		} else {
			right = n - 1
		}
		ans[right]++
	}

	for i := 1; i < n; i++ {
		ans[i] += ans[i-1] //ans[i] means that there are ans[i] number of values that <= i
	}

	val := n * (n - 1) / 2

	for i := 0; i < n; i++ {
		left := 0

		if A[i] > i {
			left = 0
		} else {
			left = i - A[i] // Find the positive i-A[i].
		}

		if left > 0 {
			val -= ans[left-1] //Find the number that is smaller than 1-A[i], sum[n-1] will never be used as we only need sum[n-1-1] at most.
		}
	}

	if val > 10000000 {
		return -1
	}

	return val
}
