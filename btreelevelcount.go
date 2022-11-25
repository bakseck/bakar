package bakar

func BTreeLevelCount(root *TreeNode) int {
	count := 0 
	if root == nil {
		return 0
	}
	count++
	rightCount := count + BTreeLevelCount(root.Right)
	leftCount := count + BTreeLevelCount(root.Left)
	if rightCount > leftCount {
		return rightCount
	} else {
		return leftCount
	}
}

