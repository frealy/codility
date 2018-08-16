// https://app.codility.com/demo/results/training3NGBSY-JWJ/
package solution

import "strings"
import "regexp"

func CountAlpha(S string) int {
	re := regexp.MustCompile("[a-zA-Z]+")
	v := strings.Join(re.FindAllString(S, -1), "")
	val := len(strings.Replace(v, " ", "", -1))
	return val
}
func CountNum(S string) int {
	re := regexp.MustCompile("[0-9]+")

	v := strings.Join(re.FindAllString(S, -1), "")
	val := len(strings.Replace(v, " ", "", -1))
	return val
}
func Solution(S string) int {
	s := strings.Split(S, " ")
	ans := -1

	for _, v := range s {
		x := CountAlpha(v)
		y := CountNum(v)
		if x%2 == 0 && y%2 == 1 && len(v) > ans && x+y == len(v) {
			ans = len(v)
		}
	}
	return ans
}
