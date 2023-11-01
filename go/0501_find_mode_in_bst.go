/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findMode(root *TreeNode) []int {
	d := make(map[int]int)
	max_freq := 0

	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		_, exists := d[node.Val]
		if !exists {
			d[node.Val] = 1
		} else {
			d[node.Val]++
		}

		max_freq = max(max_freq, d[node.Val])

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	ans := []int{}

	for key, value := range d {
		if value == max_freq {
			ans = append(ans, key)
		}
	}

	return ans

}