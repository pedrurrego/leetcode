func getLastMoment(n int, left []int, right []int) int {
	ans := 0

	for _, val := range left {
		ans = max(ans, val)
	}

	for _, val := range right {
		ans = max(ans, n-val)
	}

	return ans
}