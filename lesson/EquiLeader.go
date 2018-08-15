//https://app.codility.com/demo/results/trainingRF4N52-M5G/
package solution

func Solution(A []int) int {
	n := len(A)
	dom := 0
	max := 0
	ans := 0
	val := make(map[int]int)
	for _, v := range A {
		val[v]++
	}
	for i, v := range val {
		if max < v {
			max = v
			dom = i
		}
	}

	ldrCount := 0
	for i := 0; i < n; i++ {
		if A[i] == dom {
			ldrCount++
		}
		leadersInRightPart := max - ldrCount
		if ldrCount > (i+1)/2 && leadersInRightPart > (n-i-1)/2 {
			ans++
		}
	}

	return ans
}
