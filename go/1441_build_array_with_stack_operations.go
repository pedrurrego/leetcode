func buildArray(target []int, n int) []string {
	i := 0
	// out := []string{}
	out := make([]string, 0, 2*n)

	for j := 1; j < (n + 1); j++ {
		if i < len(target) {
			if j == target[i] {
				out = append(out, "Push")
				i++
			} else {
				out = append(out, "Push")
				out = append(out, "Pop")
			}
		}
	}

	return out
}