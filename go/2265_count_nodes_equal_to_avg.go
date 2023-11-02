/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {

	count := 0

	var dfs func(*TreeNode) []int // dfs is a variable that holds a reference to a function

	dfs = func(node *TreeNode) []int {
		if node == nil {
			out := []int{0, 0}
			return out
		}

		L := dfs(node.Left)
		R := dfs(node.Right)

		cur_sum := node.Val + L[0] + R[0]
		cur_n := 1 + L[1] + R[1]

		avg := cur_sum / cur_n

		if avg == node.Val {
			count++
		}

		out := []int{cur_sum, cur_n}

		return out

	}

	dfs(root)

	return count

}