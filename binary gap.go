package solution

import (
	"strconv"
)

func Solution(N int) int {
	val := strconv.FormatInt(int64(N), 2)
	count := 0
	start := false
	end := false
	max := 0
	for i := 0; i < len(val); i++ {
		if string(val[i]) == "0" && !start {
			count = count + 1
			start = true
		} else if string(val[i]) == "0" && start {
			count = count + 1
		}
		if string(val[i]) == "1" && start {
			if count > max {
				max = count
				count = 0
			} else {
				count = 0
			}
			start = false
			end = true
		}
	}
	if !end {
		max = 0
	}
	return max
}
