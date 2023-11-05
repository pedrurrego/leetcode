func getWinner(arr []int, k int) int {
	max_elem := arr[0]

	for _, val := range arr {
		if val > max_elem {
			max_elem = val
		}
	}

	curr := arr[0]
	win_streak := 0

	for i := 1; i < len(arr); i++ {
		opponent := arr[i]

		if curr > opponent {
			win_streak += 1
		} else {
			curr = opponent
			win_streak = 1
		}

		if curr == max_elem || win_streak == k {
			return curr
		}
	}

	return 0
}