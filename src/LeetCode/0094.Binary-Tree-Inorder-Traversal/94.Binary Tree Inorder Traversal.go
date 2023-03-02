package leetcode

import "structures"

func inorderTraversal(root *structures.TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *structures.TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}
