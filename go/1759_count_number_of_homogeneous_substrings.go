import (
	"fmt"
	"math"
)

func countHomogenous(s string) int {
	ans := 0
	cur_ans := 1
	mod := int(math.Pow(10, 9) + 7)

	if len(s) == 1 {
		return 1
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cur_ans += 1

		} else {
			ans += (cur_ans * (cur_ans + 1)) / 2 % mod
			cur_ans = 1
		}

		if i == len(s)-2 {
			ans += (cur_ans * (cur_ans + 1)) / 2 % mod
		}
	}

	return ans
}