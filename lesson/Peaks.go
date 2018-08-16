// https://app.codility.com/demo/results/trainingTNB258-ZK8/
package solution

func Solution(A []int) int {
	n := len(A)
	var ans []int
	//find all the peak
	for i := 1; i < n-1; i++ {
		if A[i-1] < A[i] && A[i+1] < A[i] {
			ans = append(ans, i)
		}
	}

	for i := 1; i <= n; i++ {
		if n%i != 0 { //skip if non divisible
			continue
		}
		find := 0
		group := n / i
		check := true
		// Find whether every group has a peak
		for _, peakId := range ans {
			if peakId/i > find {
				check = false
				break
			}
			if peakId/i == find {
				find++
			}
		}
		if find != group {
			check = false
		}
		if check {
			return group
		}
	}
	return 0
}
